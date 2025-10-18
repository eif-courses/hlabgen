package handlers

import (
	"EventAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateAttendee(w http.ResponseWriter, r *http.Request) {
	var attendee models.Attendee
	if err := json.NewDecoder(r.Body).Decode(&attendee); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(attendee)
}

func GetAttendee(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}

func UpdateAttendee(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}

func DeleteAttendee(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}
