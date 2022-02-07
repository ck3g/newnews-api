package mockdb

import (
	"errors"
	"strings"
	"time"

	"github.com/ck3g/newnews-api/data"
	"golang.org/x/crypto/bcrypt"
)

type MockUserModel struct{}

var users = []*data.User{}
var userLastID int64 = 0

func (m *MockUserModel) Create(username, password string) (int64, error) {
	if m.Exists(username) {
		return 0, data.ErrUserExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	user := &data.User{
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

func (m *MockUserModel) Find(id int64) (*data.User, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}

	return nil, errors.New("user not found")
}

func (m *MockUserModel) FindByUsername(username string) (*data.User, error) {
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
	users = []*data.User{}
	userLastID = 0
}
