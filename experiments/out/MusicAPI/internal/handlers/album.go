package handlers

import (
	"MusicAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateAlbum(w http.ResponseWriter, r *http.Request) {
	var album models.Album
	if err := json.NewDecoder(r.Body).Decode(&album); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(album)
}

// Other CRUD functions for Album
