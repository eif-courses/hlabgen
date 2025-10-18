package handlers

import (
    "encoding/json"
    "net/http"
    "HotelAPI/internal/models"
)

func CreateReview(w http.ResponseWriter, r *http.Request) {
    var review models.Review
    if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(review)
}

func GetReviews(w http.ResponseWriter, r *http.Request) {
    // Implementation for fetching reviews
}

func UpdateReview(w http.ResponseWriter, r *http.Request) {
    // Implementation for updating a review
}

func DeleteReview(w http.ResponseWriter, r *http.Request) {
    // Implementation for deleting a review
}