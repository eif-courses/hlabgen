package handlers

import (
	"ResearchAPI/internal/models"
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

func GetReview() {
	// Implementation here
}
func UpdateReview() {
	// Implementation here
}
func DeleteReview() {
	// Implementation here
}
