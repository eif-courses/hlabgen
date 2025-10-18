package handlers

import (
	"BlogAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreatePost() {
	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func GetPost() {
	// Implementation here
}
func UpdatePost() {
	// Implementation here
}
func DeletePost() {
	// Implementation here
}
