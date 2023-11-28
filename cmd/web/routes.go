package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

// Registering a Request Handler
func (a *application) routes() http.Handler {

	//create a new router
	mux := chi.NewRouter()

	// A good base middleware stack
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Recoverer)

	if a.debug {
		mux.Use(middleware.Logger)
	}

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(a.appName))
	})

	mux.Get("/api/comments", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("dit me son linh"))
	})
	return mux
}
