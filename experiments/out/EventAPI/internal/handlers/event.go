package handlers

import (
	"EventAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event models.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(event)
}

func GetEvent(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}
