package app

import (
	"encoding/json"
	"fmt"
	"github.com/Gary-Gs/go-clean-arch/config"
	"github.com/Gary-Gs/go-clean-arch/delivery"
	"github.com/Gary-Gs/go-clean-arch/domain"
	mw "github.com/Gary-Gs/go-clean-arch/middleware"
	repo "github.com/Gary-Gs/go-clean-arch/repository/mysql"
	"github.com/Gary-Gs/go-clean-arch/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	log2 "log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type App struct {
	Dependencies Dependencies
	Usecase      Usecase
	Configs      config.Configs
}

type Dependencies struct {
	Database *gorm.DB
}

type Usecase struct {
	ArticleUsecase domain.ArticleUsecase
}

type Repositories struct {
	ArticleRepository domain.ArticleRepository
	AuthorRepository  domain.AuthorRepository
}

type CustomValidator struct {
	validator *validator.Validate
}

func (a *App) InitApiEcho() *echo.Echo {
	e := echo.New()
	e.HTTPErrorHandler = customHTTPErrorHandler
	midW := mw.InitMiddleware(a.Configs)
	e.Use(midW.GenerateRequestID)
	e.Use(midW.CORS)
	e.Use(midW.MiddlewareLogging)
	e.Use(middleware.Recover())
	e.Validator = &CustomValidator{validator: validator.New()}
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: &mw.CustomLogWriter{Config: a.Configs},
	}))
	// handlers
	delivery.NewArticleHandler(e, a.Usecase.ArticleUsecase)
	return e
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		log.Debugf("validation failed, err=%s", err.Error())
		return err
	}
	return nil
}

func getLumberjackLogger() *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   "app.log", // Log file path
		MaxSize:    10,        // Maximum file size before rotation in MB
		MaxBackups: 30,        // Maximum number of old log files to retain
		MaxAge:     28,        // Maximum number of days to retain old log files
		Compress:   false,     // Whether to compress the rotated log files
	}
}

func GetNewInstance(file string) (App, error) {
	var app App
	configs, err := LoadConfigFile(file)
	if err != nil {
		return app, fmt.Errorf("load config file error: %v", err.Error())
	}
	deps, err := initDependencies(configs)
	if err != nil {
		return app, fmt.Errorf("failed to initialize dependencies err: %s", err.Error())
	}
	repos := initRepos(deps.Database)
	uc := initUseCases(repos, configs)
	app = App{
		Dependencies: deps,
		Configs:      configs,
		Usecase:      uc,
	}
	return app, nil
}

func LoadConfigFile(file string) (config.Configs, error) {
	var c config.Configs
	configs := viper.NewWithOptions()
	configs.SetConfigFile(file)
	err := configs.ReadInConfig()
	if err != nil {
		return c, fmt.Errorf("failed to read config file, err=%s", err.Error())
	}
	replacer := strings.NewReplacer(".", "_")
	configs.SetEnvKeyReplacer(replacer)
	configs.AutomaticEnv()
	err = configs.Unmarshal(&c)
	if err != nil {
		return c, fmt.Errorf("failed to unmarshal config file, err=%s", err.Error())
	}
	// logger initialization
	multiWriter := io.MultiWriter(os.Stdout, getLumberjackLogger())
	log.SetOutput(multiWriter)
	log.SetFormatter(&log.TextFormatter{
		ForceColors:            true,
		DisableColors:          false,
		FullTimestamp:          true,
		DisableLevelTruncation: true,
	})
	if c.AppConfig.LogLevel == "debug" {
		log.SetLevel(log.DebugLevel)
		log.SetReportCaller(false)
		data, err := json.MarshalIndent(c, "", "    ")
		if err != nil {
			return c, fmt.Errorf("failed to marshal config file, err=%s", err.Error())
		}
		log.Debugf("\n----------loaded configs----------\n %s", data)
		log.Debug("\n----------end configs------")
	} else if c.AppConfig.LogLevel == "info" {
		log.SetLevel(log.InfoLevel)
	} else if c.AppConfig.LogLevel == "warn" {
		log.SetLevel(log.WarnLevel)
	}
	return c, err
}

func getDBConnection(dbConf config.Database) (*gorm.DB, error) {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbConf.Username, dbConf.Password, dbConf.Host, dbConf.Port,
		dbConf.Name)
	val := url.Values{}
	val.Add("parseTime", "1")
	dns := fmt.Sprintf("%s?%s", connection, val.Encode())
	newLogger := logger.New(
		log2.New(os.Stdout, "\r\n", log2.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             300 * time.Millisecond, // Slow SQL threshold
			LogLevel:                  logger.Warn,            // Log level
			IgnoreRecordNotFoundError: false,                  // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                   // Disable color
		},
	)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger:          newLogger,
		CreateBatchSize: 100,
	})
	if err != nil {
		return nil, err
	}
	sDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sDB.SetMaxIdleConns(dbConf.MaxIdleConnections)
	sDB.SetConnMaxLifetime(time.Duration(dbConf.MaxConnectionLifeTimeSecs) * time.Second)
	if dbConf.Debug {
		return db.Debug(), err
	}
	return db, err
}

func initRepos(database *gorm.DB) Repositories {
	return Repositories{
		ArticleRepository: repo.NewMysqlArticleRepository(database),
		AuthorRepository:  repo.NewMysqlAuthorRepository(database),
	}
}

func initDependencies(configs config.Configs) (d Dependencies, err error) {
	if configs.FeatureFlag.EnableDB {
		d.Database, err = getDBConnection(configs.Database)
		if err != nil {
			return d, fmt.Errorf("failed to initialize DB connection: %s", err.Error())
		}
	}
	return
}

func initUseCases(repos Repositories, cnf config.Configs) Usecase {
	contextTO := time.Duration(cnf.AppConfig.ContextTimeOut) * time.Second
	au := usecase.NewArticleUsecase(repos.ArticleRepository, repos.AuthorRepository, contextTO)
	uc := Usecase{
		ArticleUsecase: au,
	}
	return uc
}

func customHTTPErrorHandler(err error, c echo.Context) {
	log.Error("server error: ", err.Error())
	_ = c.JSON(http.StatusInternalServerError, struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}{http.StatusInternalServerError, err.Error()})
}
