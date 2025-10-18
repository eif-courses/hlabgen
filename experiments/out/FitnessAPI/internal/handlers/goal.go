package handlers

import (
	"FitnessAPI/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateGoal(c *gin.Context) {
	var goal models.Goal
	if err := c.ShouldBindJSON(&goal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Logic to save goal to database
	c.JSON(http.StatusCreated, goal)
}

func GetGoal(c *gin.Context) {
	// Logic to get goal from database
}

func UpdateGoal(c *gin.Context) {
	// Logic to update goal in database
}

func DeleteGoal(c *gin.Context) {
	// Logic to delete goal from database
}
