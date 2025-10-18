package handlers

import (
	"LibraryAPI/internal/models"
	"encoding/json"
	"net/http"
)

var loans []models.Loan

func CreateLoan(w http.ResponseWriter, r *http.Request) {
	var loan models.Loan
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&loan); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	loans = append(loans, loan)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(loan)
}

func GetLoans(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(loans)
}

func UpdateLoan(w http.ResponseWriter, r *http.Request) {
	// Update logic here
}

func DeleteLoan(w http.ResponseWriter, r *http.Request) {
	// Delete logic here
}
