package handlers

import (
	"EventAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateVenue(w http.ResponseWriter, r *http.Request) {
	var venue models.Venue
	if err := json.NewDecoder(r.Body).Decode(&venue); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(venue)
}

func GetVenue(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}

func UpdateVenue(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}

func DeleteVenue(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}
