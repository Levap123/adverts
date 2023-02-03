package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) InitRoutes() http.Handler {
	r := chi.NewRouter()

	r.Route("/auth", func(auth chi.Router) {
		auth.Post("/sign-up", h.signUp)
	})
	return r
}
