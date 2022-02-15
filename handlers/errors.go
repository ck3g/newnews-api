package handlers

import "net/http"

type responseError struct {
	Message []string `json:"message"`
}

func responseErrorMessage(msg string) []responseError {
	return []responseError{
		{Message: []string{msg}},
	}
}

func (h *Handlers) errorResponse(w http.ResponseWriter, status int, message interface{}) {
	env := envelope{
		"errors": message,
	}

	h.writeJSON(w, status, env, nil)
}

func (h *Handlers) badRequestResponse(w http.ResponseWriter) {
	h.errorResponse(w, http.StatusBadRequest, responseErrorMessage("Bad request"))
}

func (h *Handlers) validationErrorResponse(w http.ResponseWriter, message interface{}) {
	h.errorResponse(w, http.StatusUnprocessableEntity, message)
}
