package handlers

import (
	"MusicAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateAlbum() {
	var album models.Album
	if err := json.NewDecoder(r.Body).Decode(&album); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(album)
}

func GetAlbums() {
	// Implementation for fetching albums
}
func UpdateAlbum() {
	// Implementation for updating an album
}
func DeleteAlbum() {
	// Implementation for deleting an album
}
