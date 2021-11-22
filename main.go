package main

import (
	"net/http"
	"os"

	"github.com/ck3g/newnews-api/handlers"
)

type application struct {
	Handlers *handlers.Handlers
}

func main() {
	app := &application{
		Handlers: &handlers.Handlers{},
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	srv.ErrorLog.Fatal(err)
}
