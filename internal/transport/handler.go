package handler

import (
	"sync"

	"github.com/Levap123/adverts/internal/service"
	"github.com/Levap123/adverts/internal/validator"
	"github.com/Levap123/adverts/pkg/json"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	service *service.Service
	js      *json.JSONSerializer
	lg      *logrus.Logger
	v       *validator.Validator
	mu      *sync.Mutex
}

func NewHandler(service *service.Service, js *json.JSONSerializer, lg *logrus.Logger, v *validator.Validator) *Handler {
	return &Handler{
		service: service,
		js:      js,
		lg:      lg,
		v:       v,
		mu:      &sync.Mutex{},
	}
}

type ErrorResponse struct {
	Message string `json:"message,omitempty"`
}
