package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/health", app.Handlers.Health)
	r.Get("/", app.Handlers.Home)
	r.Post("/users/create", app.Handlers.UsersCreate)
	r.Post("/token", app.Handlers.TokenCreate)

	return r
}
