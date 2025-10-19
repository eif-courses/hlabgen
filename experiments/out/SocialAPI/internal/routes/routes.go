package routes

import (
	"SocialAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/posts", handlers.CreatePost).Methods("POST")
	r.HandleFunc("/comments", handlers.CreateComment).Methods("POST")
}
