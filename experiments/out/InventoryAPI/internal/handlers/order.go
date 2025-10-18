package handlers

import (
	"InventoryAPI/internal/models"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

// CreateOrder handles the creation of a new order.
func CreateOrder() {
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	order.ID = primitive.NewObjectID()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

// GetOrders handles fetching all orders.
func GetOrders() {
	// Implementation for fetching orders
}

// UpdateOrder handles updating an existing order.
func UpdateOrder() {
	// Implementation for updating an order
}

// DeleteOrder handles deleting an order.
func DeleteOrder() {
	// Implementation for deleting an order
}
