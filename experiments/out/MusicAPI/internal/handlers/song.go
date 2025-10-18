package handlers

import (
	"MusicAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateSong(w http.ResponseWriter, r *http.Request) {
	var song models.Song
	if err := json.NewDecoder(r.Body).Decode(&song); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(song)
}

func GetSong(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}

func UpdateSong(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}

func DeleteSong(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}
