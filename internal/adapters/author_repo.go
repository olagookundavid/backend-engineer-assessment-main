package adapters

import (
	"context"

	"github.com/masena-dev/bookstore-api/internal/db"

	"github.com/masena-dev/bookstore-api/internal/domain"

	"github.com/masena-dev/bookstore-api/internal/transport"
)

type SQLCAuthorRepository struct {
	Queries *db.Queries
}

func NewSQLCAuthorRepository(queries *db.Queries) *SQLCAuthorRepository {
	return &SQLCAuthorRepository{Queries: queries}
}

func (r *SQLCAuthorRepository) GetAuthor(ctx context.Context, id int64) (*domain.Author, error) {
	author, err := r.Queries.GetAuthor(ctx, id)
	if err != nil {
		return nil, err
	}
	return &domain.Author{
		ID:        author.ID,
		Name:      author.Name,
		Bio:       author.Bio.String,
		CreatedAt: author.CreatedAt.Time,
		UpdatedAt: author.UpdatedAt.Time,
	}, nil
}

func (r *SQLCAuthorRepository) GetAllAuthors(ctx context.Context) ([]*domain.Author, error) {
	dbAuthors, err := r.Queries.ListAuthors(ctx)
	if err != nil {
		return nil, err
	}
	authors := transport.ConvertToDomainAuthors(dbAuthors)
	return authors, nil
}

func (r *SQLCAuthorRepository) GetAuthorStats(ctx context.Context, authorID int64) (*domain.AuthorStats, error) {

	dbAuthors, err := r.Queries.ListAuthors(ctx)
	if err != nil {
		return nil, err
	}
	print(dbAuthors)

	return &domain.AuthorStats{}, nil
}
