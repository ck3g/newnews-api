package data

import (
	"errors"
	"time"
)

var (
	ErrUserExists       = errors.New("data: user already exists")
	ErrUserDoesNotExist = errors.New("data: user does not exist")
)

type Models struct {
	Items        ItemsDataStorage
	Users        UsersDataStorage
	AuthSessions AuthSessionDataStorage
}

type ItemsDataStorage interface {
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

type UsersDataStorage interface {
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

type AuthSessionDataStorage interface {
	GenerateForUserID(id int64) (string, error)
}

type AuthSession struct {
	ID        int64     `db:"id" json:"id"`
	Token     string    `db:"token" json:"token"`
	UserID    int64     `db:"user_id" json:"user_id"`
	ExpiredAt time.Time `db:"expired_at" json:"expired_at"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
