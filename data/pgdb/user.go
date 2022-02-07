package pgdb

import (
	"context"
	"errors"
	"strings"

	"github.com/ck3g/newnews-api/data"
	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	DB *pgx.Conn
}

func (m *UserModel) Create(username, password string) (int64, error) {
	var id int64

	username = strings.Trim(username, " ")

	if m.Exists(username) {
		return id, data.ErrUserExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return id, err
	}

	query := "INSERT INTO users (username, hashed_password) VALUES ($1, $2) RETURNING id"
	err = m.DB.QueryRow(context.Background(), query, username, hashedPassword).Scan(&id)
	if err != nil {
		return id, errors.New("error creating user")
	}

	return id, nil
}

func (m *UserModel) Find(id int64) (*data.User, error) {
	var u data.User

	query := "SELECT id, username, email, hashed_password, karma, created_at, updated_at FROM users WHERE id = $1"
	err := m.DB.QueryRow(context.Background(), query, id).Scan(
		&u.ID, &u.Username, &u.Email, &u.HashedPassword, &u.Karma, &u.CreatedAt, &u.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (m *UserModel) FindByUsername(username string) (*data.User, error) {
	var u data.User

	query := "SELECT id, username, email, hashed_password, karma, created_at, updated_at FROM users WHERE LOWER(username) = $1"
	err := m.DB.QueryRow(context.Background(), query, strings.ToLower(username)).Scan(
		&u.ID, &u.Username, &u.Email, &u.HashedPassword, &u.Karma, &u.CreatedAt, &u.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (m *UserModel) Exists(username string) bool {
	var count int64

	query := "SELECT COUNT(*) FROM users WHERE LOWER(username) = $1"
	err := m.DB.QueryRow(context.Background(), query, strings.ToLower(username)).Scan(&count)
	if err != nil {
		return false
	}

	return count != 0
}
