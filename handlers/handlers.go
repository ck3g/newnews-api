package handlers

import "net/http"

type Handlers struct {
}

func (h *Handlers) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"Healthy"}`))
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"items":[{"title":"Google", "link":"https://google.com"}, {"title":"Apple", "link":"https://apple.com"}]}`))
}
