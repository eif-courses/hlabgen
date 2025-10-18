package handlers

import (
	"TaskAPI/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Save task to database
	c.JSON(http.StatusCreated, task)
}

func GetTasks(c *gin.Context) {
	// Retrieve tasks from database
	c.JSON(http.StatusOK, []models.Task{})
}

func UpdateTask(c *gin.Context) {
	// Update task in database
	c.JSON(http.StatusOK, gin.H{"message": "Task updated"})
}

func DeleteTask(c *gin.Context) {
	// Delete task from database
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
