package handlers

import (
	"BlogAPI/internal/models"
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

// Additional comment handlers would go here
