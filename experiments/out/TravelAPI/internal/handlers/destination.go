package handlers

import (
	"TravelAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateDestination() {
	var destination models.Destination
	if err := json.NewDecoder(r.Body).Decode(&destination); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(destination)
}

func GetDestination() {
	// Implementation for getting a destination
}
func UpdateDestination() {
	// Implementation for updating a destination
}
func DeleteDestination() {
	// Implementation for deleting a destination
}
