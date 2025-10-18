package handlers

import (
	"RecipeAPI/internal/models"
	"encoding/json"
	"net/http"
)

// CreateRecipe handles the creation of a new recipe.
func CreateRecipe(w http.ResponseWriter, r *http.Request) {
	var recipe models.Recipe
	if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(recipe)
}

// GetRecipes handles fetching all recipes.
func GetRecipes(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}

// UpdateRecipe handles updating an existing recipe.
func UpdateRecipe(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}

// DeleteRecipe handles deleting a recipe.
func DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}
