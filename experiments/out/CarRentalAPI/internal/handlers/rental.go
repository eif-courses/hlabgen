package handlers

import (
	"CarRentalAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateRental() {
	var rental models.Rental
	if err := json.NewDecoder(r.Body).Decode(&rental); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(rental)
}

func GetRentals() {
	// Implementation for fetching rentals
}
func UpdateRental() {
	// Implementation for updating a rental
}
func DeleteRental() {
	// Implementation for deleting a rental
}
