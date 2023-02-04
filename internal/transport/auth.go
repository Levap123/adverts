package handler

import (
	"context"
	"net/http"
	"strings"
	"time"
)

type signUpBody struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	var input signUpBody
	if err := h.js.Read(r, &input); err != nil {
		h.lg.Errorln(err.Error())
		if err := h.js.Send(w, http.StatusBadRequest, ErrorResponse{err.Error()}); err != nil {
			h.lg.Errorln(err.Error())
		}
		return
	}
	if !h.v.IsEmailValid(input.Email) || !h.v.IsPasswordValid(input.Password) {
		if err := h.js.Send(w, http.StatusBadRequest, ErrorResponse{"password or email is invalid"}); err != nil {
			h.lg.Errorln(err.Error())
		}
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	userID, err := h.service.Create(ctx, input.Email, input.Password)
	if err != nil {
		h.lg.Errorln(err.Error())
		switch {
		case strings.Contains(err.Error(), "duplicate"):
			if err := h.js.Send(w, http.StatusBadRequest, ErrorResponse{"email is already in use"}); err != nil {
				h.lg.Errorln(err.Error())
			}
		default:
			if err := h.js.Send(w, http.StatusInternalServerError, ErrorResponse{err.Error()}); err != nil {
				h.lg.Errorln(err.Error())
			}
		}
		return
	}
	if err := h.js.Send(w, http.StatusOK, map[string]int{"userID": userID}); err != nil {
		h.lg.Errorln(err.Error())
	}
	h.lg.Println("signup - ok")
}
