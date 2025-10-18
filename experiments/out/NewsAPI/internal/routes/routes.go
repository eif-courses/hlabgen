package routes

import (
	"NewsAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register() {
	r.HandleFunc("/articles", handlers.CreateArticle).Methods("POST")
	r.HandleFunc("/articles", handlers.GetArticles).Methods("GET")
	r.HandleFunc("/articles/{id}", handlers.UpdateArticle).Methods("PUT")
	r.HandleFunc("/articles/{id}", handlers.DeleteArticle).Methods("DELETE")
	r.HandleFunc("/categories", handlers.CreateCategory).Methods("POST")
	r.HandleFunc("/categories", handlers.GetCategories).Methods("GET")
	r.HandleFunc("/categories/{id}", handlers.UpdateCategory).Methods("PUT")
	r.HandleFunc("/categories/{id}", handlers.DeleteCategory).Methods("DELETE")
	r.HandleFunc("/authors", handlers.CreateAuthor).Methods("POST")
	r.HandleFunc("/authors", handlers.GetAuthors).Methods("GET")
	r.HandleFunc("/authors/{id}", handlers.UpdateAuthor).Methods("PUT")
	r.HandleFunc("/authors/{id}", handlers.DeleteAuthor).Methods("DELETE")
}
