package main

import (
	"net/http"
	"os"

	"log"

	"github.com/ck3g/newnews-api/data"
	"github.com/ck3g/newnews-api/handlers"
	"github.com/joho/godotenv"
)

type application struct {
	Handlers *handlers.Handlers
}

func main() {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	models := data.New()
	app := &application{
		Handlers: &handlers.Handlers{
			Models: models,
		},
	}

	port := os.Getenv("PORT")

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	srv.ErrorLog.Fatal(err)
}
