package adapters

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

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
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, transport.ErrNoAuthorFound
		}
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

	dbAuthorStat, err := r.Queries.GetAuthorStats(ctx, authorID)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, transport.ErrNoAuthorStatsFound
		}
		return nil, err
	}
	var bookByYear map[string]int
	if dbAuthorStat.BooksByYear != nil {
		var booksByYear map[string]int
		err := json.Unmarshal(dbAuthorStat.BooksByYear, &booksByYear)
		if err != nil {
			return nil, fmt.Errorf("failed to parse BooksByYear: %v", err)
		}
		bookByYear = booksByYear
	} else {
		bookByYear = map[string]int{}
	}

	return &domain.AuthorStats{
		ID:                  dbAuthorStat.AuthorID,
		Name:                dbAuthorStat.AuthorName,
		TotalBooks:          int(dbAuthorStat.TotalBooks),
		AverageBookPrice:    dbAuthorStat.AverageBookPrice,
		EarliestPublication: dbAuthorStat.EarliestPublication,
		LatestPublication:   dbAuthorStat.LatestPublication,
		TotalRevenue:        dbAuthorStat.TotalRevenue,
		BooksByYear:         bookByYear,
	}, nil
}
