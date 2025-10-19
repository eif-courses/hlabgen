package handlers

import (
	"TravelAPI/internal/models"
	"encoding/json"
	"net/http"
)

var trips []models.Trip

func CreateTrip(w http.ResponseWriter, r *http.Request) {
	var trip models.Trip
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&trip); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	trips = append(trips, trip)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(trip)
}

func GetTrips(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trips)
}
