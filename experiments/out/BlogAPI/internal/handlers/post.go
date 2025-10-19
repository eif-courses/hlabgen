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
