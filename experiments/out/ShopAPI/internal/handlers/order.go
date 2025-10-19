package handlers

import (
	"ShopAPI/internal/models"
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
