package mockdb

import (
	"errors"
	"time"

	"github.com/ck3g/newnews-api/data"
)

type MockItemModel struct {
}

var items []*data.Item
var lastID int64

func (m *MockItemModel) AllNew() ([]*data.Item, error) {
	return items, nil
}

func (m *MockItemModel) Create(item data.Item) (int64, error) {
	item.ID = lastID + 1
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()

	items = append(items, &item)

	lastID = item.ID

	return item.ID, nil
}

func (m *MockItemModel) Find(id int64) (*data.Item, error) {
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
	items = []*data.Item{}
	lastID = 0
}
