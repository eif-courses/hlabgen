package handlers

import (
	"MusicAPI/internal/models"
	"encoding/json"
	"net/http"
)

// CreateAlbum handles the creation of a new album.
func CreateAlbum() {
	var album models.Album
	if err := json.NewDecoder(r.Body).Decode(&album); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(album)
}

// GetAlbums handles fetching all albums.
func GetAlbums() {
	// Implementation for fetching albums
}
