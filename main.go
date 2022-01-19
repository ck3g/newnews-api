package main

import (
	"context"
	"net/http"
	"os"

	"log"

	"github.com/ck3g/newnews-api/data"
	"github.com/ck3g/newnews-api/handlers"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

type application struct {
	Handlers *handlers.Handlers
	Models   data.Models
}

func main() {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	db, err := dbConnect()
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close(context.Background())

	models := data.New(db)
	app := &application{
		Handlers: &handlers.Handlers{
			Models: models,
		},
		Models: models,
	}

	port := os.Getenv("PORT")

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: app.routes(),
	}

	err = srv.ListenAndServe()
	srv.ErrorLog.Fatal(err)
}

func dbConnect() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	return conn, nil
}
