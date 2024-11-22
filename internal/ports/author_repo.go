package ports

import (
	"context"

	"github.com/masena-dev/bookstore-api/internal/domain"
)

type AuthorRepository interface {
	GetAuthor(ctx context.Context, id int64) (*domain.Author, error)
	GetAllAuthors(ctx context.Context) ([]*domain.Author, error)
	GetAuthorStats(ctx context.Context, authorID int64) (*domain.AuthorStats, error)
}
