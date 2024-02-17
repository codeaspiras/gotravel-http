package http

import (
	"github.com/codeaspiras/gotravel-http/internal/http/handlers"

	"github.com/go-chi/chi/v5"
)

func NewRouter(r *chi.Mux) {
	r.Route("/api/v1/", func(r chi.Router) {
		r.Post("/calculate", handlers.Calculate)
	})
}
