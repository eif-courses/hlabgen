package routes

import (
	"BlogAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	// Post routes
	r.HandleFunc("/posts", handlers.CreatePost).Methods("POST")
	r.HandleFunc("/posts", handlers.GetPosts).Methods("GET")
	// Comment routes
	r.HandleFunc("/comments", handlers.CreateComment).Methods("POST")
	r.HandleFunc("/comments", handlers.GetComments).Methods("GET")
	// Author routes
	r.HandleFunc("/authors", handlers.CreateAuthor).Methods("POST")
	r.HandleFunc("/authors", handlers.GetAuthors).Methods("GET")
}
