package data

import "testing"

type TestItem struct {
	Title    string
	Link     string
	FromSite string
	Points   int
}

func TestItem_AllNew(t *testing.T) {
	db := newTestDB(t)

	model := ItemModel{DB: db}

	items, err := model.AllNew()
	if err != nil {
		t.Error("Error fetching all new items")
	}

	if len(items) != 1 {
		t.Errorf("Wrong items count. Want 1; got %d\n", len(items))
	}

	expected := TestItem{
		Title:    "Google",
		Link:     "https://google.com",
		FromSite: "google.com",
		Points:   10,
	}
	actual := TestItem{
		Title:    items[0].Title,
		Link:     items[0].Link,
		FromSite: items[0].FromSite,
		Points:   items[0].Points,
	}

	if expected != actual {
		t.Errorf("Incorrect item returned. Want: %+v; got %+v\n", expected, actual)
	}
}

func TestItem_Create(t *testing.T) {
	db := newTestDB(t)

	model := ItemModel{DB: db}

	item := Item{
		Title:    "Apple",
		Link:     "https://apple.com",
		FromSite: "apple.com",
		Points:   20,
	}

	id, err := model.Create(item)
	if err != nil {
		t.Error("Could not create an item ", err)
	}

	if id == 0 {
		t.Error("Invalid item ID. Want > 0; Got 0;")
	}

	items, err := model.AllNew()
	if err != nil {
		t.Error("Error fetching items")
	}

	if len(items) != 2 {
		t.Errorf("Wrong items count. Want 2; got %d\n", len(items))
	}

	found, err := model.Find(id)
	if err != nil {
		t.Error("Could not find an item ", err)
	}

	expected := TestItem{item.Title, item.Link, item.FromSite, item.Points}
	actual := TestItem{found.Title, found.Link, found.FromSite, found.Points}

	if expected != actual {
		t.Errorf("Incorrect item returned. Want: %+v; got %+v\n", expected, actual)
	}
}

func TestItem_Find(t *testing.T) {
	db := newTestDB(t)

	model := ItemModel{DB: db}

	item := Item{
		Title:    "Apple",
		Link:     "https://apple.com",
		FromSite: "apple.com",
		Points:   20,
	}

	id, err := model.Create(item)
	if err != nil {
		t.Error("Error creating item ", err)
	}

	found, err := model.Find(id)
	if err != nil {
		t.Error("Could not find an item ", err)
	}

	expected := TestItem{item.Title, item.Link, item.FromSite, item.Points}
	actual := TestItem{found.Title, found.Link, found.FromSite, found.Points}

	if expected != actual {
		t.Errorf("Incorrect item returned. Want: %+v; got %+v\n", expected, actual)
	}

	nonExistingID := int64(-1)
	_, err = model.Find(nonExistingID)
	if err == nil {
		t.Errorf("Should get an error when finding a non existing item, but didn't get one")
	}
}

func TestItem_Destroy(t *testing.T) {
	db := newTestDB(t)

	model := ItemModel{DB: db}
	items, err := model.AllNew()
	if err != nil {
		t.Error(err)
	}

	model.Destroy(items[0].ID)

	items, err = model.AllNew()
	if err != nil {
		t.Error(err)
	}

	if len(items) != 0 {
		t.Errorf("Wrong items count returned. Want no items; got %d\n", len(items))
	}
}
