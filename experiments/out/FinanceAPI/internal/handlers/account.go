package handlers

import (
	"FinanceAPI/internal/models"
	"encoding/json"
	"net/http"
)

var accounts []models.Account

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account models.Account
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	accounts = append(accounts, account)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)
}
