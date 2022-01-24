package data

import (
	"time"

	"github.com/jackc/pgx/v4"
)

var db *pgx.Conn

type Models struct {
	Items ItemsDatastorage
}

type ItemsDatastorage interface {
	AllNew() ([]*Item, error)
	Create(item Item) (int64, error)
	Find(id int64) (*Item, error)
	Destroy(id int64)
}

type Item struct {
	ID        int64     `db:"id" json:"id"`
	Title     string    `db:"title" json:"title"`
	Link      string    `db:"link" json:"link"`
	FromSite  string    `db:"from_site" json:"from_site"`
	Points    int       `db:"points" json:"points"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func New(databasePool *pgx.Conn) Models {
	db = databasePool

	return Models{
		Items: &ItemsModel{DB: db},
	}
}
