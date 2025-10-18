package handlers

import (
	"BlogAPI/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateAuthor(c *gin.Context) {
	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Save author to database
	c.JSON(http.StatusCreated, author)
}

func GetAuthors(c *gin.Context) {
	// Fetch authors from database
	c.JSON(http.StatusOK, []models.Author{})
}

func GetAuthor(c *gin.Context) {
	// Fetch single author by ID
	c.JSON(http.StatusOK, models.Author{})
}

func UpdateAuthor(c *gin.Context) {
	// Update author logic
	c.JSON(http.StatusOK, models.Author{})
}

func DeleteAuthor(c *gin.Context) {
	// Delete author logic
	c.Status(http.StatusNoContent)
}
