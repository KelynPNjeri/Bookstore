package main

import (
	"github.com/KelynPNjeri/Bookstore/controllers"
	"github.com/KelynPNjeri/Bookstore/database"
	"github.com/KelynPNjeri/Bookstore/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	db, err := database.OpenDBConnection()

	if err != nil {
		log.Fatalln(err)
	}
	err = db.DB().Ping()
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&model.Book{})
	defer db.Close()

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Routes
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to the Library."})

	})
	r.GET("/books", controllers.GetAllBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.GetBook)
	r.PUT("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
	r.Run()
}
