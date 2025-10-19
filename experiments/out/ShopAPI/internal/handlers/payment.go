package handlers

import (
	"ShopAPI/internal/models"
	"encoding/json"
	"net/http"
)

var payments []models.Payment

func CreatePayment(w http.ResponseWriter, r *http.Request) {
	var payment models.Payment
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	payments = append(payments, payment)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(payment)
}
