package services

import "github.com/kryast/crud-15.git/repositories"

type BookService interface {
}

type bookService struct {
	repo repositories.BookRepository
}

func NewBookService(repo repositories.BookRepository) BookService {
	return &bookService{repo}
}
