package routes

import (
	"LibraryAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	r.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	r.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")

	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	r.HandleFunc("/loans", handlers.CreateLoan).Methods("POST")
	r.HandleFunc("/loans", handlers.GetLoans).Methods("GET")
	r.HandleFunc("/loans/{id}", handlers.UpdateLoan).Methods("PUT")
	r.HandleFunc("/loans/{id}", handlers.DeleteLoan).Methods("DELETE")
}
