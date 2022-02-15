package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ck3g/newnews-api/pkg/validator"
	"golang.org/x/crypto/bcrypt"
)

type createTokenRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handlers) TokenCreate(w http.ResponseWriter, r *http.Request) {
	var req createTokenRequestBody
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

	user, err := h.Models.Users.FindByUsername(req.Username)
	if err != nil {
		h.validationErrorResponse(w, responseErrorMessage("Invalid username or password"))
		return
	}

	err = bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(req.Password))
	if err != nil {
		h.validationErrorResponse(w, responseErrorMessage("Invalid username or password"))
		return
	}

	token, err := h.Models.AuthSessions.Authenticate(user.ID)
	if err != nil {
		h.validationErrorResponse(w, responseErrorMessage("Invalid username or password"))
		return
	}

	env := envelope{"token": token}

	h.writeJSON(w, http.StatusCreated, env, nil)
}
