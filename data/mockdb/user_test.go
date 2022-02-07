package mockdb

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestMockUser_Create(t *testing.T) {
	model := UserModel{}
	model.Truncate()

	id, err := model.Create("user", "password")
	if err != nil {
		t.Errorf("expected to create a user, but error returned")
	}

	user, err := model.Find(id)
	if err != nil {
		t.Error("user not found but it should")
	}

	if user.Username != "user" {
		t.Errorf("user has wrong username; want %s; got %s", "user", user.Username)
	}

	err = bcrypt.CompareHashAndPassword(user.HashedPassword, []byte("password"))
	if err != nil {
		t.Error("expected password to be properly hashed")
	}

	_, err = model.Create("user", "password")
	if err == nil {
		t.Error("expected to get an error when creating a user with existing username, got nothing")
	}
}

func TestMockUser_Find(t *testing.T) {
	model := UserModel{}
	model.Truncate()

	id, _ := model.Create("user", "password")

	user, err := model.Find(id)
	if err != nil {
		t.Error("expected to find a user but it wasn't found")
	}

	if user.ID != 1 {
		t.Errorf("wrong user found; want ID: %d; got %d", 1, user.ID)
	}

	_, err = model.Find(-1)
	if err == nil {
		t.Error("expected not to find a user, but it was found")
	}
}

func TestMockUser_FindByUsername(t *testing.T) {
	model := UserModel{}
	model.Truncate()

	id, _ := model.Create("UserName", "password")

	user, err := model.FindByUsername("userNAME")
	if err != nil {
		t.Error("expected to find a user, but it wasn't found")
	}

	if user.ID != id {
		t.Errorf("wrong user found; want ID %d; got %d", id, user.ID)
	}

	_, err = model.FindByUsername("unknown")
	if err == nil {
		t.Error("expected not to find a user, but it was found")
	}
}

func TestMockUser_Exists(t *testing.T) {
	model := UserModel{}
	model.Truncate()

	model.Create("UserName", "password")

	if !model.Exists("userNAME") {
		t.Error("expected to find a user, but it wasn't found")
	}

	if model.Exists("unknown") {
		t.Error("expected not to find a user, but it was found")
	}
}
