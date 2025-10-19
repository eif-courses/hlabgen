package handlers

import (
	"MusicAPI/internal/models"
	"encoding/json"
	"net/http"
)

var artists []models.Artist

func CreateArtist(w http.ResponseWriter, r *http.Request) {
	var artist models.Artist
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&artist); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	artists = append(artists, artist)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(artist)
}

func GetArtists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(artists)
}
func GetArtist(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a single artist
	w.WriteHeader(http.StatusOK)
}
func UpdateArtist(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating an artist
	w.WriteHeader(http.StatusOK)
}
func DeleteArtist(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting an artist
	w.WriteHeader(http.StatusNoContent)
}
