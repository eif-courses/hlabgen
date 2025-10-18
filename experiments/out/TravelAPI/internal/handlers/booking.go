package handlers

import (
	"TravelAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateBooking() {
	var booking models.Booking
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(booking)
}

func GetBooking() {
	// Implementation here
}
func UpdateBooking() {
	// Implementation here
}
func DeleteBooking() {
	// Implementation here
}
