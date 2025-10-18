package handlers

import (
	"MusicAPI/internal/models"
	"encoding/json"
	"net/http"
)

// CreateSong handles the creation of a new song.
func CreateSong() {
	var song models.Song
	if err := json.NewDecoder(r.Body).Decode(&song); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(song)
}

// GetSongs handles fetching all songs.
func GetSongs() {
	// Implementation for fetching songs
}
