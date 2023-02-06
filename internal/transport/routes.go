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

	r.Route("/api", func(api chi.Router) {
		api.Use(h.userIdentity)

		api.Route("/adverts", func(adverts chi.Router) {
			adverts.Get("/{advertId}", h.getAdvertById)
			adverts.Post("/", h.createAdvert)
			adverts.Get("/", h.getAllAdverts)

			adverts.Route("/{advertId}/bets", func(bets chi.Router) {
				bets.Post("/", h.MakeBet)
			})
		})
	})

	return r
}
