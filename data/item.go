package data

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4"
)

type Item struct {
	ID        int       `db:"id" json:"id"`
	Title     string    `db:"title" json:"title"`
	Link      string    `db:"link" json:"link"`
	FromSite  string    `db:"from_site" json:"from_site"`
	Points    int       `db:"points" json:"points"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func (i *Item) GetAllNew() ([]*Item, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	rows, _ := conn.Query(context.Background(), "SELECT id, title, link, from_site, points, created_at, updated_at FROM items")

	all := []*Item{}

	for rows.Next() {
		var i Item
		err := rows.Scan(&i.ID, &i.Title, &i.Link, &i.FromSite, &i.Points, &i.CreatedAt, &i.UpdatedAt)
		if err != nil {
			return nil, err
		}

		all = append(all, &i)
	}

	return all, nil
}
