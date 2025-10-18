package handlers

import (
	"MovieAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateReview() {
	var review models.Review
	if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(review)
}

func GetReviews() {
	// Implementation for retrieving reviews
}
