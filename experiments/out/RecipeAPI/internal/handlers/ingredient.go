package handlers

import (
	"RecipeAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateIngredient() {
	var ingredient models.Ingredient
	if err := json.NewDecoder(r.Body).Decode(&ingredient); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ingredient)
}

func GetIngredients() {
	// Implementation for getting ingredients
}
func UpdateIngredient() {
	// Implementation for updating an ingredient
}
func DeleteIngredient() {
	// Implementation for deleting an ingredient
}
