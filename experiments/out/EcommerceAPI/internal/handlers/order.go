package handlers

import (
	"EcommerceAPI/internal/models"
	"encoding/json"
	"net/http"
)

var orders []models.Order

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	orders = append(orders, order)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}
func GetOrder(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a single order
	w.WriteHeader(http.StatusOK)
}
func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating an order
	w.WriteHeader(http.StatusOK)
}
func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting an order
	w.WriteHeader(http.StatusNoContent)
}
