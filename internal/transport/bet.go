package handler

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/Levap123/adverts/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

type BetRequest struct {
	Bet int `json:"bet,omitempty"`
}

func (h *Handler) MakeBet(w http.ResponseWriter, r *http.Request) {
	var input BetRequest
	if err := h.js.Read(r, &input); err != nil {
		if err := h.js.Send(w, http.StatusBadRequest, ErrorResponse{"request body is invalid"}); err != nil {
			h.lg.Errorln(err)
		}
		return
	}
	userId := r.Context().Value("userId")
	advertId, err := strconv.Atoi(chi.URLParam(r, "advertId"))
	if err != nil {
		h.lg.Errorln(err)
		if err := h.js.Send(w, http.StatusNotFound, ErrorResponse{"not found"}); err != nil {
			h.lg.Errorln(err.Error())
		}
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	h.mu.Lock()
	betId, err := h.service.MakeBet(ctx, userId.(int), advertId, input.Bet)
	h.mu.Unlock()
	if err != nil {
		h.lg.Errorln(err)
		switch {
		case errors.Is(err, service.ErrPriceSmall):
			if err := h.js.Send(w, http.StatusBadRequest, ErrorResponse{"your bet is too small"}); err != nil {
				h.lg.Errorln(err.Error())
			}
		case errors.Is(err, service.ErrAdvertIsNotActive):
			if err := h.js.Send(w, http.StatusBadRequest, ErrorResponse{service.ErrAdvertIsNotActive.Error()}); err != nil {
				h.lg.Errorln(err.Error())
			}
		case errors.Is(err, pgx.ErrNoRows):
			if err := h.js.Send(w, http.StatusNotFound, ErrorResponse{"advert not found"}); err != nil {
				h.lg.Errorln(err.Error())
			}
		default:
			if err := h.js.Send(w, http.StatusInternalServerError, ErrorResponse{err.Error()}); err != nil {
				h.lg.Errorln(err.Error())
			}
		}
		return
	}
	if err := h.js.Send(w, http.StatusOK, map[string]int{"bet_id": betId}); err != nil {
		h.lg.Errorln(err.Error())
	}
}
