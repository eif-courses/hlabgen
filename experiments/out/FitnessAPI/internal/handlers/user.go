package handlers

import (
	"FitnessAPI/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Logic to save user to database
	c.JSON(http.StatusCreated, user)
}

func GetUser(c *gin.Context) {
	// Logic to get user from database
}

func UpdateUser(c *gin.Context) {
	// Logic to update user in database
}

func DeleteUser(c *gin.Context) {
	// Logic to delete user from database
}
