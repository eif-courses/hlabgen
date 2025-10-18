package handlers

import (
	"BlogAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author
	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Here you would typically save the author to the database
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(author)
}

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	// Logic to retrieve authors
}
