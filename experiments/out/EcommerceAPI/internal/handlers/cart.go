package handlers

import (
	"EcommerceAPI/internal/models"
	"encoding/json"
	"net/http"
)

var carts []models.Cart

func CreateCart(w http.ResponseWriter, r *http.Request) {
	var cart models.Cart
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&cart); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	carts = append(carts, cart)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cart)
}

func GetCarts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(carts)
}
func GetCart(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a single cart
	w.WriteHeader(http.StatusOK)
}
func UpdateCart(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating a cart
	w.WriteHeader(http.StatusOK)
}
func DeleteCart(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting a cart
	w.WriteHeader(http.StatusNoContent)
}
