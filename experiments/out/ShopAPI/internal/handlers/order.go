package handlers

import (
	"ShopAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateOrder() {
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

func GetOrders() {
	// Implementation for getting orders
}
func UpdateOrder() {
	// Implementation for updating an order
}
func DeleteOrder() {
	// Implementation for deleting an order
}
