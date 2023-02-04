package handler

import (
	"context"
	"net/http"
	"time"
)

type AdvertRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Price int    `json:"price"`
}

func (h *Handler) createAdvert(w http.ResponseWriter, r *http.Request) {
	var input AdvertRequest
	if err := h.js.Read(r, &input); err != nil {
		if err := h.js.Send(w, http.StatusBadRequest, ErrorResponse{err.Error()}); err != nil {
			h.lg.Errorln(err)
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	if !h.v.IsAdvertValid(input.Title, input.Body, input.Price) {
		if err := h.js.Send(w, http.StatusBadRequest, ErrorResponse{"advert is invalid"}); err != nil {
			h.lg.Errorln(err.Error())
		}
		return
	}
	advertId, err := h.service.AdvertService.Create(ctx, input.Title, input.Body, input.Price)
}
