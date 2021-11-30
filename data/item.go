package data

import (
	"time"
)

type Item struct {
	ID        int       `db:"id" json:"id"`
	Title     string    `db:"title" json:"title"`
	Link      string    `db:"link" json:"link"`
	FromSite  string    `db:"from_site" json:"from_site"`
	Points    int       `db:"points" json:"points"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func (i *Item) GetAllNew() ([]*Item, error) {
	all := []*Item{
		{
			ID:        1,
			Title:     "Google",
			Link:      "https://google.com",
			FromSite:  "google.com",
			Points:    5,
			CreatedAt: time.Now().Add(time.Duration(-1) * time.Hour),
			UpdatedAt: time.Now().Add(time.Duration(-1) * time.Hour),
		},
		{
			ID:        2,
			Title:     "Apple",
			Link:      "https://apple.com",
			FromSite:  "apple.com",
			Points:    10,
			CreatedAt: time.Now().Add(time.Duration(-30) * time.Minute),
			UpdatedAt: time.Now().Add(time.Duration(-30) * time.Minute),
		},
	}

	return all, nil
}
