package adapters

import (
	"context"

	"github.com/masena-dev/bookstore-api/internal/db" // SQLC generated package

	"github.com/masena-dev/bookstore-api/internal/domain"
	"github.com/masena-dev/bookstore-api/internal/ports"
)

type SQLCAuthorRepository struct {
	Queries *db.Queries
}

func NewSQLCAuthorRepository(queries *db.Queries) ports.AuthorRepository {
	return &SQLCAuthorRepository{Queries: queries}
}

func (r *SQLCAuthorRepository) GetAuthor(ctx context.Context, id int64) (*domain.Author, error) {
	author, err := r.Queries.GetAuthor(ctx, id)
	if err != nil {
		return nil, err
	}
	return &domain.Author{
		ID: author.ID,
	}, nil
}

func (r *SQLCAuthorRepository) CreateAuthor(ctx context.Context, author *domain.Author) (*domain.Author, error) {
	dbAuthor, err := r.Queries.GetAuthor(ctx, author.ID)
	if err != nil {
		return nil, err
	}
	return &domain.Author{
		ID: dbAuthor.ID,
	}, nil
}
