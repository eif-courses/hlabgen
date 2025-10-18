package handlers

import (
	"MusicAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateSong() {
	var song models.Song
	if err := json.NewDecoder(r.Body).Decode(&song); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(song)
}

func GetSongs() {
	// Implementation for fetching songs
}
func UpdateSong() {
	// Implementation for updating a song
}
func DeleteSong() {
	// Implementation for deleting a song
}
