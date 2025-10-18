package handlers

import (
	"RecipeAPI/internal/models"
	"encoding/json"
	"net/http"
)

// CreateIngredient handles the creation of a new ingredient.
func CreateIngredient(w http.ResponseWriter, r *http.Request) {
	var ingredient models.Ingredient
	if err := json.NewDecoder(r.Body).Decode(&ingredient); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ingredient)
}

// GetIngredients handles fetching all ingredients.
func GetIngredients(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}
