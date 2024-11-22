package transport

import (
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/masena-dev/bookstore-api/internal/apis"
	"github.com/masena-dev/bookstore-api/internal/db"
	"github.com/masena-dev/bookstore-api/internal/domain"
)

func ToCreateBookParams(req apis.CreateBookRequest) (db.CreateBookParams, error) {

	err := req.Validate()
	if err != nil {
		return db.CreateBookParams{}, err
	}
	publishedDate, err := time.Parse("2006-01-02", req.PublishedDate)
	if err != nil {
		return db.CreateBookParams{}, fmt.Errorf("invalid published date: %w", err)
	}
	return db.CreateBookParams{
		Title:         req.Title,
		Isbn:          req.ISBN,
		Description:   toPgtext(req.Description),
		Price:         toPgNumeric(req.Price),
		AuthorID:      req.AuthorID,
		PublishedDate: toPgDate(publishedDate),
	}, nil
}

func ToUpdateBookParams(req apis.UpdateBookRequest, Id int64) (db.UpdateBookParams, error) {

	err := req.Validate()
	if err != nil {
		return db.UpdateBookParams{}, err
	}
	publishedDate, err := time.Parse("2006-01-02", *req.PublishedDate)
	if err != nil {
		return db.UpdateBookParams{}, fmt.Errorf("invalid published date: %w", err)
	}
	return db.UpdateBookParams{
		Title:         *req.Title,
		Description:   toPgtext(*req.Description),
		Price:         toPgNumeric(*req.Price),
		PublishedDate: toPgDate(publishedDate),
		ID:            Id,
	}, nil
}

func toPgtext(str string) pgtype.Text {
	return pgtype.Text{
		String: str, Valid: true,
	}
}

func toPgNumeric(value float64) pgtype.Numeric {
	exp := 0

	for value != math.Trunc(value) {
		value *= 10
		exp--
	}

	intPart := big.NewInt(int64(value))

	return pgtype.Numeric{
		Int:   intPart,
		Exp:   int32(exp),
		Valid: true,
	}
}

func toPgDate(time time.Time) pgtype.Date {
	return pgtype.Date{
		Valid: true,
		Time:  time,
	}
}

func ConvertToDomainBooks(rows []db.ListBooksRow) []*domain.Book {
	books := make([]*domain.Book, 0, len(rows))

	for _, row := range rows {
		value, _ := row.Price.Float64Value()
		book := &domain.Book{
			ID:            row.ID,
			CreatedAt:     row.CreatedAt.Time,
			UpdatedAt:     row.UpdatedAt.Time,
			Title:         row.Title,
			ISBN:          row.Isbn,
			Description:   row.Description.String,
			PublishedDate: row.PublishedDate.Time.String(),
			Price:         value.Float64,
			Author: domain.Author{
				ID:        row.AuthorID,
				Name:      row.AuthorName,
				Bio:       row.AuthorBio.String,
				CreatedAt: row.AuthorCreatedAt.Time,
				UpdatedAt: row.AuthorUpdatedAt.Time,
			},
		}
		books = append(books, book)
	}
	return books
}

func ConvertToDomainAuthors(rows []db.Author) []*domain.Author {
	authors := make([]*domain.Author, 0, len(rows))

	for _, row := range rows {
		book := &domain.Author{
			ID:        row.ID,
			Name:      row.Name,
			Bio:       row.Bio.String,
			CreatedAt: row.CreatedAt.Time,
			UpdatedAt: row.UpdatedAt.Time,
		}
		authors = append(authors, book)
	}
	return authors
}

// ID:        author.ID,
// 		Name:      author.Name,
// 		Bio:       author.Bio.String,
// 		CreatedAt: author.CreatedAt.Time,
// 		UpdatedAt: author.UpdatedAt.Time,
