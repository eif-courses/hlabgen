package handlers

import (
	"FitnessAPI/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateProgress(c *gin.Context) {
	var progress models.Progress
	if err := c.ShouldBindJSON(&progress); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Logic to save progress to database
	c.JSON(http.StatusCreated, progress)
}

func GetProgress(c *gin.Context) {
	// Logic to get progress from database
}

func UpdateProgress(c *gin.Context) {
	// Logic to update progress in database
}

func DeleteProgress(c *gin.Context) {
	// Logic to delete progress from database
}
