package handlers

import (
	"TaskAPI/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateTeam(c *gin.Context) {
	var team models.Team
	if err := c.ShouldBindJSON(&team); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Save team to database
	c.JSON(http.StatusCreated, team)
}

func GetTeams(c *gin.Context) {
	// Retrieve teams from database
	c.JSON(http.StatusOK, []models.Team{})
}

func UpdateTeam(c *gin.Context) {
	// Update team in database
	c.JSON(http.StatusOK, gin.H{"message": "Team updated"})
}

func DeleteTeam(c *gin.Context) {
	// Delete team from database
	c.JSON(http.StatusOK, gin.H{"message": "Team deleted"})
}
