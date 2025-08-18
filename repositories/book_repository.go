package repositories

import (
	"github.com/kryast/crud-15.git/models"
	"gorm.io/gorm"
)

type BookRepository interface {
	Create(book *models.Book) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db}
}

func (br *bookRepository) Create(book *models.Book) error {
	return br.db.Create(book).Error
}
