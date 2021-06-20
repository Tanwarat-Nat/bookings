package main

import (
	"net/http"

	"github.com/Tanwarat-Nat/bookings/pkg/config"
	"github.com/Tanwarat-Nat/bookings/pkg/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/home", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
