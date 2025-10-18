package handlers

import (
	"FinanceAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateTransaction() {
	var transaction models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(transaction)
}

func GetTransactions() {
	// Implementation for getting transactions
}
func UpdateTransaction() {
	// Implementation for updating a transaction
}
func DeleteTransaction() {
	// Implementation for deleting a transaction
}
