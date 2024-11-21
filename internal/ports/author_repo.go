package ports

import (
	"context"

	"github.com/masena-dev/bookstore-api/internal/domain"
)

type AuthorRepository interface {
	GetAuthor(ctx context.Context, id int64) (*domain.Author, error)
	CreateAuthor(ctx context.Context, user *domain.Author) (*domain.Author, error)
}
