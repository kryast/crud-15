package main

import (
	"log"

	"github.com/kryast/crud-15.git/database"
	"github.com/kryast/crud-15.git/models"
	"github.com/kryast/crud-15.git/router"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Failed to Connect Database", err)
	}

	db.AutoMigrate(&models.Book{})

	r := router.SetupRouter(db)
	r.Run(":8080")
}
