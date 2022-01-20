package handlers

import (
	"net/http"

	"github.com/ck3g/newnews-api/data"
)

type Handlers struct {
	Models data.Models
}

func (h *Handlers) Health(w http.ResponseWriter, r *http.Request) {
	env := envelope{
		"status": "Healthy",
	}

	h.writeJSON(w, http.StatusOK, env, nil)
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	items, err := h.Models.Items.AllNew()
	if err != nil {
		env := envelope{
			"error": "Internal Server Error",
		}
		h.writeJSON(w, http.StatusInternalServerError, env, nil)
	}

	env := envelope{
		"items": items,
	}

	h.writeJSON(w, http.StatusOK, env, nil)
}
