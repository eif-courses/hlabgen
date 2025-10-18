package handlers

import (
	"InventoryAPI/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Logic to save order
	c.JSON(http.StatusCreated, order)
}

func GetOrders(c *gin.Context) {
	// Logic to get orders
	c.JSON(http.StatusOK, []models.Order{})
}

func UpdateOrder(c *gin.Context) {
	// Logic to update order
}

func DeleteOrder(c *gin.Context) {
	// Logic to delete order
}
