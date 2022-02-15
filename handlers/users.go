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

func (h *Handlers) UsersCreate(w http.ResponseWriter, r *http.Request) {
	var req createUserRequestBody
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.badRequestResponse(w)
		return
	}

	v := validator.New()
	validateUsernameAndPassword(v, req.Username, req.Password)

	if !v.Valid() {
		h.validationErrorResponse(w, v.ErrorMessages())
		return
	}

	userID, err := h.Models.Users.Create(req.Username, req.Password)
	if err != nil {
		h.validationErrorResponse(w, createUserErrorMessage(err))
		return
	}

	token, err := h.Models.AuthSessions.Authenticate(userID)
	if err != nil {
		h.validationErrorResponse(w, createUserErrorMessage(err))
		return
	}

	env := envelope{"token": token}

	h.writeJSON(w, http.StatusCreated, env, nil)
}

func createUserErrorMessage(err error) []responseError {
	if errors.Is(err, data.ErrUserExists) {
		return responseErrorMessage("user already exists")
	}

	return responseErrorMessage("Something went wrong")
}
