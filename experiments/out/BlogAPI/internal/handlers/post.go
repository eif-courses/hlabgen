package handlers

import (
	"BlogAPI/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Save post to database
	c.JSON(http.StatusCreated, post)
}

func GetPosts(c *gin.Context) {
	// Fetch posts from database
	c.JSON(http.StatusOK, []models.Post{})
}

func GetPost(c *gin.Context) {
	// Fetch single post by ID
	c.JSON(http.StatusOK, models.Post{})
}

func UpdatePost(c *gin.Context) {
	// Update post logic
	c.JSON(http.StatusOK, models.Post{})
}

func DeletePost(c *gin.Context) {
	// Delete post logic
	c.Status(http.StatusNoContent)
}
