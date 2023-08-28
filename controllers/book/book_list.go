package book

import (
	"net/http"

	"github.com/beinganukul/rest_api_golang/initializers"
	"github.com/beinganukul/rest_api_golang/models"
	"github.com/gin-gonic/gin"
)

func GetAllBooks(c *gin.Context) {
	var books []models.Book
	result := initializers.DB.Find(&books)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "couldn't load books from database!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"books": books,
	})
}
