package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ck3g/newnews-api/data"
)

type createUserRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handlers) UsersCreate(w http.ResponseWriter, r *http.Request) {
	var req createUserRequestBody
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		env := envelope{"error": "Bad Request"}
		h.writeJSON(w, http.StatusBadRequest, env, nil)
		return
	}

	// TODO: add validation

	_, err = h.Models.Users.Create(req.Username, req.Password)
	if err != nil {
		env := envelope{"error": createUserErrorMessage(err)}
		h.writeJSON(w, http.StatusUnprocessableEntity, env, nil)
		return
	}

	env := envelope{"token": "fake-token"}

	h.writeJSON(w, http.StatusCreated, env, nil)
}

func createUserErrorMessage(err error) string {
	if errors.Is(err, data.ErrUserExists) {
		return "User already exists"
	}

	return "Something went wrong"
}
