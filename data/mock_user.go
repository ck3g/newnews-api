package data

import (
	"errors"
	"strings"
	"time"
)

type MockUserModel struct{}

var users = []*User{}
var userLastID int64 = 0

func (m *MockUserModel) Create(username, password string) (int64, error) {
	hashedPassword := []byte(password)

	user := &User{
		ID:             userLastID + 1,
		Username:       username,
		HashedPassword: hashedPassword,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	users = append(users, user)
	userLastID++

	return user.ID, nil
}

func (m *MockUserModel) Find(id int64) (*User, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}

	return nil, errors.New("user not found")
}

func (m *MockUserModel) FindByUsername(username string) (*User, error) {
	for _, user := range users {
		if strings.EqualFold(user.Username, username) {
			return user, nil
		}
	}

	return nil, errors.New("user not found")
}

func (m *MockUserModel) Exists(username string) bool {
	_, err := m.FindByUsername(username)

	return err == nil
}

func (m *MockUserModel) Truncate() {
	users = []*User{}
	userLastID = 0
}
