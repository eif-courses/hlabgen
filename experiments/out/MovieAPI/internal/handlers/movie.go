package handlers

import (
	"MovieAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateMovie() {
	var movie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(movie)
}

func GetMovies() {
	// Implementation for retrieving movies
}
func UpdateMovie() {
	// Implementation for updating a movie
}
func DeleteMovie() {
	// Implementation for deleting a movie
}
