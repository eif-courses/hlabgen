package handlers

import (
	"SocialAPI/internal/models"
	"encoding/json"
	"net/http"
)

var comments []models.Comment

func CreateComment(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	comments = append(comments, comment)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(comment)
}
