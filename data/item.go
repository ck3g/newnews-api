package data

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
)

type ItemsModel struct {
	DB *pgx.Conn
}

type Item struct {
	ID        int       `db:"id" json:"id"`
	Title     string    `db:"title" json:"title"`
	Link      string    `db:"link" json:"link"`
	FromSite  string    `db:"from_site" json:"from_site"`
	Points    int       `db:"points" json:"points"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func (i *ItemsModel) AllNew() ([]*Item, error) {
	rows, _ := i.DB.Query(context.Background(), "SELECT id, title, link, from_site, points, created_at, updated_at FROM items")

	all := []*Item{}

	for rows.Next() {
		var item Item
		err := rows.Scan(
			&item.ID, &item.Title, &item.Link, &item.FromSite, &item.Points, &item.CreatedAt, &item.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		all = append(all, &item)
	}

	return all, nil
}
