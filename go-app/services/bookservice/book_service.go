package bookservice

import (
	"go-app/domain/book"
	"go-app/repositories/bookrepo"
)

// BookService interface
type BookService interface {
	CreateBook(book *book.Book) (*book.Book, error)
}

type bookService struct {
	Repo bookrepo.Repo
}

// NewBookService will instantiate User Service
func NewBookService(
	repo bookrepo.Repo,
) BookService {

	return &bookService{
		Repo: repo,
	}
}

func (bs *bookService) CreateBook(book *book.Book) (*book.Book, error) {
	return bs.Repo.CreateBook(book)
}
