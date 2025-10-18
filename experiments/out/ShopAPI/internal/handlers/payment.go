package handlers

import (
	"ShopAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreatePayment() {
	var payment models.Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(payment)
}

func GetPayments() {
	// Implementation for getting payments
}
func UpdatePayment() {
	// Implementation for updating a payment
}
func DeletePayment() {
	// Implementation for deleting a payment
}
