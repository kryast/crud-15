package handlers

import "github.com/kryast/crud-15.git/services"

type BookHandler struct {
	service services.BookService
}

func NewBookHandler(service services.BookService) *BookHandler {
	return &BookHandler{service}
}
