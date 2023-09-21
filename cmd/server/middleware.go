package server

import (
	"context"
	"net/http"
	"strings"

	"github.com/hyperversalblocks/averveil/pkg/jwt"
)

func (c *Container) ValidateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if len(authHeader) == 0 {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		info, err := c.jwt.ValidateToken(token, c.config.JWTSecret)
		if err != nil {
			http.Error(w, "token expired", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), jwt.PublicKeyUser, info.PublicKey)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (c *Container) ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				c.logger.Error(err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
