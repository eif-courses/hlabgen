package handlers

import (
	"MovieAPI/internal/models"
	"encoding/json"
	"net/http"
)

var directors []models.Director

func CreateDirector(w http.ResponseWriter, r *http.Request) {
	var director models.Director
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&director); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	directors = append(directors, director)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(director)
}

func GetDirectors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(directors)
}
