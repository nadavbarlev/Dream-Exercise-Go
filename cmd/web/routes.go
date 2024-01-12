package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nadavbarlev/dream-exercise-go/internal/handlers"
)

func routes() http.Handler {
	mux := chi.NewRouter()
	mux.Post("/location", handlers.LocationHandler)
	return mux
}