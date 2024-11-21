package services

import (
	"context"

	"github.com/masena-dev/bookstore-api/internal/domain"
	"github.com/masena-dev/bookstore-api/internal/ports"
)

type AuthorService struct {
	AuthorRepository ports.AuthorRepository
}

func NewAuthorService(repo ports.AuthorRepository) *AuthorService {
	return &AuthorService{AuthorRepository: repo}
}

func (s *AuthorService) GetAuthor(ctx context.Context, id int64) (*domain.Author, error) {
	return s.AuthorRepository.GetAuthor(ctx, id)
}

func (s *AuthorService) CreateAuthor(ctx context.Context, author *domain.Author) (*domain.Author, error) {
	return s.AuthorRepository.CreateAuthor(ctx, author)
}
