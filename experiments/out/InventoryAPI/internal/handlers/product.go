package handlers

import (
	"InventoryAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}
