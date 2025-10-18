package handlers

import (
	"FitnessAPI/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateWorkout(c *gin.Context) {
	var workout models.Workout
	if err := c.ShouldBindJSON(&workout); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Logic to save workout to database
	c.JSON(http.StatusCreated, workout)
}

func GetWorkout(c *gin.Context) {
	// Logic to get workout from database
}

func UpdateWorkout(c *gin.Context) {
	// Logic to update workout in database
}

func DeleteWorkout(c *gin.Context) {
	// Logic to delete workout from database
}
