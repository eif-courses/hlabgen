package routes

import (
	"ResearchAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register() {
	r.HandleFunc("/papers", handlers.CreatePaper).Methods("POST")
	r.HandleFunc("/papers/{id:[0-9]+}", handlers.GetPaper).Methods("GET")
	r.HandleFunc("/papers/{id:[0-9]+}", handlers.UpdatePaper).Methods("PUT")
	r.HandleFunc("/papers/{id:[0-9]+}", handlers.DeletePaper).Methods("DELETE")

	r.HandleFunc("/authors", handlers.CreateAuthor).Methods("POST")
	r.HandleFunc("/authors/{id:[0-9]+}", handlers.GetAuthor).Methods("GET")
	r.HandleFunc("/authors/{id:[0-9]+}", handlers.UpdateAuthor).Methods("PUT")
	r.HandleFunc("/authors/{id:[0-9]+}", handlers.DeleteAuthor).Methods("DELETE")

	r.HandleFunc("/reviews", handlers.CreateReview).Methods("POST")
	r.HandleFunc("/reviews/{id:[0-9]+}", handlers.GetReview).Methods("GET")
	r.HandleFunc("/reviews/{id:[0-9]+}", handlers.UpdateReview).Methods("PUT")
	r.HandleFunc("/reviews/{id:[0-9]+}", handlers.DeleteReview).Methods("DELETE")
}
