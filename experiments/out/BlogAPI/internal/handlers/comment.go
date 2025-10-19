package handlers

import (
	"BlogAPI/internal/models"
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

func GetComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comments)
}
func GetComment(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a single comment
	w.WriteHeader(http.StatusOK)
}
func UpdateComment(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating a comment
	w.WriteHeader(http.StatusOK)
}
func DeleteComment(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting a comment
	w.WriteHeader(http.StatusNoContent)
}
