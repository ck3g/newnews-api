package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ck3g/newnews-api/data"
	"github.com/ck3g/newnews-api/data/mockdb"
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
	h := Handlers{
		Models: mockdb.NewMock(),
	}
	item1 := data.Item{
		Title:    "Google",
		Link:     "https://google.com",
		FromSite: "google.com",
		Points:   10,
	}
	item2 := data.Item{
		Title:    "Apple",
		Link:     "https://apple.com",
		FromSite: "apple.com",
		Points:   20,
	}
	h.Models.Items.Create(item1)
	h.Models.Items.Create(item2)
	handler := http.HandlerFunc(h.Home)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("wrong status code: want %v; got %v", http.StatusOK, status)
	}

	type itemResponse struct {
		Title    string `json:"title"`
		Link     string `json:"link"`
		FromSite string `json:"from_site"`
		Points   int    `json:"points"`
	}

	type homeResponse struct {
		Items []itemResponse `json:"items"`
	}

	var resp homeResponse

	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		t.Fatal(err)
	}

	if len(resp.Items) != 2 {
		t.Errorf("wrong items number returned; want %d; got %d", 2, len(resp.Items))
	}

	expected := homeResponse{
		Items: []itemResponse{
			{
				Title:    item1.Title,
				Link:     item1.Link,
				FromSite: item1.FromSite,
				Points:   item1.Points,
			},
			{
				Title:    item2.Title,
				Link:     item2.Link,
				FromSite: item2.FromSite,
				Points:   item2.Points,
			},
		},
	}

	if expected.Items[0] != resp.Items[0] || expected.Items[1] != resp.Items[1] {
		t.Errorf("wrong items returned")
	}

}
