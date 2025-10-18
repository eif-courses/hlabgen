package handlers

import (
	"InventoryAPI/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateStock(c *gin.Context) {
	var stock models.Stock
	if err := c.ShouldBindJSON(&stock); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Logic to save stock
	c.JSON(http.StatusCreated, stock)
}

func GetStocks(c *gin.Context) {
	// Logic to get stocks
	c.JSON(http.StatusOK, []models.Stock{})
}

func UpdateStock(c *gin.Context) {
	// Logic to update stock
}

func DeleteStock(c *gin.Context) {
	// Logic to delete stock
}
