package handlers

import (
	"ShopAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	// Implementation for fetching orders
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating an order
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting an order
}
