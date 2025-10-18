package handlers

import (
	"EcommerceAPI/internal/models"
	"encoding/json"
	"net/http"
)

var products []models.Product

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	products = append(products, product)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
func GetProduct(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a single product
	w.WriteHeader(http.StatusOK)
}
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating a product
	w.WriteHeader(http.StatusOK)
}
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting a product
	w.WriteHeader(http.StatusNoContent)
}
