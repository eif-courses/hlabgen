package routes

import (
	"BlogAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register() {
	r.HandleFunc("/posts", handlers.CreatePost).Methods("POST")
	r.HandleFunc("/posts", handlers.GetPosts).Methods("GET")
	r.HandleFunc("/posts/{id}", handlers.UpdatePost).Methods("PUT")
	r.HandleFunc("/posts/{id}", handlers.DeletePost).Methods("DELETE")
	r.HandleFunc("/comments", handlers.CreateComment).Methods("POST")
	r.HandleFunc("/authors", handlers.CreateAuthor).Methods("POST")
}
