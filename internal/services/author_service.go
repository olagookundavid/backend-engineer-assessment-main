package services

import (
	"context"

	"github.com/masena-dev/bookstore-api/internal/domain"
	"github.com/masena-dev/bookstore-api/internal/ports"
)

type IAuthorService interface {
	GetAuthor(ctx context.Context, id int64) (*domain.Author, error)
	GetAllAuthors(ctx context.Context) ([]*domain.Author, error)
	GetAuthorStats(ctx context.Context, authorID int64) (*domain.AuthorStats, error)
}

type AuthorService struct {
	AuthorRepository ports.AuthorRepository
}

func NewAuthorService(repo ports.AuthorRepository) *AuthorService {
	return &AuthorService{AuthorRepository: repo}
}

func (s *AuthorService) GetAuthor(ctx context.Context, id int64) (*domain.Author, error) {
	return s.AuthorRepository.GetAuthor(ctx, id)
}

func (s *AuthorService) GetAllAuthors(ctx context.Context) ([]*domain.Author, error) {
	return s.AuthorRepository.GetAllAuthors(ctx)
}

func (s *AuthorService) GetAuthorStats(ctx context.Context, authorID int64) (*domain.AuthorStats, error) {
	return s.AuthorRepository.GetAuthorStats(ctx, authorID)
}
