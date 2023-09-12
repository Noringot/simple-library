package repository

import (
	"context"
	"fmt"
	"github.com/Noringotq/simple-grpc/pkg/driver/mysql"
	"github.com/Noringotq/simple-grpc/pkg/lib"
)

type LibRepository struct {
	*mysql.Driver
}

func NewLibRepository(user, pass, host, db string) *LibRepository {
	driver := mysql.New(user, pass, host, db)
	return &LibRepository{driver}
}

func (r *LibRepository) GetBooks(ctx context.Context, author string) ([]*lib.Book, error) {
	var books []*lib.Book
	query := fmt.Sprintf("SELECT `books`.title FROM books JOIN authors ON `authors`.id=`books`.author_id WHERE name='%s'", author)

	rows, err := r.Db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var book lib.Book
		if err := rows.Scan(&book.Title); err != nil {
			return nil, err
		}

		books = append(books, &book)
	}

	return books, nil
}

func (r *LibRepository) GetAuthors(ctx context.Context, book string) ([]*lib.Author, error) {
	var authors []*lib.Author
	query := fmt.Sprintf("SELECT authors.name FROM books JOIN authors ON authors.id=books.author_id WHERE books.title='%s'", book)

	rows, err := r.Db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var author lib.Author
		if err := rows.Scan(&author.Name); err != nil {
			return nil, err
		}

		authors = append(authors, &author)
	}

	return authors, nil
}
