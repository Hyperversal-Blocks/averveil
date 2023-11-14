package api

import (
	"context"
	"net/http"
	"strings"

	"github.com/hyperversal-blocks/averveil/pkg/jwt"
	"github.com/hyperversal-blocks/averveil/pkg/util"
)

func (s *Services) ValidateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if len(authHeader) == 0 {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		info, err := s.jwtService.ValidateToken(token)
		if err != nil {
			http.Error(w, "token expired", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), jwt.PublicKeyUser, info.PublicKey)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (s *Services) ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				s.logger.Error(err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func (s *Services) HBLOCKAccessHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !s.hblockSemaphore.TryAcquire(1) {
			s.logger.Debug("hblock access: simultaneous on-chain operations not supported")
			s.logger.Error(nil, "hblock access: simultaneous on-chain operations not supported")
			util.WriteJson(w, "simultaneous on-chain operations not supported", http.StatusTooManyRequests)
			return
		}
		defer s.hblockSemaphore.Release(1)
	})
}
