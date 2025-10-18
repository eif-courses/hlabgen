package handlers

import (
	"SocialAPI/internal/models"
	"encoding/json"
	"net/http"
)

// CreateComment handles the creation of a new comment.
func CreateComment(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(comment)
}

// GetComment handles fetching a comment by ID.
func GetComment(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}

// UpdateComment handles updating a comment.
func UpdateComment(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}

// DeleteComment handles deleting a comment.
func DeleteComment(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}
