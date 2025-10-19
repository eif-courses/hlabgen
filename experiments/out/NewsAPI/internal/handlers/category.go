package handlers

import (
	"NewsAPI/internal/models"
	"encoding/json"
	"net/http"
)

var categories []models.Category

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	categories = append(categories, category)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}

func GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}
