package handlers

import (
	"MusicAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateArtist(w http.ResponseWriter, r *http.Request) {
	var artist models.Artist
	if err := json.NewDecoder(r.Body).Decode(&artist); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(artist)
}

// Other CRUD functions for Artist
