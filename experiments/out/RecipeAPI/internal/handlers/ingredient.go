package handlers

import (
	"RecipeAPI/internal/models"
	"encoding/json"
	"net/http"
)

var ingredients []models.Ingredient

func CreateIngredient(w http.ResponseWriter, r *http.Request) {
	var ingredient models.Ingredient
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&ingredient); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ingredients = append(ingredients, ingredient)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ingredient)
}

func GetIngredients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ingredients)
}
func GetIngredient(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a single ingredient
	w.WriteHeader(http.StatusOK)
}
func UpdateIngredient(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating an ingredient
	w.WriteHeader(http.StatusOK)
}
func DeleteIngredient(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting an ingredient
	w.WriteHeader(http.StatusNoContent)
}
