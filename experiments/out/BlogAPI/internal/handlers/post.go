package handlers

import (
	"BlogAPI/internal/models"
	"encoding/json"
	"net/http"
)

var posts []models.Post

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	posts = append(posts, post)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
func GetPost(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a single post
	w.WriteHeader(http.StatusOK)
}
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating a post
	w.WriteHeader(http.StatusOK)
}
func DeletePost(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting a post
	w.WriteHeader(http.StatusNoContent)
}
