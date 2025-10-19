package routes

import (
	"ResearchAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	// Paper routes
	r.HandleFunc("/papers", handlers.CreatePaper).Methods("POST")
	r.HandleFunc("/papers", handlers.GetPapers).Methods("GET")
	r.HandleFunc("/papers/{id}", handlers.GetPaper).Methods("GET")
	r.HandleFunc("/papers/{id}", handlers.UpdatePaper).Methods("PUT")
	r.HandleFunc("/papers/{id}", handlers.DeletePaper).Methods("DELETE")

	// Author routes
	r.HandleFunc("/authors", handlers.CreateAuthor).Methods("POST")
	r.HandleFunc("/authors", handlers.GetAuthors).Methods("GET")
	r.HandleFunc("/authors/{id}", handlers.GetAuthor).Methods("GET")
	r.HandleFunc("/authors/{id}", handlers.UpdateAuthor).Methods("PUT")
	r.HandleFunc("/authors/{id}", handlers.DeleteAuthor).Methods("DELETE")

	// Review routes
	r.HandleFunc("/reviews", handlers.CreateReview).Methods("POST")
	r.HandleFunc("/reviews", handlers.GetReviews).Methods("GET")
	r.HandleFunc("/reviews/{id}", handlers.GetReview).Methods("GET")
	r.HandleFunc("/reviews/{id}", handlers.UpdateReview).Methods("PUT")
	r.HandleFunc("/reviews/{id}", handlers.DeleteReview).Methods("DELETE")
}
