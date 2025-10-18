package handlers

import (
	"VotingAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateElection() {
	var election models.Election
	if err := json.NewDecoder(r.Body).Decode(&election); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(election)
}

func GetElections() {
	// Implementation for getting elections
}
