package controllers

import (
	"github.com/KelynPNjeri/Bookstore/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func GetAllBooks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var books []model.Book
	db.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": &books})
}

func CreateBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book = model.Book{BookName: book.BookName, BookDesc: book.BookDesc, Category: book.Category}
	db.Create(&book)
	c.JSON(http.StatusCreated, gin.H{"data": book})
}
func GetBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var book model.Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})

}

func UpdateBook(c *gin.Context) {

}

func DeleteBook(c *gin.Context)  {

}