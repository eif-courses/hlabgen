package handlers

import (
	"TicketingAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateVenue() {
	var venue models.Venue
	if err := json.NewDecoder(r.Body).Decode(&venue); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(venue)
}
