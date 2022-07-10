package config

type Configs struct {
	AppConfig AppConfig
	Database  Database
}

type AppConfig struct {
	LogLevel          string
	ApiPort           string
	ContextTimeOut    int
	TimeZone          string
	LatencyWarningSec float64
}

type Database struct {
	Host                      string
	Port                      string
	Username                  string
	Password                  string
	Name                      string
	Debug                     bool
	MaxIdleConnections        int
	MaxConnectionLifeTimeSecs int
}
