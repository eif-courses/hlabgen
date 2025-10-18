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
	// Logic to save vote to database
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(vote)
}

func GetVotes() {
	// Logic to retrieve votes from database
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode([]models.Vote{})
}
