package routes

import (
	"NewsAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	// Article routes
	r.HandleFunc("/articles", handlers.CreateArticle).Methods("POST")
	r.HandleFunc("/articles", handlers.GetArticles).Methods("GET")
	// Category routes
	r.HandleFunc("/categories", handlers.CreateCategory).Methods("POST")
	r.HandleFunc("/categories", handlers.GetCategories).Methods("GET")
	// Author routes
	r.HandleFunc("/authors", handlers.CreateAuthor).Methods("POST")
	r.HandleFunc("/authors", handlers.GetAuthors).Methods("GET")
}
