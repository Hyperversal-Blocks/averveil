package api

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func (s *Services) Cors() {
	s.router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           int(12 * time.Hour),
	}))
}

func (s *Services) Routes() {
	s.router.Route("/balance", func(r chi.Router) {
		s.HBLOCKAccessHandler()
		r.Get("/get", s.user.GetBalance)
	})

	s.router.Route("/upload", func(r chi.Router) {
		r.Post("/csv", s.upload.CSV)
	})

	s.router.Route("/view", func(r chi.Router) {
		r.Get("/csv", s.view.CSV)
	})
}

func (s *Services) GetRouter() *chi.Mux {
	return s.router
}
