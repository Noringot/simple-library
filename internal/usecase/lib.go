package usecase

import (
	"context"
	"github.com/Noringotq/simple-grpc/internal/adapters/repository"
	"github.com/Noringotq/simple-grpc/pkg/lib"
)

type LibStorage struct {
	Repo *repository.LibRepository
	lib.UnimplementedLibServer
}

func NewLibStorage(repo *repository.LibRepository) *LibStorage {
	return &LibStorage{
		Repo: repo,
	}
}

func (s *LibStorage) GetAuthors(ctx context.Context, book *lib.BookRequest) (*lib.AuthorsResponse, error) {
	authors, err := s.Repo.GetAuthors(ctx, book.GetTitle())
	if err != nil {
		return nil, err
	}

	return &lib.AuthorsResponse{
		Authors: authors,
	}, nil
}

func (s *LibStorage) GetBooks(ctx context.Context, author *lib.AuthorRequest) (*lib.BooksResponse, error) {
	books, err := s.Repo.GetBooks(ctx, author.GetName())

	if err != nil {
		return nil, err
	}

	return &lib.BooksResponse{Books: books}, nil
}
