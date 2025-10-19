package handlers

import (
	"FinanceAPI/internal/models"
	"encoding/json"
	"net/http"
)

var transactions []models.Transaction

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction models.Transaction
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	transactions = append(transactions, transaction)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(transaction)
}
