package handlers

import (
	"InventoryAPI/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateWarehouse(c *gin.Context) {
	var warehouse models.Warehouse
	if err := c.ShouldBindJSON(&warehouse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Logic to save warehouse
	c.JSON(http.StatusCreated, warehouse)
}

func GetWarehouses(c *gin.Context) {
	// Logic to get warehouses
	c.JSON(http.StatusOK, []models.Warehouse{})
}

func UpdateWarehouse(c *gin.Context) {
	// Logic to update warehouse
}

func DeleteWarehouse(c *gin.Context) {
	// Logic to delete warehouse
}
