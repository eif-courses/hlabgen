package handlers

import (
	"EventAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateOrganizer(w http.ResponseWriter, r *http.Request) {
	var organizer models.Organizer
	if err := json.NewDecoder(r.Body).Decode(&organizer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(organizer)
}

func GetOrganizer(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}

func UpdateOrganizer(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}

func DeleteOrganizer(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}
