package router

import (
	"shortener/controller"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Shortener() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Compress(1))
	router.Use(middleware.Timeout(60 * time.Second))

	// Health check endpoint
	router.Get("/health-check", controller.HealthCheckHandler)

	// Shortener endpoints
	router.Post("/shorten", controller.ShortenHandler)
	router.Get("/{slug}", controller.ShortURLHandler)

	return router
}
