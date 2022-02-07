package pgdb

import (
	"github.com/ck3g/newnews-api/data"
	"github.com/jackc/pgx/v4"
)

var db *pgx.Conn

func New(databasePool *pgx.Conn) data.Models {
	db = databasePool

	return data.Models{
		Items: &ItemModel{DB: db},
		Users: &UserModel{DB: db},
	}
}
