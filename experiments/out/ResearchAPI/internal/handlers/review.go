package handlers

import (
	"ResearchAPI/internal/models"
	"encoding/json"
	"net/http"
)

var reviews []models.Review

func CreateReview(w http.ResponseWriter, r *http.Request) {
	var review models.Review
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	reviews = append(reviews, review)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(review)
}
