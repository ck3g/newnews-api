package data

import (
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

func New(databasePool *pgx.Conn) Models {
	db = databasePool

	return Models{
		Items: &ItemsModel{DB: db},
	}
}
