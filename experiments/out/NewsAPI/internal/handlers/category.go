package handlers

import (
	"NewsAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateCategory() {
	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}

func GetCategories() {
	// Implementation for fetching categories
}
func UpdateCategory() {
	// Implementation for updating a category
}
func DeleteCategory() {
	// Implementation for deleting a category
}
