package routes

import (
	"SocialAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register() {
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")
	r.HandleFunc("/posts", handlers.CreatePost).Methods("POST")
	r.HandleFunc("/posts/{id:[0-9]+}", handlers.GetPost).Methods("GET")
	r.HandleFunc("/posts/{id:[0-9]+}", handlers.UpdatePost).Methods("PUT")
	r.HandleFunc("/posts/{id:[0-9]+}", handlers.DeletePost).Methods("DELETE")
	r.HandleFunc("/comments", handlers.CreateComment).Methods("POST")
	r.HandleFunc("/comments/{id:[0-9]+}", handlers.GetComment).Methods("GET")
	r.HandleFunc("/comments/{id:[0-9]+}", handlers.UpdateComment).Methods("PUT")
	r.HandleFunc("/comments/{id:[0-9]+}", handlers.DeleteComment).Methods("DELETE")
}
