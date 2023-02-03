package handler

import (
	"github.com/Levap123/adverts/internal/service"
	"github.com/Levap123/adverts/pkg/json"
)

type Handler struct {
	service *service.Service
	js      *json.JSONSerializer
}

func NewHandler(service *service.Service, js *json.JSONSerializer) *Handler {
	return &Handler{
		service: service,
		js:      js,
	}
}
