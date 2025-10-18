package handlers

import (
	"LibraryAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateLoan(w http.ResponseWriter, r *http.Request) {
	var loan models.Loan
	if err := json.NewDecoder(r.Body).Decode(&loan); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(loan)
}

func GetLoans(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting all loans
}
