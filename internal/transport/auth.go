package handler

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/Levap123/adverts/internal/service"
)

type signUpBody struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	var input signUpBody
	if err := h.js.Read(r, &input); err != nil {
		h.lg.Errorln(err.Error())
		h.js.Send(w, http.StatusBadRequest, ErrorResponse{err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	userID, err := h.service.Create(ctx, input.Email, input.Password)
	if err != nil {
		h.lg.Errorln(err.Error())
		switch {
		case errors.Is(err, service.ErrInvalidEmail):
			h.js.Send(w, http.StatusBadRequest, ErrorResponse{service.ErrInvalidEmail.Error()})
		case errors.Is(err, service.ErrInvalidPassword):
			h.js.Send(w, http.StatusBadRequest, ErrorResponse{service.ErrInvalidPassword.Error()})
		default:
			h.js.Send(w, http.StatusInternalServerError, ErrorResponse{err.Error()})
		}
		return
	}
	if err := h.js.Send(w, http.StatusOK, map[string]int{"userID": userID}); err != nil {
		h.lg.Errorln(err.Error())
	}
	h.lg.Println("signup - ok")
}
