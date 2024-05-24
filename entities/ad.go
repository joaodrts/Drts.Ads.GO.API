package entities

import "time"

type Ad struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DateStart   time.Time `json:"date_start"`
	DateEnd     time.Time `json:"date_end"`
	Status      bool      `json:"status"`
	Link        string    `json:"link"`
}
