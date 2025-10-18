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
	// Logic to save recipe to database
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(recipe)
}

func GetRecipes() {
	// Logic to retrieve recipes from database
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode([]models.Recipe{})
}
func UpdateRecipe() {
	// Logic to update recipe in database
}
func DeleteRecipe() {
	// Logic to delete recipe from database
}
