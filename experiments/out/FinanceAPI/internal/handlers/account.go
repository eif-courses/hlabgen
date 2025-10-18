package handlers

import (
	"FinanceAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateAccount() {
	var account models.Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)
}

func GetAccounts() {
	// Implementation for getting accounts
}
func UpdateAccount() {
	// Implementation for updating an account
}
func DeleteAccount() {
	// Implementation for deleting an account
}
