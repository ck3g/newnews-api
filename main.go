package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/ck3g/newnews-api/handlers"
)

type application struct {
	Handlers *handlers.Handlers
}

func main() {
	app := &application{
		Handlers: &handlers.Handlers{},
	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/health", app.Handlers.Health)
	r.Get("/", app.Handlers.Home)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	http.ListenAndServe(":"+port, r)
}
