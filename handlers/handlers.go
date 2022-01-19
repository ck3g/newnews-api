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
	items, _ := h.Models.Items.AllNew()
	env := envelope{
		"items": items,
	}

	h.writeJSON(w, http.StatusOK, env, nil)
}
