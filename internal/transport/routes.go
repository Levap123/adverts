package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) InitRoutes() http.Handler {
	r := chi.NewRouter()

	r.Route("/auth", func(auth chi.Router) {
		auth.Post("/sign-up", h.signUp)
		auth.Post("/sign-in", h.signIn)
	})

	r.Route("/adverts", func(adverts chi.Router) {
		adverts.Use(h.userIdentity)
		adverts.Post("/", h.createAdvert)
		adverts.Get("/", h.getAllAdverts)
	})
	return r
}
