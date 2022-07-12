package logging

import (
	"context"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// Infof ...
func Infof(ctx context.Context, format string, args ...interface{}) {
	log.WithFields(log.Fields{echo.HeaderXRequestID: ctx.Value(echo.HeaderXRequestID)}).
		Infof(format, args...)
}

// Debugf ...
func Debugf(ctx context.Context, format string, args ...interface{}) {
	log.WithFields(log.Fields{echo.HeaderXRequestID: ctx.Value(echo.HeaderXRequestID)}).
		Debugf(format, args...)
}

// Warnf ...
func Warnf(ctx context.Context, format string, args ...interface{}) {
	log.WithFields(log.Fields{echo.HeaderXRequestID: ctx.Value(echo.HeaderXRequestID)}).
		Warnf(format, args...)
}

// Errorf ...
func Errorf(ctx context.Context, format string, args ...interface{}) {
	log.WithFields(log.Fields{echo.HeaderXRequestID: ctx.Value(echo.HeaderXRequestID)}).
		Errorf(format, args...)
}
