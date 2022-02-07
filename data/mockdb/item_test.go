package mockdb

import (
	"testing"

	"github.com/ck3g/newnews-api/data"
)

func TestMockItem_AllNew(t *testing.T) {
	model := ItemModel{}
	model.Truncate()

	items, _ := model.AllNew()
	if len(items) != 0 {
		t.Error("Items returned when not expected")
	}

	id, _ := model.Create(data.Item{
		Title:    "Google",
		Link:     "https://google.com",
		Points:   10,
		FromSite: "google.com",
	})

	items, _ = model.AllNew()
	if len(items) != 1 {
		t.Errorf("Wrong number of items returned. Want 1; got %d", len(items))
	}

	if items[0].ID != id {
		t.Errorf("Invalid item ID; Want %d; got %d;", id, items[0].ID)
	}
}

func TestMockItem_Create(t *testing.T) {
	model := ItemModel{}
	model.Truncate()

	item := data.Item{
		Title:    "Google",
		Link:     "https://google.com",
		Points:   10,
		FromSite: "google.com",
	}

	id, err := model.Create(item)
	if err != nil {
		t.Error(err)
	}

	if id != 1 {
		t.Error("Wrong item ID returned")
	}

	items, _ := model.AllNew()

	if len(items) != 1 {
		t.Error("Wrong items count returned")
	}
}

func TestMockItem_Find(t *testing.T) {
	model := ItemModel{}
	model.Truncate()

	item, err := model.Find(1)
	if err == nil {
		t.Error("Error not returned but it should")
	}
	if item != nil {
		t.Error("Item returned but it should not")
	}

	id, _ := model.Create(data.Item{
		Title:    "Google",
		Link:     "https://google.com",
		Points:   10,
		FromSite: "google.com",
	})

	item, err = model.Find(id)
	if err != nil {
		t.Error("Error returned when it's not expected")
	}
	if item.ID != id {
		t.Errorf("Invalid item found; Want ID %d; got %d", id, item.ID)
	}
}

func TestMockItem_Destroy(t *testing.T) {
	model := ItemModel{}
	model.Truncate()

	id1, _ := model.Create(data.Item{
		Title: "Google",
	})
	id2, _ := model.Create(data.Item{
		Title: "Microsoft",
	})
	id3, _ := model.Create(data.Item{
		Title: "Apple",
	})

	_, err := model.Find(id1)
	if err != nil {
		t.Error("First item not found, but it should")
	}

	_, err = model.Find(id2)
	if err != nil {
		t.Error("Second item not found, but it should")
	}

	_, err = model.Find(id3)
	if err != nil {
		t.Error("Third item not found, but it should")
	}

	model.Destroy(id2)

	_, err = model.Find(id1)
	if err != nil {
		t.Error("First item not found, but it should")
	}

	_, err = model.Find(id2)
	if err == nil {
		t.Error("Second item found, but it should not")
	}

	_, err = model.Find(id3)
	if err != nil {
		t.Error("Third item not found, but it should")
	}
}
