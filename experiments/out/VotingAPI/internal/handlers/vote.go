package handlers

import (
	"VotingAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateVote() {
	var vote models.Vote
	if err := json.NewDecoder(r.Body).Decode(&vote); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(vote)
}

func GetVotes() {
	// Implementation for getting votes
}
