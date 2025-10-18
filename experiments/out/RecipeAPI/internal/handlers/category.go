package handlers

import (
	"RecipeAPI/internal/models"
	"encoding/json"
	"net/http"
)

// CreateCategory handles the creation of a new category.
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}

// GetCategories handles fetching all categories.
func GetCategories(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}
