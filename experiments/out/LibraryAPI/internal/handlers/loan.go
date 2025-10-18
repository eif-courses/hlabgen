package handlers

import (
	"LibraryAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateLoan() {
	var loan models.Loan
	if err := json.NewDecoder(r.Body).Decode(&loan); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(loan)
}

func GetLoans() {
	// Implementation for retrieving loans
}
func UpdateLoan() {
	// Implementation for updating a loan
}
func DeleteLoan() {
	// Implementation for deleting a loan
}
