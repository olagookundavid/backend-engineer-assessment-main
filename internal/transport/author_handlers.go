package transport

import (
	"net/http"

	"github.com/masena-dev/bookstore-api/internal/helpers"
	"github.com/masena-dev/bookstore-api/internal/services"
)

type AuthorHandler struct {
	AuthorService services.IAuthorService
}

type IAuthorHandler interface {
	GetAuthor(w http.ResponseWriter, r *http.Request)
	GetAllAuthors(w http.ResponseWriter, r *http.Request)
	GetAuthorStats(w http.ResponseWriter, r *http.Request)
}

func NewAuthorHandler(service services.IAuthorService) *AuthorHandler {
	return &AuthorHandler{AuthorService: service}
}

func (h *AuthorHandler) GetAuthor(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.ReadIDParam(r)
	if err != nil {
		helpers.NotFoundResponseWithMsg(w, r, err.Error())
		return
	}
	author, err := h.AuthorService.GetAuthor(r.Context(), id)
	if err != nil {
		helpers.ServerErrorResponse(w, r, err)
		return
	}

	env := envelope{
		"message": "Retrieved author details",
		"data":    author}

	err = helpers.WriteJSON(w, http.StatusOK, env, nil)
	if err != nil {
		helpers.ServerErrorResponse(w, r, err)
	}
}

func (h *AuthorHandler) GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	authors, err := h.AuthorService.GetAllAuthors(r.Context())
	if err != nil {
		helpers.ServerErrorResponse(w, r, err)
		return
	}

	env := envelope{
		"message": "Retrieved all authors",
		"data":    authors}

	err = helpers.WriteJSON(w, http.StatusOK, env, nil)
	if err != nil {
		helpers.ServerErrorResponse(w, r, err)
	}
}

func (h *AuthorHandler) GetAuthorStats(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.ReadIDParam(r)
	if err != nil {
		helpers.NotFoundResponseWithMsg(w, r, err.Error())
		return
	}
	author, err := h.AuthorService.GetAuthor(r.Context(), id)
	if err != nil {
		helpers.ServerErrorResponse(w, r, err)
		return
	}

	env := envelope{
		"message": "Retrieved author details",
		"data":    author}

	err = helpers.WriteJSON(w, http.StatusOK, env, nil)
	if err != nil {
		helpers.ServerErrorResponse(w, r, err)
	}
}
