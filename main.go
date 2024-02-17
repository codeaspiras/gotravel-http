package main

import (
	"net/http"
	"os"
	"time"

	myhttp "github.com/codeaspiras/gotravel-http/internal/http"
	"github.com/codeaspiras/gotravel-http/internal/middlewares"
	"github.com/codeaspiras/gotravel-http/internal/validator"
	"github.com/codeaspiras/gotravel/stdio"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	// dependencies
	io := stdio.New()

	// run
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "8080"
	}

	io.Echo("# Inicializando HTTP server na porta %s", port)

	validator := validator.New()

	router := chi.NewRouter()

	// midlewares
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	router.Use(middleware.Timeout(5 * time.Second))
	router.Use(middlewares.WithIO(io))
	router.Use(middlewares.WithValidator(validator))

	// routes
	myhttp.NewRouter(router)

	// serve
	if err := http.ListenAndServe(":"+port, router); err != nil {
		io.Echo("# Um erro aconteceu: %s\n", err)
	}

	io.Echo("# Servidor encerrado")
}
