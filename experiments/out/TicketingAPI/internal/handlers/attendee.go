package handlers

import (
	"TicketingAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateAttendee() {
	var attendee models.Attendee
	if err := json.NewDecoder(r.Body).Decode(&attendee); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(attendee)
}
