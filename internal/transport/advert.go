package handler

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/Levap123/adverts/internal/service"
)

type AdvertRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Price int    `json:"price"`
}

func (h *Handler) createAdvert(w http.ResponseWriter, r *http.Request) {
	var input AdvertRequest
	if err := h.js.Read(r, &input); err != nil {
		if err := h.js.Send(w, http.StatusBadRequest, ErrorResponse{"request body is invalid"}); err != nil {
			h.lg.Errorln(err)
		}
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	if !h.v.IsTitleValid(input.Title) {
		if err := h.js.Send(w, http.StatusBadRequest, ErrorResponse{ErrInvalidTitle.Error()}); err != nil {
			h.lg.Errorln(err.Error())
		}
		return
	}

	if !h.v.IsBodyValid(input.Body) {
		if err := h.js.Send(w, http.StatusBadRequest, ErrorResponse{ErrInvalidBody.Error()}); err != nil {
			h.lg.Errorln(err.Error())
		}
		return
	}

	if !h.v.IsPriceValid(input.Price) {
		if err := h.js.Send(w, http.StatusBadRequest, ErrorResponse{"price too much"}); err != nil {
			h.lg.Errorln(err.Error())
		}
		return
	}

	userId := r.Context().Value("userId")
	advertId, err := h.service.AdvertService.Create(ctx, input.Title, input.Body, input.Price, userId.(int))
	if err != nil {
		h.lg.Errorln(err)
		switch {
		case errors.Is(err, service.ErrInorrectTitle):
			if err := h.js.Send(w, http.StatusBadRequest, ErrorResponse{err.Error()}); err != nil {
				h.lg.Errorln(err.Error())
			}
		default:
			if err := h.js.Send(w, http.StatusInternalServerError, ErrorResponse{err.Error()}); err != nil {
				h.lg.Errorln(err.Error())
			}
		}
		return
	}
	if err := h.js.Send(w, http.StatusOK, map[string]int{"advertId": advertId}); err != nil {
		h.lg.Errorln(err.Error())
	}
}

func (h *Handler) getAllAdverts(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userId")
}
