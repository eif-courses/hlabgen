package routes

import (
	"MovieAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	// Movie routes
	r.HandleFunc("/movies", handlers.CreateMovie).Methods("POST")
	r.HandleFunc("/movies", handlers.GetMovies).Methods("GET")
	// Director routes
	r.HandleFunc("/directors", handlers.CreateDirector).Methods("POST")
	r.HandleFunc("/directors", handlers.GetDirectors).Methods("GET")
	// Review routes
	r.HandleFunc("/reviews", handlers.CreateReview).Methods("POST")
	r.HandleFunc("/reviews", handlers.GetReviews).Methods("GET")
}
