package handlers

import (
	"github.com/masena-dev/bookstore-api/internal/db"
	"github.com/masena-dev/bookstore-api/internal/transport"
)

type Handlers struct {
	Authors transport.AuthorHandler
}

func NewHandlers(db *db.Queries) Handlers {
	return Handlers{
		//TODO
		Authors: transport.AuthorHandler{},
	}
}
