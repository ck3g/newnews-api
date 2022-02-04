package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ck3g/newnews-api/data"
	"github.com/ck3g/newnews-api/pkg/validator"
)

type createUserRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type responseError struct {
	Message []string `json:"message"`
}

func (h *Handlers) UsersCreate(w http.ResponseWriter, r *http.Request) {
	var req createUserRequestBody
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		env := envelope{"errors": []responseError{
			{Message: []string{"Bad request"}},
		}}
		h.writeJSON(w, http.StatusBadRequest, env, nil)
		return
	}

	v := validator.New()
	v.ValidatePresenseOf("username", req.Username)
	v.ValidateLengthOf("username", req.Username, 3, 20)
	v.ValidatePresenseOf("password", req.Password)
	v.ValidateLengthOf("password", req.Password, 6, 128)

	if !v.Valid() {
		env := envelope{"errors": v.ErrorMessages()}
		h.writeJSON(w, http.StatusUnprocessableEntity, env, nil)
		return
	}

	_, err = h.Models.Users.Create(req.Username, req.Password)
	if err != nil {
		env := envelope{"errors": createUserErrorMessage(err)}
		h.writeJSON(w, http.StatusUnprocessableEntity, env, nil)
		return
	}

	env := envelope{"token": "fake-token"}

	h.writeJSON(w, http.StatusCreated, env, nil)
}

func createUserErrorMessage(err error) []responseError {
	if errors.Is(err, data.ErrUserExists) {
		return []responseError{
			{Message: []string{"user already exists"}},
		}
	}

	return []responseError{
		{Message: []string{"Something went wrong"}},
	}
}
