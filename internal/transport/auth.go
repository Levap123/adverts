package handler

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/Levap123/adverts/internal/service"
	"github.com/jackc/pgx/v5"
)

type userRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	var input userRequest
	if err := h.js.Read(r, &input); err != nil {
		h.lg.Errorln(err.Error())
		if err := h.js.Send(w, http.StatusBadRequest, ErrorResponse{"request body is invalid"}); err != nil {
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

	userID, err := h.service.AuthService.Create(ctx, input.Email, input.Password)
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

type Tokens struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	var input userRequest
	if err := h.js.Read(r, &input); err != nil {
		h.lg.Errorln(err.Error())
		if err := h.js.Send(w, http.StatusBadRequest, ErrorResponse{"request body is invalid"}); err != nil {
			h.lg.Errorln(err.Error())
		}
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	accessToken, refreshToken, err := h.service.AuthService.GetTokens(ctx, input.Email, input.Password)
	if err != nil {
		h.lg.Errorln(err.Error())
		switch {
		case errors.Is(err, service.ErrInvalidPassword):
			if err := h.js.Send(w, http.StatusBadRequest, ErrorResponse{err.Error()}); err != nil {
				h.lg.Errorln(err.Error())
			}
		case errors.Is(err, pgx.ErrNoRows):
			if err := h.js.Send(w, http.StatusBadRequest, ErrorResponse{"user with this email does not exist"}); err != nil {
				h.lg.Errorln(err.Error())
			}
		default:
			if err := h.js.Send(w, http.StatusInternalServerError, ErrorResponse{err.Error()}); err != nil {
				h.lg.Errorln(err.Error())
			}
		}
		return
	}

	if err := h.js.Send(w, http.StatusOK, Tokens{Access: accessToken, Refresh: refreshToken}); err != nil {
		h.lg.Errorln(err.Error())
	}

	h.lg.Println("signin - ok")
}
