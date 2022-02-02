package data

import (
	"errors"
	"time"
)

type MockItemModel struct {
}

var items []*Item
var lastID int64

func (m *MockItemModel) AllNew() ([]*Item, error) {
	return items, nil
}

func (m *MockItemModel) Create(item Item) (int64, error) {
	item.ID = lastID + 1
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()

	items = append(items, &item)

	lastID = item.ID

	return item.ID, nil
}

func (m *MockItemModel) Find(id int64) (*Item, error) {
	for _, item := range items {
		if item.ID == id {
			return item, nil
		}
	}

	return nil, errors.New("not found")
}

func (m *MockItemModel) Destroy(id int64) {
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

func (m *MockItemModel) Truncate() {
	items = []*Item{}
	lastID = 0
}
