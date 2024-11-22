package ports

import (
	"context"

	"github.com/masena-dev/bookstore-api/internal/db"
	"github.com/masena-dev/bookstore-api/internal/domain"
)

type BookRepository interface {
	GetBook(ctx context.Context, id int64) (*domain.Book, error)
	GetAllBooks(ctx context.Context) ([]*domain.Book, error)
	CreateBook(ctx context.Context, arg db.CreateBookParams) (*domain.Book, error)
	UpdateBook(ctx context.Context, arg db.UpdateBookParams) (*domain.Book, error)
	DeleteBook(ctx context.Context, id int64) error
}
