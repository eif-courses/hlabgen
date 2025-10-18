package handlers

import (
	"FitnessAPI/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateExercise(c *gin.Context) {
	var exercise models.Exercise
	if err := c.ShouldBindJSON(&exercise); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Logic to save exercise to database
	c.JSON(http.StatusCreated, exercise)
}

func GetExercise(c *gin.Context) {
	// Logic to get exercise from database
}

func UpdateExercise(c *gin.Context) {
	// Logic to update exercise in database
}

func DeleteExercise(c *gin.Context) {
	// Logic to delete exercise from database
}
