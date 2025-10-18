package handlers

import (
	"BlogAPI/internal/models"
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

// Additional author handlers would go here
