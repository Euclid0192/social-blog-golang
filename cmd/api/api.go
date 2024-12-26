package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Euclid0192/social-blog-golang/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Application interface for later injection
type application struct {
	config config
	store  store.Storage
}

type config struct {
	addr string
	db     dbConfig
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

func (app *application) mount() http.Handler {
	// /// Kinda router
	// mux := http.NewServeMux()

	// /// Good practice to include version number
	// mux.HandleFunc("/v1/health", app.healthCheckHandler)

	// Better and easier to not use std lib and use external package
	// In this case, chi: lightweight router for HTTP services
	r := chi.NewRouter()

	// Middlewares
	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	// Route group
	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
	})

	return r
}

func (app *application) run(mux http.Handler) error {
	/// Start server
	srv := &http.Server{
		Addr:    app.config.addr,
		Handler: mux,
		/// Add read/write, idle timeout
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("server has started at %s", app.config.addr)

	return srv.ListenAndServe()
}
