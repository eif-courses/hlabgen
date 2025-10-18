package handlers

import (
	"LibraryAPI/internal/models"
	"encoding/json"
	"net/http"
)

// CreateLoan handles the creation of a new loan.
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
	// Logic to save loan to database
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(loan)
}

// GetLoans handles fetching all loans.
func GetLoans(w http.ResponseWriter, r *http.Request) {
	// Logic to fetch loans from database
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode([]models.Loan{})
}

// UpdateLoan handles updating a loan.
func UpdateLoan(w http.ResponseWriter, r *http.Request) {
	var loan models.Loan
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&loan); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Logic to update loan in database
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(loan)
}

// DeleteLoan handles deleting a loan.
func DeleteLoan(w http.ResponseWriter, r *http.Request) {
	// Logic to delete loan from database
	w.WriteHeader(http.StatusNoContent)
}
