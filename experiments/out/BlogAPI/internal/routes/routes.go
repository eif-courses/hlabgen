package routes

import (
	"BlogAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	// Post routes
	r.HandleFunc("/posts", handlers.CreatePost).Methods("POST")
	r.HandleFunc("/posts", handlers.GetPosts).Methods("GET")
	r.HandleFunc("/posts/{id}", handlers.GetPost).Methods("GET")
	r.HandleFunc("/posts/{id}", handlers.UpdatePost).Methods("PUT")
	r.HandleFunc("/posts/{id}", handlers.DeletePost).Methods("DELETE")
	// Comment routes
	r.HandleFunc("/comments", handlers.CreateComment).Methods("POST")
	r.HandleFunc("/comments", handlers.GetComments).Methods("GET")
	r.HandleFunc("/comments/{id}", handlers.GetComment).Methods("GET")
	r.HandleFunc("/comments/{id}", handlers.UpdateComment).Methods("PUT")
	r.HandleFunc("/comments/{id}", handlers.DeleteComment).Methods("DELETE")
	// Author routes
	r.HandleFunc("/authors", handlers.CreateAuthor).Methods("POST")
	r.HandleFunc("/authors", handlers.GetAuthors).Methods("GET")
	r.HandleFunc("/authors/{id}", handlers.GetAuthor).Methods("GET")
	r.HandleFunc("/authors/{id}", handlers.UpdateAuthor).Methods("PUT")
	r.HandleFunc("/authors/{id}", handlers.DeleteAuthor).Methods("DELETE")
}
