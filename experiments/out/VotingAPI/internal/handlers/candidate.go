package handlers

import (
	"VotingAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateCandidate() {
	var candidate models.Candidate
	if err := json.NewDecoder(r.Body).Decode(&candidate); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(candidate)
}

func GetCandidates() {
	// Implementation for getting candidates
}
