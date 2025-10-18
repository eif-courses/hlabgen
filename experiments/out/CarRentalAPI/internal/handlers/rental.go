package handlers

import (
	"CarRentalAPI/internal/models"
	"encoding/json"
	"net/http"
)

// CreateRental handles the creation of a new rental.
func CreateRental() {
	var rental models.Rental
	if err := json.NewDecoder(r.Body).Decode(&rental); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(rental)
}

// GetRentals handles fetching all rentals.
func GetRentals() {
	// Implementation for fetching rentals
}

// UpdateRental handles updating an existing rental.
func UpdateRental() {
	// Implementation for updating a rental
}

// DeleteRental handles deleting a rental.
func DeleteRental() {
	// Implementation for deleting a rental
}
