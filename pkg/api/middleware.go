package api

import (
	"context"
	"net/http"
	"strings"

	"github.com/hyperversal-blocks/averveil/pkg/jwt"
)

func (a *Api) ValidateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if len(authHeader) == 0 {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		info, err := a.jwt.ValidateToken(token)
		if err != nil {
			http.Error(w, "token expired", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), jwt.PublicKeyUser, info.PublicKey)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (a *Api) ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				a.logger.Error(err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func (a *Api) HBLOCKAccessHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !a.hblockSemaphore.TryAcquire(1) {
			a.logger.Debug("hblock access: simultaneous on-chain operations not supported")
			a.logger.Error(nil, "hblock access: simultaneous on-chain operations not supported")
			WriteJson(w, "simultaneous on-chain operations not supported", http.StatusTooManyRequests)
			return
		}
		defer a.hblockSemaphore.Release(1)
	})
}
