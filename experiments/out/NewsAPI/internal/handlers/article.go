package handlers

import (
	"NewsAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateArticle() {
	var article models.Article
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(article)
}

func GetArticle() {
	// Implementation for getting an article
}
func UpdateArticle() {
	// Implementation for updating an article
}
func DeleteArticle() {
	// Implementation for deleting an article
}
