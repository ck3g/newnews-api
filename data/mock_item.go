package data

import (
	"errors"
	"time"
)

// TODO: rename MockItemsModel to MockItemModel
type MockItemsModel struct {
}

var items []*Item
var lastID int64

func (m *MockItemsModel) AllNew() ([]*Item, error) {
	return items, nil
}

func (m *MockItemsModel) Create(item Item) (int64, error) {
	item.ID = lastID + 1
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()

	items = append(items, &item)

	lastID = item.ID

	return item.ID, nil
}

func (m *MockItemsModel) Find(id int64) (*Item, error) {
	for _, item := range items {
		if item.ID == id {
			return item, nil
		}
	}

	return nil, errors.New("not found")
}

func (m *MockItemsModel) Destroy(id int64) {
	var index int
	for i, item := range items {
		if item.ID == id {
			index = i
		}
	}

	if index != 0 {
		items = append(items[:index], items[index+1:]...)
	}
}

func (m *MockItemsModel) Truncate() {
	items = []*Item{}
	lastID = 0
}
