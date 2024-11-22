package domain

import "time"

type Book struct {
	ID            int64     `json:"id"`
	Title         string    `json:"title"`
	ISBN          string    `json:"isbn"`
	Description   string    `json:"description"`
	Price         float64   `json:"price"`
	Author        Author    `json:"author"`
	PublishedDate string    `json:"published_date"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
