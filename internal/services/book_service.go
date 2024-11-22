package services

import (
	"context"

	"github.com/masena-dev/bookstore-api/internal/db"
	"github.com/masena-dev/bookstore-api/internal/domain"
	"github.com/masena-dev/bookstore-api/internal/ports"
)

type IBookService interface {
	GetBook(ctx context.Context, id int64) (*domain.Book, error)
	GetAllBooks(ctx context.Context) ([]*domain.Book, error)
	CreateBook(ctx context.Context, arg db.CreateBookParams) (*domain.Book, error)
	UpdateBook(ctx context.Context, arg db.UpdateBookParams) (*domain.Book, error)
	DeleteBook(ctx context.Context, id int64) error
}

type BookService struct {
	BookRepository ports.BookRepository
}

func NewBookService(repo ports.BookRepository) *BookService {
	return &BookService{BookRepository: repo}
}

func (s *BookService) GetBook(ctx context.Context, id int64) (*domain.Book, error) {
	return s.BookRepository.GetBook(ctx, id)
}

func (s *BookService) GetAllBooks(ctx context.Context) ([]*domain.Book, error) {
	return s.BookRepository.GetAllBooks(ctx)
}
func (s *BookService) CreateBook(ctx context.Context, arg db.CreateBookParams) (*domain.Book, error) {
	return s.BookRepository.CreateBook(ctx, arg)
}

func (s *BookService) UpdateBook(ctx context.Context, arg db.UpdateBookParams) (*domain.Book, error) {
	return s.BookRepository.UpdateBook(ctx, arg)
}

func (s *BookService) DeleteBook(ctx context.Context, id int64) error {
	return s.BookRepository.DeleteBook(ctx, id)
}
