package handlers

import (
	"InventoryAPI/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateSupplier(c *gin.Context) {
	var supplier models.Supplier
	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Logic to save supplier
	c.JSON(http.StatusCreated, supplier)
}

func GetSuppliers(c *gin.Context) {
	// Logic to get suppliers
	c.JSON(http.StatusOK, []models.Supplier{})
}

func UpdateSupplier(c *gin.Context) {
	// Logic to update supplier
}

func DeleteSupplier(c *gin.Context) {
	// Logic to delete supplier
}
