package handlers

import (
	"MovieAPI/internal/models"
	"encoding/json"
	"net/http"
)

// CreateDirector handles the creation of a new director.
func CreateDirector() {
	var director models.Director
	if err := json.NewDecoder(r.Body).Decode(&director); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(director)
}

// GetDirectors handles retrieving all directors.
func GetDirectors() {
	// Implementation for retrieving directors
}
