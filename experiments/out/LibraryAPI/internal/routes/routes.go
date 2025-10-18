package routes

import (
	"LibraryAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register() {
	r.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	r.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")
	r.HandleFunc("/members", handlers.CreateMember).Methods("POST")
	r.HandleFunc("/members", handlers.GetMembers).Methods("GET")
	r.HandleFunc("/members/{id}", handlers.UpdateMember).Methods("PUT")
	r.HandleFunc("/members/{id}", handlers.DeleteMember).Methods("DELETE")
	r.HandleFunc("/loans", handlers.CreateLoan).Methods("POST")
	r.HandleFunc("/loans", handlers.GetLoans).Methods("GET")
	r.HandleFunc("/loans/{id}", handlers.UpdateLoan).Methods("PUT")
	r.HandleFunc("/loans/{id}", handlers.DeleteLoan).Methods("DELETE")
}
