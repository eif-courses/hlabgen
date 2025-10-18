package handlers

import (
	"MusicAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateLike(w http.ResponseWriter, r *http.Request) {
	var like models.Like
	if err := json.NewDecoder(r.Body).Decode(&like); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(like)
}

// Other CRUD functions for Like
