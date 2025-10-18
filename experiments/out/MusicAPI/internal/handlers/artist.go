package handlers

import (
	"MusicAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateArtist() {
	var artist models.Artist
	if err := json.NewDecoder(r.Body).Decode(&artist); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(artist)
}

func GetArtists() {
	// Implementation for fetching artists
}
func UpdateArtist() {
	// Implementation for updating an artist
}
func DeleteArtist() {
	// Implementation for deleting an artist
}
