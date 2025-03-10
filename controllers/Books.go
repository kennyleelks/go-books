package controllers

import (
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"github.com/guregu/null"
	"net/http"
)

type AddBookInput struct {
	Title  string      `json:"title" binding:"required"`
	Author null.String `json:"author"`
	ISBN   null.String `json:"isbn"`
}

func GetBooks(c *gin.Context) {
	var books []models.Book

	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func CreateBook(c *gin.Context) {
	var input AddBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{
		Title:  input.Title,
		Author: input.Author,
		ISBN:   input.ISBN,
	}

	models.DB.Create(&book)

	c.JSON(http.StatusCreated, gin.H{"data": book})
}
