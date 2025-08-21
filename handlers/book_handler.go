package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kryast/crud-15.git/models"
	"github.com/kryast/crud-15.git/services"
)

type BookHandler struct {
	service services.BookService
}

func NewBookHandler(service services.BookService) *BookHandler {
	return &BookHandler{service}
}

func (bh *BookHandler) Create(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}

	if err := bh.service.Create(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (bh *BookHandler) GetAll(c *gin.Context) {
	books, err := bh.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err})
		return
	}

	c.JSON(http.StatusOK, books)
}

func (bh *BookHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	book, err := bh.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (bh *BookHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}

	book.ID = uint(id)
	if err := bh.service.Update(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err})
		return
	}

	c.JSON(http.StatusOK, book)
}
