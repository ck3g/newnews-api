package mockdb

import (
	"errors"
	"strings"
	"time"

	"github.com/ck3g/newnews-api/data"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct{}

var users = []*data.User{}
var userLastID int64 = 0

func (m *UserModel) Create(username, password string) (int64, error) {
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

func (m *UserModel) Find(id int64) (*data.User, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}

	return nil, errors.New("user not found")
}

func (m *UserModel) FindByUsername(username string) (*data.User, error) {
	for _, user := range users {
		if strings.EqualFold(user.Username, username) {
			return user, nil
		}
	}

	return nil, errors.New("user not found")
}

func (m *UserModel) Exists(username string) bool {
	_, err := m.FindByUsername(username)

	return err == nil
}

func (m *UserModel) Truncate() {
	users = []*data.User{}
	userLastID = 0
}
