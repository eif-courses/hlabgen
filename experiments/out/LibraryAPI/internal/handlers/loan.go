package handlers

import (
	"LibraryAPI/internal/models"
	"encoding/json"
	"net/http"
)

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
	// Add logic to save loan to database
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(loan)
}

func GetLoans(w http.ResponseWriter, r *http.Request) {
	// Add logic to retrieve loans from database
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode([]models.Loan{})
}

func UpdateLoan(w http.ResponseWriter, r *http.Request) {
	// Add logic to update loan in database
	w.WriteHeader(http.StatusOK)
}

func DeleteLoan(w http.ResponseWriter, r *http.Request) {
	// Add logic to delete loan from database
	w.WriteHeader(http.StatusNoContent)
}
