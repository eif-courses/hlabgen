package handlers

import (
	"TaskAPI/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Save user to database
	c.JSON(http.StatusCreated, user)
}

func GetUsers(c *gin.Context) {
	// Retrieve users from database
	c.JSON(http.StatusOK, []models.User{})
}

func UpdateUser(c *gin.Context) {
	// Update user in database
	c.JSON(http.StatusOK, gin.H{"message": "User updated"})
}

func DeleteUser(c *gin.Context) {
	// Delete user from database
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
