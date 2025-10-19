package handlers

import (
	"CarRentalAPI/internal/models"
	"encoding/json"
	"net/http"
)

var rentals []models.Rental

func CreateRental(w http.ResponseWriter, r *http.Request) {
	var rental models.Rental
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&rental); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	rentals = append(rentals, rental)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(rental)
}

func GetRentals(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rentals)
}
func GetRental(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a single rental
	w.WriteHeader(http.StatusOK)
}
func UpdateRental(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating a rental
	w.WriteHeader(http.StatusOK)
}
func DeleteRental(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting a rental
	w.WriteHeader(http.StatusNoContent)
}
