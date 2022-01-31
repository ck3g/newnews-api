package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ck3g/newnews-api/data"
)

func TestUsers_Create(t *testing.T) {
	tests := []struct {
		name         string
		username     string
		password     string
		wantStatus   int
		wantError    bool
		errorMessage string
		wantToken    string
	}{
		{
			"successful create",
			"user@example.com",
			"password",
			http.StatusCreated,
			false,
			"",
			"fake-token",
		},
		{
			"existing user",
			"exists@example.com",
			"password",
			http.StatusUnprocessableEntity,
			true,
			"User already exists",
			"",
		},
	}

	h := Handlers{
		Models: data.NewMock(),
	}
	handler := http.HandlerFunc(h.UsersCreate)

	type createResponse struct {
		Error string `json:"error"`
		Token string `json:"token"`
	}

	var resp createResponse

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()

			mockUserModel := data.MockItemsModel{}
			mockUserModel.Truncate()
			h.Models.Users.Create("exists@example.com", "password")

			body := []byte(fmt.Sprintf(`{"username":"%s","password":"%s"}`, tt.username, tt.password))
			req, err := http.NewRequest("POST", "/users/create", bytes.NewBuffer(body))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/json")

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.wantStatus {
				t.Errorf("wrong status code: want %v; got %v", tt.wantStatus, status)
			}

			if !h.Models.Users.Exists(tt.username) {
				t.Error("expected user to be existed")
			}

			err = json.Unmarshal(rr.Body.Bytes(), &resp)
			if err != nil {
				t.Fatal(err)
			}

			if tt.wantError {
				if resp.Error != tt.errorMessage {
					t.Errorf("expected to have error '%s'; got '%s'", tt.errorMessage, resp.Error)
				}
			} else {
				if resp.Token != tt.wantToken {
					t.Errorf("wrong token returned; want %s; got %s", tt.wantToken, resp.Token)
				}
			}
		})
	}
}
