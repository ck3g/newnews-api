package data

import "testing"

func TestUser_Create(t *testing.T) {
	db := newTestDB(t)
	model := UserModel{DB: db}

	id, err := model.Create("username", "password")
	if err != nil {
		t.Error("expected no error, but got one: ", err)
	}

	user, err := model.FindByUsername("username")
	if err != nil {
		t.Error("Expected to find a user, but it wasn't found: ", err)
	}

	if user.ID != id {
		t.Errorf("wrong user ID returned; want %d; got %d", id, user.ID)
	}

	_, err = model.Create("username", "password")
	if err == nil {
		t.Error("expected to have an error, got no errors")
	}

	_, err = model.Create(" username ", "password")
	if err == nil {
		t.Error("expected to have an error, got no errors")
	}
}

func TestUser_Find(t *testing.T) {
	db := newTestDB(t)
	model := UserModel{DB: db}

	id, _ := model.Create("user", "password")
	user, err := model.Find(id)
	if err != nil {
		t.Error("expected to find a user, but it wasn't found")
	}

	if user.ID != id {
		t.Errorf("wrong user returned; want ID %d; got %d", id, user.ID)
	}

	_, err = model.Find(-1)
	if err == nil {
		t.Error("expected not to find a user, but found one")
	}
}

func TestUser_FindByUsername(t *testing.T) {
	db := newTestDB(t)
	model := UserModel{DB: db}

	model.Create("user", "password")

	_, err := model.FindByUsername("User")
	if err != nil {
		t.Error("expected to find a user, but it wasn't found")
	}

	_, err = model.FindByUsername("unknown")
	if err == nil {
		t.Error("expected not to find a user, but found one")
	}
}

func TestUser_Exists(t *testing.T) {
	db := newTestDB(t)
	model := UserModel{DB: db}

	exists := model.Exists("User")
	if exists {
		t.Error("expected user to not exist, but it exists")
	}

	model.Create("user", "password")

	exists = model.Exists("User")
	if !exists {
		t.Error("expected user to exist, but it doesn't")
	}

}
