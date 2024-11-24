package adapters

import (
	"context"
	"strings"

	"github.com/masena-dev/bookstore-api/internal/db"
	"github.com/masena-dev/bookstore-api/internal/domain"
	"github.com/masena-dev/bookstore-api/internal/transport"
)

type SQLCBookRepository struct {
	Queries *db.Queries
}

func NewSQLCBookRepository(queries *db.Queries) *SQLCBookRepository {
	return &SQLCBookRepository{Queries: queries}
}

func (r *SQLCBookRepository) CreateBook(ctx context.Context, arg db.CreateBookParams) (*domain.Book, error) {
	book, err := r.Queries.CreateBook(ctx, arg)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil, transport.ErrBookExist
		}
		return nil, err
	}
	value, _ := book.Price.Float64Value()
	return &domain.Book{
		ID:            book.ID,
		CreatedAt:     book.CreatedAt.Time,
		UpdatedAt:     book.UpdatedAt.Time,
		Title:         book.Title,
		ISBN:          book.Isbn,
		Description:   book.Description.String,
		PublishedDate: book.PublishedDate.Time.String(),
		Price:         value.Float64,
		Author: domain.Author{
			ID:        book.AuthorID,
			Name:      book.AuthorName,
			Bio:       book.AuthorBio.String,
			CreatedAt: book.AuthorCreatedAt.Time,
			UpdatedAt: book.AuthorUpdatedAt.Time,
		},
	}, nil
}

func (r *SQLCBookRepository) GetBook(ctx context.Context, id int64) (*domain.Book, error) {
	book, err := r.Queries.GetBook(ctx, id)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, transport.ErrNoBookFound
		}
		return nil, err
	}
	value, _ := book.Price.Float64Value()
	return &domain.Book{
		ID:            book.ID,
		CreatedAt:     book.CreatedAt.Time,
		UpdatedAt:     book.UpdatedAt.Time,
		Title:         book.Title,
		ISBN:          book.Isbn,
		Description:   book.Description.String,
		PublishedDate: book.PublishedDate.Time.String(),
		Price:         value.Float64,
		Author: domain.Author{
			ID:        book.AuthorID,
			Name:      book.AuthorName,
			Bio:       book.AuthorBio.String,
			CreatedAt: book.AuthorCreatedAt.Time,
			UpdatedAt: book.AuthorUpdatedAt.Time,
		},
	}, nil
}
func (r *SQLCBookRepository) UpdateBook(ctx context.Context, arg db.UpdateBookParams) (*domain.Book, error) {

	book, err := r.Queries.UpdateBook(ctx, arg)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, transport.ErrNoBookFound
		}
		return nil, err
	}

	value, _ := book.Price.Float64Value()
	return &domain.Book{
		ID:            book.ID,
		CreatedAt:     book.CreatedAt.Time,
		UpdatedAt:     book.UpdatedAt.Time,
		Title:         book.Title,
		ISBN:          book.Isbn,
		Description:   book.Description.String,
		PublishedDate: book.PublishedDate.Time.String(),
		Price:         value.Float64,
		Author: domain.Author{
			ID: book.AuthorID,
			// Name:      book.AuthorName,
			// Bio:       book.AuthorBio.String,
			// CreatedAt: book.AuthorCreatedAt.Time,
			// UpdatedAt: book.AuthorUpdatedAt.Time,
		},
	}, nil
}
func (r *SQLCBookRepository) DeleteBook(ctx context.Context, id int64) error {
	err := r.Queries.DeleteBook(ctx, id)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return transport.ErrNoBookFound
		}
		return err
	}
	return nil
}

func (r *SQLCBookRepository) GetAllBooks(ctx context.Context) ([]*domain.Book, error) {
	dbBooks, err := r.Queries.ListBooks(ctx)
	if err != nil {
		return nil, err
	}
	books := transport.ConvertToDomainBooks(dbBooks)
	return books, nil
}
