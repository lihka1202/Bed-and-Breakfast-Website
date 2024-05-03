package main

import (
	"basicgotesting/pkg/config"
	"basicgotesting/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	//! In the event of panics, does a better job at handling them and gracefully letting it down
	mux.Use(middleware.Recoverer)

	//! Custom middleware
	mux.Use(NoSurf)

	//! Use Sessions middleware
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
