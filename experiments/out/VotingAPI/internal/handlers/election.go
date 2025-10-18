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
	// Logic to save election to database
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(election)
}

func GetElections() {
	// Logic to retrieve elections from database
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode([]models.Election{})
}
