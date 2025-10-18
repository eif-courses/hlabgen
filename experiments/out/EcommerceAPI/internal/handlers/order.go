package handlers

import (
	"EcommerceAPI/internal/models"
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
