package handlers

import (
	"RecipeAPI/internal/models"
	"encoding/json"
	"net/http"
)

var recipes []models.Recipe

func CreateRecipe(w http.ResponseWriter, r *http.Request) {
	var recipe models.Recipe
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	recipes = append(recipes, recipe)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(recipe)
}

func GetRecipes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipes)
}
func GetRecipe(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a single recipe
	w.WriteHeader(http.StatusOK)
}
func UpdateRecipe(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating a recipe
	w.WriteHeader(http.StatusOK)
}
func DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting a recipe
	w.WriteHeader(http.StatusNoContent)
}
