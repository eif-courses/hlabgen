package handlers

import (
	"TicketingAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateEvent() {
	var event models.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(event)
}

func GetEvents() {
	// Implementation for getting events
}
func UpdateEvent() {
	// Implementation for updating an event
}
func DeleteEvent() {
	// Implementation for deleting an event
}
