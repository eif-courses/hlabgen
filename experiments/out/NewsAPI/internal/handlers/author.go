package handlers

import (
	"NewsAPI/internal/models"
	"encoding/json"
	"net/http"
)

var authors []models.Author

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	authors = append(authors, author)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(author)
}

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authors)
}
