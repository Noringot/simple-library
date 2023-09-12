package tests

import (
	"context"
	"github.com/Noringotq/simple-grpc/internal/adapters/repository"
	"github.com/Noringotq/simple-grpc/internal/usecase"
	"github.com/Noringotq/simple-grpc/pkg/lib"
	"testing"
)

func TestGetAuthors(t *testing.T) {
	ctx := context.Background()
	book := lib.BookRequest{Title: "A"}
	expected := []*lib.Author{{Name: "John"}, {Name: "William"}}

	rep := repository.NewLibRepository("root", "root", "127.0.0.1:5432", "lib")
	usc := usecase.NewLibStorage(rep)
	res, err := usc.GetAuthors(ctx, &book)

	if err != nil {
		t.Errorf("Incorrect data: %v", err)
	}

	if len(expected) != len(res.GetAuthors()) {
		t.Errorf("Incorrect result. Expected %v, got %v", expected, res.Authors)
	}

	for k, v := range res.GetAuthors() {
		if v.Name != expected[k].GetName() {
			t.Errorf("Incorrect result. Expected %v, got %v", expected, res.Authors)
		}
	}
}

func TestGetBooks(t *testing.T) {
	ctx := context.Background()
	author := lib.AuthorRequest{Name: "John"}
	expected := []*lib.Book{{Title: "A"}, {Title: "B"}}

	rep := repository.NewLibRepository("root", "root", "127.0.0.1:5432", "lib")
	usc := usecase.NewLibStorage(rep)
	res, err := usc.GetBooks(ctx, &author)

	if err != nil {
		t.Errorf("Incorrect data: %v", err)
	}

	if len(expected) != len(res.GetBooks()) {
		t.Errorf("Incorrect result. Expected %v, got %v", expected, res.GetBooks())
	}

	for k, v := range res.GetBooks() {
		if v.Title != expected[k].GetTitle() {
			t.Errorf("Incorrect result. Expected %v, got %v", expected, res.GetBooks())
		}
	}
}
