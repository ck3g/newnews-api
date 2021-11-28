package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	h := Handlers{}
	handler := http.HandlerFunc(h.Health)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("wrong status code: want %v; got %v", http.StatusOK, status)
	}

	wantBody := `{"status":"Healthy"}`
	gotBody := rr.Body.String()
	if wantBody != gotBody {
		t.Errorf("wrong body: want %v; got %v", wantBody, gotBody)
	}
}

func TestHomeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	h := Handlers{}
	handler := http.HandlerFunc(h.Home)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("wrong status code: want %v; got %v", http.StatusOK, status)
	}

	wantBody := `{"items":[{"title":"Google", "link":"https://google.com"}, {"title":"Apple", "link":"https://apple.com"}]}`
	gotBody := rr.Body.String()
	if wantBody != gotBody {
		t.Errorf("wrong body: want %v; got %v", wantBody, gotBody)
	}
}
