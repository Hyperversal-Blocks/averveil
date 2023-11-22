package api

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func (a *Api) Cors() {
	a.router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           int(12 * time.Hour),
	}))
}

func (a *Api) Routes() {
	a.router.Route("/balance", func(r chi.Router) {
		a.HBLOCKAccessHandler()
		r.Get("/get", a.user.GetBalance)
	})

	a.router.Route("/upload", func(r chi.Router) {
		r.Post("/csv", a.upload.CSV)
	})

	a.router.Route("/config", func(r chi.Router) {
		r.Get("/", a.user.GetConfig)
	})

	a.router.Route("/view", func(r chi.Router) {
		r.Get("/csv", a.view.CSV)
	})

	a.router.Route("/swarm", func(r chi.Router) {
		r.Get("/health", a.swarm.GetNodeHealth)
	})
}

func (a *Api) GetRouter() *chi.Mux {
	return a.router
}
