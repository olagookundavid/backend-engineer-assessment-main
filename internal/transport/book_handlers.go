package transport

import (
	"net/http"
	"strings"

	"github.com/masena-dev/bookstore-api/internal/helpers"
	"github.com/masena-dev/bookstore-api/internal/services"
	"github.com/masena-dev/bookstore-api/internal/types"
)

type BookHandler struct {
	BookService services.IBookService
}

type IBookHandler interface {
	GetBook(w http.ResponseWriter, r *http.Request)
	GetAllBooks(w http.ResponseWriter, r *http.Request)
	CreateBook(w http.ResponseWriter, r *http.Request)
	UpdateBook(w http.ResponseWriter, r *http.Request)
	DeleteBook(w http.ResponseWriter, r *http.Request)
}

func NewBookHandler(service services.IBookService) *BookHandler {
	return &BookHandler{BookService: service}
}

func (h *BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.ReadIDParam(r)
	if err != nil {
		helpers.NotFoundResponseWithMsg(w, r, "Book not found")
		return
	}
	book, err := h.BookService.GetBook(r.Context(), id)
	if err != nil {
		if err == ErrNoBookFound {
			helpers.NotFoundResponseWithMsg(w, r, err.Error())
			return
		}
		helpers.ServerErrorResponse(w, r, err)
		return
	}

	data := types.GetBookResponse{
		Message: "Retrieved book",
		Book:    book,
	}
	err = helpers.WriteJSON(w, http.StatusOK, data, nil)
	if err != nil {
		helpers.ServerErrorResponse(w, r, err)
	}
}

func (h *BookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.BookService.GetAllBooks(r.Context())
	if err != nil {
		helpers.ServerErrorResponse(w, r, err)
		return
	}

	data := types.BooksResponse{
		Message: "Retrieved all books",
		Books:   books,
	}
	err = helpers.WriteJSON(w, http.StatusOK, data, nil)
	if err != nil {
		helpers.ServerErrorResponse(w, r, err)
	}
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var req types.CreateBookRequest
	err := helpers.ReadJSON(w, r, &req)
	if err != nil {
		helpers.BadRequestResponse(w, r, err)
		return
	}

	createBookParams, err := ToCreateBookParams(req)
	if err != nil {
		helpers.BadRequestResponse(w, r, err)
		return
	}

	book, err := h.BookService.CreateBook(r.Context(), createBookParams)
	if err != nil {
		if err == ErrBookExist {
			helpers.RecordAlreadyExistsResponse(w, r, err)
			return
		}
		helpers.ServerErrorResponse(w, r, err)
		return
	}

	data := types.CreateBookResponse{
		Message: "Created book",
		Book:    book}

	err = helpers.WriteJSON(w, http.StatusCreated, data, nil)
	if err != nil {
		helpers.ServerErrorResponse(w, r, err)
	}
}

func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {

	id, err := helpers.ReadIDParam(r)
	if err != nil {
		helpers.NotFoundResponseWithMsg(w, r, err.Error())
		return
	}

	book, err := h.BookService.GetBook(r.Context(), id)
	if err != nil {
		helpers.NotFoundResponseWithMsg(w, r, "Book not found")
		return
	}

	var req types.UpdateBookRequest
	err = helpers.ReadJSON(w, r, &req)
	if err != nil {
		helpers.BadRequestResponse(w, r, err)
		return
	}
	if req.Title == nil {
		req.Title = &book.Title
	}
	if req.Description == nil {
		req.Description = &book.Description
	}
	if req.Price == nil {
		req.Price = &book.Price
	}
	if req.PublishedDate == nil {
		req.PublishedDate = &(strings.Split(book.PublishedDate, " ")[0])
	}

	updateBookParams, err := ToUpdateBookParams(req, id)
	if err != nil {
		helpers.BadRequestResponse(w, r, err)
		return
	}
	book, err = h.BookService.UpdateBook(r.Context(), updateBookParams)
	if err != nil {
		if err == ErrNoBookFound {
			helpers.NotFoundResponseWithMsg(w, r, err.Error())
			return
		}
		helpers.ServerErrorResponse(w, r, err)
		return
	}

	data := types.UpdateBookResponse{
		Message: "Updated book",
		Book:    book}

	err = helpers.WriteJSON(w, http.StatusOK, data, nil)
	if err != nil {
		helpers.ServerErrorResponse(w, r, err)
	}
}

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.ReadIDParam(r)
	if err != nil {
		helpers.NotFoundResponseWithMsg(w, r, err.Error())
		return
	}
	err = h.BookService.DeleteBook(r.Context(), id)
	if err != nil {
		if err == ErrNoBookFound {
			helpers.NotFoundResponseWithMsg(w, r, err.Error())
			return
		}
		helpers.ServerErrorResponse(w, r, err)
		return
	}

	err = helpers.WriteJSON(w, http.StatusNoContent, nil, nil)
	if err != nil {
		helpers.ServerErrorResponse(w, r, err)
	}
}
