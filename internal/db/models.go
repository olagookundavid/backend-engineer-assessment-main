// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Author struct {
	ID        int64
	Name      string
	Bio       pgtype.Text
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}

type Book struct {
	ID            int64
	Title         string
	Isbn          string
	Description   pgtype.Text
	Price         pgtype.Numeric
	AuthorID      int64
	PublishedDate pgtype.Date
	CreatedAt     pgtype.Timestamptz
	UpdatedAt     pgtype.Timestamptz
}

type BookSale struct {
	ID          int64
	BookID      int64
	PurchasedAt pgtype.Timestamptz
}