package handlers

import (
	"RecipeAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateRecipe() {
	var recipe models.Recipe
	if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(recipe)
}

func GetRecipes() {
	// Implementation for getting recipes
}
func UpdateRecipe() {
	// Implementation for updating a recipe
}
func DeleteRecipe() {
	// Implementation for deleting a recipe
}
