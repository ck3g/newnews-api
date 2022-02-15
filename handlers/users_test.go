package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ck3g/newnews-api/data/mockdb"
	"golang.org/x/crypto/bcrypt"
)

func TestUsers_Create(t *testing.T) {
	tests := []struct {
		name       string
		username   string
		password   string
		wantStatus int
		wantError  bool
		wantBody   []byte
	}{
		{
			"successful create",
			"user@example.com",
			"password",
			http.StatusCreated,
			false,
			[]byte(`{"token":"fake-token"}`),
		},
		{
			"existing user",
			"exists@example.com",
			"password",
			http.StatusUnprocessableEntity,
			true,
			[]byte(`{"errors":[{"message":["user already exists"]}]}`),
		},
		{
			"blank username",
			"",
			"password",
			http.StatusUnprocessableEntity,
			true,
			[]byte(`{"errors":[{"message":["username cannot be blank"]}]}`),
		},
		{
			"blank password",
			"username",
			"",
			http.StatusUnprocessableEntity,
			true,
			[]byte(`{"errors":[{"message":["password cannot be blank"]}]}`),
		},
		{
			"username too short",
			"us",
			"password",
			http.StatusUnprocessableEntity,
			true,
			[]byte(`{"errors":[{"message":["username is too short (minimum is 3 characters)"]}]}`),
		},
		{
			"username too long",
			strings.Repeat("a", 21),
			"password",
			http.StatusUnprocessableEntity,
			true,
			[]byte(`{"errors":[{"message":["username is too long (maximum is 20 characters)"]}]}`),
		},
		{
			"password too short",
			"username",
			"pass",
			http.StatusUnprocessableEntity,
			true,
			[]byte(`{"errors":[{"message":["password is too short (minimum is 6 characters)"]}]}`),
		},
		{
			"password too long",
			"username",
			strings.Repeat("a", 129),
			http.StatusUnprocessableEntity,
			true,
			[]byte(`{"errors":[{"message":["password is too long (maximum is 128 characters)"]}]}`),
		},
	}

	h := Handlers{
		Models: mockdb.New(),
	}
	handler := http.HandlerFunc(h.UsersCreate)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()

			mockUserModel := mockdb.UserModel{}
			mockUserModel.Truncate()
			h.Models.Users.Create("exists@example.com", "password")

			body := []byte(fmt.Sprintf(`{"username":"%s","password":"%s"}`, tt.username, tt.password))
			req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/json")

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.wantStatus {
				t.Errorf("wrong status code: want %v; got %v", tt.wantStatus, status)
			}

			if rr.Body.String() != string(tt.wantBody) {
				t.Errorf("wrong response body; want %s; got %s", tt.wantBody, rr.Body.String())
			}

			if !tt.wantError {
				user, err := h.Models.Users.FindByUsername(tt.username)
				if err != nil {
					t.Error("expected user to be existed")
				}

				err = bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(tt.password))
				if err != nil {
					t.Error("expected hashed password to match the password, but it didn't")
				}
			}
		})
	}
}
