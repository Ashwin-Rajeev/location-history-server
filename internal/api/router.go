package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// InitRouter initialize a new chi router instance.
func (s *App) InitRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	r.Route("/location/{order_id}", func(r chi.Router) {
		// Basic set of handler routes
		r.Post("/now", s.addLocationHistory)
		r.Get("/", s.getLocationHistory)
		r.Delete("/", s.deleteLocationHistory)
	})

	return r
}
