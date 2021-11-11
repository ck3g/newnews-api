package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"Healthy"}`))
	})
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"items":[{"title":"Google", "link":"https://google.com"}, {"title":"Apple", "link":"https://apple.com"}]}`))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	http.ListenAndServe(":"+port, r)
}
