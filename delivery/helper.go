package delivery

import (
	"github.com/Gary-Gs/go-clean-arch/domain"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	switch err.(type) {
	case validator.ValidationErrors:
		return http.StatusBadRequest
	}

	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	case domain.ErrContextTimeout:
		return http.StatusRequestTimeout
	default:
		return http.StatusBadRequest
	}
}
