package handlers

import (
	"TaskAPI/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateProject(c *gin.Context) {
	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Save project to database
	c.JSON(http.StatusCreated, project)
}

func GetProjects(c *gin.Context) {
	// Retrieve projects from database
	c.JSON(http.StatusOK, []models.Project{})
}

func UpdateProject(c *gin.Context) {
	// Update project in database
	c.JSON(http.StatusOK, gin.H{"message": "Project updated"})
}

func DeleteProject(c *gin.Context) {
	// Delete project from database
	c.JSON(http.StatusOK, gin.H{"message": "Project deleted"})
}
