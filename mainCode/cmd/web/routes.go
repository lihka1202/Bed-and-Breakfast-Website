package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lihka1202/mainCode/pkg/config"
	"github.com/lihka1202/mainCode/pkg/handlers"
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
