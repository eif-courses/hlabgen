package handlers

import (
	"TravelAPI/internal/models"
	"encoding/json"
	"net/http"
)

var bookings []models.Booking

func CreateBooking(w http.ResponseWriter, r *http.Request) {
	var booking models.Booking
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	bookings = append(bookings, booking)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(booking)
}

func GetBookings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookings)
}
