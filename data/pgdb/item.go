package pgdb

import (
	"context"

	"github.com/ck3g/newnews-api/data"
	"github.com/jackc/pgx/v4"
)

type ItemModel struct {
	DB *pgx.Conn
}

func (i *ItemModel) AllNew() ([]*data.Item, error) {
	rows, err := i.DB.Query(context.Background(), "SELECT id, title, link, from_site, points, created_at, updated_at FROM items")
	if err != nil {
		return nil, err
	}

	all := []*data.Item{}

	for rows.Next() {
		var item data.Item
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

func (i *ItemModel) Create(item data.Item) (int64, error) {
	var id int64

	query := "INSERT INTO items (title, link, from_site, points) VALUES ($1, $2, $3, $4) RETURNING id"
	args := []interface{}{item.Title, item.Link, item.FromSite, item.Points}
	err := i.DB.QueryRow(context.Background(), query, args...).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (i *ItemModel) Find(id int64) (*data.Item, error) {
	var item data.Item

	query := "SELECT id, title, link, from_site, points, created_at, updated_at FROM items WHERE id=$1"
	err := i.DB.QueryRow(context.Background(), query, id).Scan(
		&item.ID, &item.Title, &item.Link, &item.FromSite, &item.Points, &item.CreatedAt, &item.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (i *ItemModel) Destroy(id int64) {
	i.DB.Exec(context.Background(), "DELETE FROM items WHERE id=$1", id)
}
