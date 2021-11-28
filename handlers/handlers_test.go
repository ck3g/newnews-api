package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ck3g/newnews-api/pkg/jsonh"
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

	wantBody := []byte(`{ "status": "Healthy" }`)
	gotBody := rr.Body.Bytes()
	if !jsonh.Equal(wantBody, gotBody) {
		t.Errorf("wrong body: want %s; got %s", wantBody, gotBody)
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

	wantBody := []byte(`
		{
			"items": [
				{ "title": "Google", "link": "https://google.com" },
				{ "title": "Apple", "link": "https://apple.com" }
			]
		}`)
	gotBody := rr.Body.Bytes()
	if !jsonh.Equal(wantBody, gotBody) {
		t.Errorf("wrong body: want %s; got %s", wantBody, gotBody)
	}
}
