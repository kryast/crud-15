package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kryast/crud-15.git/handlers"
	"github.com/kryast/crud-15.git/repositories"
	"github.com/kryast/crud-15.git/services"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	repo := repositories.NewBookRepository(db)
	svc := services.NewBookService(repo)
	h := handlers.NewBookHandler(svc)

	r.POST("/book", h.Create)
	r.GET("/book", h.GetAll)

	return r
}
