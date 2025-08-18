package services

import (
	"github.com/kryast/crud-15.git/models"
	"github.com/kryast/crud-15.git/repositories"
)

type BookService interface {
	Create(book *models.Book) error
}

type bookService struct {
	repo repositories.BookRepository
}

func NewBookService(repo repositories.BookRepository) BookService {
	return &bookService{repo}
}

func (bs *bookService) Create(book *models.Book) error {
	return bs.repo.Create(book)
}
