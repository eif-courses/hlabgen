package handlers

import (
	"CarRentalAPI/internal/models"
	"encoding/json"
	"net/http"
)

var customers []models.Customer

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	customers = append(customers, customer)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(customer)
}
