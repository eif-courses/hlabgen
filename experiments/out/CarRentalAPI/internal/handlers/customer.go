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

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
func GetCustomer(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a single customer
	w.WriteHeader(http.StatusOK)
}
func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating a customer
	w.WriteHeader(http.StatusOK)
}
func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting a customer
	w.WriteHeader(http.StatusNoContent)
}
