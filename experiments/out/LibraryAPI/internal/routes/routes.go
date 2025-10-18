package routes

import (
	"LibraryAPI/internal/handlers"
	"github.com/gorilla/mux"
)

// Register sets up the routes for the API.
func Register(router *mux.Router) {
	router.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	router.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")

	router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	router.HandleFunc("/loans", handlers.CreateLoan).Methods("POST")
	router.HandleFunc("/loans", handlers.GetLoans).Methods("GET")
	router.HandleFunc("/loans/{id}", handlers.UpdateLoan).Methods("PUT")
	router.HandleFunc("/loans/{id}", handlers.DeleteLoan).Methods("DELETE")
}
