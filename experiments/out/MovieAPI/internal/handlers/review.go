package handlers

import (
	"MovieAPI/internal/models"
	"encoding/json"
	"net/http"
)

// CreateReview handles the creation of a new review.
func CreateReview() {
	var review models.Review
	if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(review)
}

// GetReviews handles retrieving all reviews for a movie.
func GetReviews() {
	// Implementation for retrieving reviews
}
