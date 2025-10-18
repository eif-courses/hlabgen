package handlers

import (
	"MusicAPI/internal/models"
	"encoding/json"
	"net/http"
)

// CreateArtist handles the creation of a new artist.
func CreateArtist() {
	var artist models.Artist
	if err := json.NewDecoder(r.Body).Decode(&artist); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(artist)
}

// GetArtists handles fetching all artists.
func GetArtists() {
	// Implementation for fetching artists
}
