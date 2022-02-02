package data

import (
	"errors"
	"time"

	"github.com/jackc/pgx/v4"
)

var db *pgx.Conn

var (
	ErrUserExists = errors.New("data: user already exists")
)

type Models struct {
	Items ItemsDatastorage
	Users UsersDatastorage
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

type UsersDatastorage interface {
	Create(username, password string) (int64, error)
	Find(id int64) (*User, error)
	FindByUsername(username string) (*User, error)
	Exists(username string) bool
}

type User struct {
	ID             int64     `db:"id" json:"id"`
	Username       string    `db:"username" json:"username"`
	Email          []byte    `db:"email" json:"-"`
	HashedPassword []byte    `db:"hashed_password" json:"-"`
	Karma          int       `db:"karma" json:"karma"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
}

func New(databasePool *pgx.Conn) Models {
	db = databasePool

	return Models{
		Items: &ItemModel{DB: db},
		Users: &UserModel{DB: db},
	}
}

func NewMock() Models {
	return Models{
		Items: &MockItemModel{},
		Users: &MockUserModel{},
	}
}
