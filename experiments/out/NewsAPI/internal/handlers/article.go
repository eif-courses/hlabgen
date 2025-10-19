package handlers

import (
	"NewsAPI/internal/models"
	"encoding/json"
	"net/http"
)

var articles []models.Article

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	var article models.Article
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	articles = append(articles, article)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(article)
}

func GetArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}
