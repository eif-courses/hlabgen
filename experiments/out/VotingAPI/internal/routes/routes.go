package routes

import (
	"VotingAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	// Election routes
	r.HandleFunc("/elections", handlers.CreateElection).Methods("POST")
	r.HandleFunc("/elections", handlers.GetElections).Methods("GET")
}
