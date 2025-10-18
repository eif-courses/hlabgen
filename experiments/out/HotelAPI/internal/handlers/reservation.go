package handlers

import (
	"HotelAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateReservation() {
	var reservation models.Reservation
	if err := json.NewDecoder(r.Body).Decode(&reservation); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(reservation)
}

func GetReservations() {
	// Implementation for fetching reservations
}
func UpdateReservation() {
	// Implementation for updating a reservation
}
func DeleteReservation() {
	// Implementation for deleting a reservation
}
