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

// TODO: move to handlers/helpers (or handlers/errors)
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
	validateUsernameAndPassword(v, req.Username, req.Password)

	if !v.Valid() {
		env := envelope{"errors": v.ErrorMessages()}
		h.writeJSON(w, http.StatusUnprocessableEntity, env, nil)
		return
	}

	userID, err := h.Models.Users.Create(req.Username, req.Password)
	if err != nil {
		env := envelope{"errors": createUserErrorMessage(err)}
		h.writeJSON(w, http.StatusUnprocessableEntity, env, nil)
		return
	}

	token, err := h.Models.AuthSessions.Authenticate(userID)
	if err != nil {
		env := envelope{"errors": createUserErrorMessage(err)}
		h.writeJSON(w, http.StatusInternalServerError, env, nil)
		return
	}

	env := envelope{"token": token}

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
