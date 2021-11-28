package data

import (
	"time"
)

type Item struct {
	ID        int       `db:"id" json:"id"`                 // id biging
	Title     string    `db:"title" json:"title"`           // title text limit 1024
	Link      string    `db:"link" json:"link"`             // link text limit 2048
	FromSite  string    `db:"from_site" json:"from_site"`   // from_site text limit 128
	Points    int       `db:"points" json:"points"`         // points int default 0
	CreatedAt time.Time `db:"created_at" json:"created_at"` // created_at datetime
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"` // updated_at datetime
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
