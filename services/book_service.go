package services

import (
	"github.com/kryast/crud-15.git/models"
	"github.com/kryast/crud-15.git/repositories"
)

type BookService interface {
	Create(book *models.Book) error
	GetAll() ([]models.Book, error)
	GetByID(id uint) (*models.Book, error)
	Update(book *models.Book) error
	Delete(id uint) error
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

func (bs *bookService) GetAll() ([]models.Book, error) {
	return bs.repo.GetAll()
}

func (bs *bookService) GetByID(id uint) (*models.Book, error) {
	return bs.repo.GetByID(id)
}

func (bs *bookService) Update(book *models.Book) error {
	return bs.repo.Update(book)
}

func (bs *bookService) Delete(id uint) error {
	return bs.repo.Delete(id)
}
