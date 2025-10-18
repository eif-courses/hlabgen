package handlers

import (
	"ShopAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateCustomer() {
	var customer models.Customer
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(customer)
}

func GetCustomers() {
	// Implementation for getting customers
}
func UpdateCustomer() {
	// Implementation for updating a customer
}
func DeleteCustomer() {
	// Implementation for deleting a customer
}
