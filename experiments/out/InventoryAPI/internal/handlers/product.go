package handlers

import (
	"InventoryAPI/internal/models"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

// CreateProduct handles the creation of a new product.
func CreateProduct() {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	product.ID = primitive.NewObjectID()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

// GetProducts handles fetching all products.
func GetProducts() {
	// Implementation for fetching products
}

// UpdateProduct handles updating an existing product.
func UpdateProduct() {
	// Implementation for updating a product
}

// DeleteProduct handles deleting a product.
func DeleteProduct() {
	// Implementation for deleting a product
}
