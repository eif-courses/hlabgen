package handlers

import (
	"SocialAPI/internal/models"
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
	// Implementation for getting a post
}
func UpdatePost() {
	// Implementation for updating a post
}
func DeletePost() {
	// Implementation for deleting a post
}
