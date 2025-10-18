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
	// Implementation for getting a trip
}
func UpdateTrip() {
	// Implementation for updating a trip
}
func DeleteTrip() {
	// Implementation for deleting a trip
}
