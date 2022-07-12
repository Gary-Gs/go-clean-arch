package middleware

import (
	"context"
	"github.com/Gary-Gs/go-clean-arch/common"
	"github.com/Gary-Gs/go-clean-arch/config"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	// URIs to exclude in the logs
	pathToExclude = []string{
		"/health",
	}
)

// GoMiddleware represent the data-struct for middleware
type GoMiddleware struct {
	// another stuff , may be needed by middleware
	Config config.Configs
}

type CustomLogWriter struct {
	Config config.Configs
}

// InitMiddleware initialize the middleware
func InitMiddleware(conf config.Configs) *GoMiddleware {
	return &GoMiddleware{
		Config: conf,
	}
}

// CORS will handle the CORS middleware
func (m *GoMiddleware) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodOptions, http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodPost,
			http.MethodDelete},
	})(next)
}

func (m *GoMiddleware) GenerateRequestID(next echo.HandlerFunc) echo.HandlerFunc {
	return middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Generator: func() string {
			return uuid.New().String()
		},
	})(next)
}

func (m *GoMiddleware) MiddlewareLogging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// inject request ID into context
		id := c.Response().Header().Get(echo.HeaderXRequestID)
		c.SetRequest(c.Request().WithContext(context.WithValue(c.Request().Context(), echo.HeaderXRequestID, id)))

		// exclude some URI in logs
		if !common.ContainsIgnoreCase(pathToExclude, c.Path()) &&
			m.Config.FeatureFlag.EnableExcludeUrl && !strings.Contains(c.Path(), "/swagger") {
			m.makeLogEntry(c).Debug("API Request:")
		}
		return next(c)
	}
}

func (m *GoMiddleware) makeLogEntry(c echo.Context) *log.Entry {
	if c == nil {
		return log.WithFields(log.Fields{
			"timestamp": time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	return log.WithFields(log.Fields{
		"timestamp":           time.Now().Format("2006-01-02 15:04:05"),
		"method":              c.Request().Method,
		"uri":                 c.Request().URL.String(),
		echo.HeaderXRequestID: c.Response().Header().Get(echo.HeaderXRequestID),
	})
}

func (c *CustomLogWriter) Write(p []byte) (n int, err error) {
	latencyReg, _ := regexp.Compile("\"latency\":[0-9]+")
	latencyString := strings.ReplaceAll(latencyReg.FindString(string(p)), "\"latency\":", "")
	latency, err := strconv.ParseFloat(latencyString, 64)
	if err != nil {
		return 0, err
	}

	if latency/1000000000 > c.Config.AppConfig.LatencyWarningSec {
		log.Warnf("API latency: %f seconds, log: %s", latency/1000000000, string(p))
	}
	return len(p), nil
}
