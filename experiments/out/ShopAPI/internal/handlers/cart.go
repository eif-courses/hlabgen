package handlers

import (
	"ShopAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateCart() {
	var cart models.Cart
	if err := json.NewDecoder(r.Body).Decode(&cart); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cart)
}

func GetCarts() {
	// Implementation for getting carts
}
func UpdateCart() {
	// Implementation for updating a cart
}
func DeleteCart() {
	// Implementation for deleting a cart
}
