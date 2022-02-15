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
		env := envelope{"errors": []responseError{
			{Message: []string{"Bad request"}},
		}}
		h.writeJSON(w, http.StatusBadRequest, env, nil)
		return
	}

	v := validator.New()
	v.ValidatePresenceOf("username", req.Username)
	v.ValidateLengthOf("username", req.Username, 3, 20)
	v.ValidatePresenceOf("password", req.Password)
	v.ValidateLengthOf("password", req.Password, 6, 128)

	if !v.Valid() {
		env := envelope{"errors": v.ErrorMessages()}
		h.writeJSON(w, http.StatusUnprocessableEntity, env, nil)
		return
	}

	user, err := h.Models.Users.FindByUsername(req.Username)
	if err != nil {
		env := envelope{"errors": []responseError{
			{Message: []string{"Invalid username or password"}},
		}}
		h.writeJSON(w, http.StatusUnprocessableEntity, env, nil)
		return
	}

	err = bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(req.Password))
	if err != nil {
		env := envelope{"errors": []responseError{
			{Message: []string{"Invalid username or password"}},
		}}
		h.writeJSON(w, http.StatusUnprocessableEntity, env, nil)
		return
	}

	token, err := h.Models.AuthSessions.Authenticate(user.ID)
	if err != nil {
		env := envelope{"errors": []responseError{
			{Message: []string{"Something went wrong"}},
		}}
		h.writeJSON(w, http.StatusInternalServerError, env, nil)
		return
	}

	env := envelope{"token": token}

	h.writeJSON(w, http.StatusCreated, env, nil)
}
