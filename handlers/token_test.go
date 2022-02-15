package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ck3g/newnews-api/data/mockdb"
)

func TestTokens_Create(t *testing.T) {
	tests := []struct {
		name       string
		username   string
		password   string
		wantStatus int
		wantError  bool
		wantBody   []byte
	}{
		{
			"successful login",
			"user@example.com",
			"password",
			http.StatusCreated,
			false,
			[]byte(`{"token":"fake-token"}`),
		},
		{
			"invalid username",
			"unknown@example.com",
			"password",
			http.StatusUnprocessableEntity,
			true,
			[]byte(`{"errors":[{"message":["Invalid username or password"]}]}`),
		},
		{
			"invalid password",
			"user@example.com",
			"invalid-password",
			http.StatusUnprocessableEntity,
			true,
			[]byte(`{"errors":[{"message":["Invalid username or password"]}]}`),
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
	handler := http.HandlerFunc(h.TokenCreate)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()

			mockUserModel := mockdb.UserModel{}
			mockUserModel.Truncate()
			_, err := mockUserModel.Create("user@example.com", "password")
			if err != nil {
				t.Fatal(err)
			}

			body := []byte(fmt.Sprintf(`{"username":"%s","password":"%s"}`, tt.username, tt.password))
			req, err := http.NewRequest("POST", "/token", bytes.NewBuffer(body))
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
		})
	}
}
