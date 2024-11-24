package handlers

import (
	"github.com/masena-dev/bookstore-api/internal/adapters"
	"github.com/masena-dev/bookstore-api/internal/db"
	"github.com/masena-dev/bookstore-api/internal/services"
	"github.com/masena-dev/bookstore-api/internal/transport"
)

type Handlers struct {
	AuthorsHandler transport.IAuthorHandler
	BooksHandler   transport.IBookHandler
}

func NewHandlers(db *db.Queries) Handlers {
	//Authors
	authorRepo := adapters.NewSQLCAuthorRepository(db)
	authorService := services.NewAuthorService(authorRepo)
	authorHandler := transport.NewAuthorHandler(authorService)

	//Books
	bookRepo := adapters.NewSQLCBookRepository(db)
	bookService := services.NewBookService(bookRepo)
	bookHandler := transport.NewBookHandler(bookService)

	return Handlers{
		AuthorsHandler: authorHandler,
		BooksHandler:   bookHandler,
	}
}
