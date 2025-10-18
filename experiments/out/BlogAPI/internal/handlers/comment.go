package handlers

import (
	"BlogAPI/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Save comment to database
	c.JSON(http.StatusCreated, comment)
}

func GetComments(c *gin.Context) {
	// Fetch comments from database
	c.JSON(http.StatusOK, []models.Comment{})
}

func GetComment(c *gin.Context) {
	// Fetch single comment by ID
	c.JSON(http.StatusOK, models.Comment{})
}

func UpdateComment(c *gin.Context) {
	// Update comment logic
	c.JSON(http.StatusOK, models.Comment{})
}

func DeleteComment(c *gin.Context) {
	// Delete comment logic
	c.Status(http.StatusNoContent)
}
