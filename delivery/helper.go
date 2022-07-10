package delivery

import (
	"github.com/Gary-Gs/go-clean-arch/domain"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"net/http"
)

const OK = "OK"

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

type HttpResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

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
		logrus.Error(err.Error())
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		logrus.Error(err.Error())
		return http.StatusInternalServerError
	}
}
