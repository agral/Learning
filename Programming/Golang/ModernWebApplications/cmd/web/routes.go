package main

import (
	"ModernWebApplications/pkg/config"
	"ModernWebApplications/pkg/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
