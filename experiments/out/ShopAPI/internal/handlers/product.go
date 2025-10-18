package handlers

import (
	"ShopAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateProduct() {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func GetProducts() {
	// Implementation for getting products
}
func UpdateProduct() {
	// Implementation for updating a product
}
func DeleteProduct() {
	// Implementation for deleting a product
}
