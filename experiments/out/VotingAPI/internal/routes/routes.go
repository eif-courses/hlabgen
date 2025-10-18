package routes

import (
	"VotingAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register() {
	r.HandleFunc("/elections", handlers.CreateElection).Methods("POST")
	r.HandleFunc("/elections", handlers.GetElections).Methods("GET")
	r.HandleFunc("/candidates", handlers.CreateCandidate).Methods("POST")
	r.HandleFunc("/candidates", handlers.GetCandidates).Methods("GET")
	r.HandleFunc("/votes", handlers.CreateVote).Methods("POST")
	r.HandleFunc("/votes", handlers.GetVotes).Methods("GET")
}
