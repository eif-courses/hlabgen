package handlers

import (
	"InventoryAPI/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Logic to save product
	c.JSON(http.StatusCreated, product)
}

func GetProducts(c *gin.Context) {
	// Logic to get products
	c.JSON(http.StatusOK, []models.Product{})
}

func UpdateProduct(c *gin.Context) {
	// Logic to update product
}

func DeleteProduct(c *gin.Context) {
	// Logic to delete product
}
