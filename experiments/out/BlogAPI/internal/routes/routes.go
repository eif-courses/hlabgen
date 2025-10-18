package routes

import (
	"BlogAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register() {
	r.HandleFunc("/posts", handlers.CreatePost).Methods("POST")
	r.HandleFunc("/posts/{id:[0-9]+}", handlers.GetPost).Methods("GET")
	r.HandleFunc("/posts/{id:[0-9]+}", handlers.UpdatePost).Methods("PUT")
	r.HandleFunc("/posts/{id:[0-9]+}", handlers.DeletePost).Methods("DELETE")
	r.HandleFunc("/comments", handlers.CreateComment).Methods("POST")
	r.HandleFunc("/comments/{id:[0-9]+}", handlers.GetComment).Methods("GET")
	r.HandleFunc("/comments/{id:[0-9]+}", handlers.UpdateComment).Methods("PUT")
	r.HandleFunc("/comments/{id:[0-9]+}", handlers.DeleteComment).Methods("DELETE")
	r.HandleFunc("/authors", handlers.CreateAuthor).Methods("POST")
	r.HandleFunc("/authors/{id:[0-9]+}", handlers.GetAuthor).Methods("GET")
	r.HandleFunc("/authors/{id:[0-9]+}", handlers.UpdateAuthor).Methods("PUT")
	r.HandleFunc("/authors/{id:[0-9]+}", handlers.DeleteAuthor).Methods("DELETE")
}
