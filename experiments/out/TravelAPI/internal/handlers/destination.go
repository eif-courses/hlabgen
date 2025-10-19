package handlers

import (
	"TravelAPI/internal/models"
	"encoding/json"
	"net/http"
)

var destinations []models.Destination

func CreateDestination(w http.ResponseWriter, r *http.Request) {
	var destination models.Destination
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&destination); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	destinations = append(destinations, destination)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(destination)
}

func GetDestinations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(destinations)
}
