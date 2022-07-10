package middleware

import (
	"context"
	"github.com/Gary-Gs/go-clean-arch/config"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// GoMiddleware represent the data-struct for middleware
type GoMiddleware struct {
	// another stuff , may be needed by middleware
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

		m.makeLogEntry(c).Debug("API Request:")
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
