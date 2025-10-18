package handlers

import (
	"NewsAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateAuthor() {
	var author models.Author
	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(author)
}

func GetAuthors() {
	// Implementation for fetching authors
}
func UpdateAuthor() {
	// Implementation for updating an author
}
func DeleteAuthor() {
	// Implementation for deleting an author
}
