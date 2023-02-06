package handler

import (
	"context"
	"net/http"
	"strings"

	"github.com/Levap123/adverts/internal/service"
	"github.com/Levap123/adverts/pkg/jwt"
)

func (h *Handler) userIdentity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 {
			if err := h.js.Send(w, http.StatusUnauthorized, ErrorResponse{"auth header is invalid"}); err != nil {
				h.lg.Errorln(err.Error())
			}
			return
		}
		id, tokenType, err := jwt.ParseToken(authParts[1])
		if err != nil {
			if err := h.js.Send(w, http.StatusUnauthorized, ErrorResponse{"token is invalid"}); err != nil {
				h.lg.Errorln(err.Error())
			}
			return
		}
		if tokenType != service.AccessType {
			if err := h.js.Send(w, http.StatusUnauthorized, ErrorResponse{"token type is invalid"}); err != nil {
				h.lg.Errorln(err.Error())
			}
			return
		}
		ctx := context.WithValue(r.Context(), "userId", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
