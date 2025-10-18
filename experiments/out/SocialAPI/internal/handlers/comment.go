package handlers

import (
	"SocialAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateComment() {
	var comment models.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(comment)
}

func GetComment() {
	// Implementation for getting a comment
}
func UpdateComment() {
	// Implementation for updating a comment
}
func DeleteComment() {
	// Implementation for deleting a comment
}
