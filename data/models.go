package data

import (
	"github.com/jackc/pgx/v4"
)

var db *pgx.Conn

type Models struct {
	Items ItemsModel
}

func New(databasePool *pgx.Conn) Models {
	db = databasePool

	return Models{
		Items: ItemsModel{DB: db},
	}
}
