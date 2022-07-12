package domain

import (
	"context"
	"errors"
)

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrNotFound            = errors.New("not found")
	ErrConflict            = errors.New("conflict")
	ErrContextTimeout      = context.DeadlineExceeded
)
