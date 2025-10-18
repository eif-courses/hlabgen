package handlers

import (
	"BlogAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Here you would typically save the post to the database
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	// Here you would typically retrieve posts from the database
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode([]models.Post{})
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	// Logic to update a post
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	// Logic to delete a post
}
