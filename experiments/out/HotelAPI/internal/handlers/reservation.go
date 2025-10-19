package handlers

import (
	"HotelAPI/internal/models"
	"encoding/json"
	"net/http"
)

var reservations []models.Reservation

func CreateReservation(w http.ResponseWriter, r *http.Request) {
	var reservation models.Reservation
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&reservation); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	reservations = append(reservations, reservation)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(reservation)
}
