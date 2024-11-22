package domain

import "time"

type Author struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name,omitempty"`
	Bio       string    `json:"bio,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type AuthorStats struct {
	ID                  int64          `json:"id"`
	Name                string         `json:"name"`
	TotalBooks          int            `json:"total_books"`
	AverageBookPrice    float64        `json:"average_book_price"`
	EarliestPublication string         `json:"earliest_publication"` // YYYY-MM-DD
	LatestPublication   string         `json:"latest_publication"`   // YYYY-MM-DD
	TotalRevenue        float64        `json:"total_revenue"`
	BooksByYear         map[string]int `json:"books_by_year"` // e.g. {"2023": 3, "2024": 1}
}
