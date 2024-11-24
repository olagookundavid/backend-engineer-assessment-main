package types

import (
	"time"

	"github.com/masena-dev/bookstore-api/internal/domain"
)

// Response types
type AuthorsResponse struct {
	Message string           `json:"message"`
	Authors []*domain.Author `json:"authors"`
}
type AuthorResponse struct {
	Message string         `json:"message"`
	Author  *domain.Author `json:"author"`
}
type AuthorStatsResponse struct {
	Message string              `json:"message"`
	Author  *domain.AuthorStats `json:"author_stats"`
}

type BooksResponse struct {
	Message string         `json:"message"`
	Books   []*domain.Book `json:"books"`
}

type CreateBookResponse struct {
	Message string       `json:"message"`
	Book    *domain.Book `json:"book"`
}

type GetBookResponse struct {
	Message string       `json:"message"`
	Book    *domain.Book `json:"book"`
}

type UpdateBookResponse struct {
	Message string       `json:"message"`
	Book    *domain.Book `json:"book"`
}
type MesageResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

// Request types
type CreateBookRequest struct {
	Title         string  `json:"title"`
	ISBN          string  `json:"isbn"`
	Description   string  `json:"description,omitempty"`
	Price         float64 `json:"price"`
	AuthorID      int64   `json:"author_id"`
	PublishedDate string  `json:"published_date"` // Format: YYYY-MM-DD
}

type UpdateBookRequest struct {
	Title         *string  `json:"title,omitempty"`
	Description   *string  `json:"description,omitempty"`
	Price         *float64 `json:"price,omitempty"`
	PublishedDate *string  `json:"published_date,omitempty"` // Format: YYYY-MM-DD
}

// Validation methods
func (r CreateBookRequest) Validate() error {
	if r.Title == "" || len(r.Title) > 255 {
		return ValidationError{Field: "title", Message: "must be between 1 and 255 characters"}
	}

	if !isValidISBN(r.ISBN) {
		return ValidationError{Field: "isbn", Message: "must be a valid ISBN format"}
	}

	if r.Price < 0 {
		return ValidationError{Field: "price", Message: "must be greater than or equal to 0"}
	}

	if !isValidDate(r.PublishedDate) {
		return ValidationError{Field: "published_date", Message: "must be in YYYY-MM-DD format"}
	}

	return nil
}

func (r UpdateBookRequest) Validate() error {
	if *r.Title != "" && len(*r.Title) > 255 {
		return ValidationError{Field: "title", Message: "must be between 1 and 255 characters"}
	}

	if r.Price != nil && *r.Price < 0 {
		return ValidationError{Field: "price", Message: "must be greater than or equal to 0"}
	}

	if *r.PublishedDate != "" && !isValidDate(*r.PublishedDate) {
		return ValidationError{Field: "published_date", Message: "must be in YYYY-MM-DD format"}
	}

	return nil
}

// ValidationError represents a validation error for a specific field
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return e.Field + ": " + e.Message
}

// Helper functions for validation
func isValidISBN(isbn string) bool {
	// Simple check for now - feel free to improve this
	if len(isbn) < 10 || len(isbn) > 17 {
		return false
	}
	for _, r := range isbn {
		if r != '-' && (r < '0' || r > '9') {
			return false
		}
	}
	return true
}

func isValidDate(date string) bool {
	_, err := time.Parse("2006-01-02", date)
	return err == nil
}
