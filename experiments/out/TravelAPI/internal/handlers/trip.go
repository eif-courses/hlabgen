package handlers

import (
	"TravelAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateTrip() {
	var trip models.Trip
	if err := json.NewDecoder(r.Body).Decode(&trip); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(trip)
}

func GetTrip() {
	// Implementation here
}
func UpdateTrip() {
	// Implementation here
}
func DeleteTrip() {
	// Implementation here
}
