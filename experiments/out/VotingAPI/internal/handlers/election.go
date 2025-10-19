package handlers

import (
	"VotingAPI/internal/models"
	"encoding/json"
	"net/http"
)

var elections []models.Election

func CreateElection(w http.ResponseWriter, r *http.Request) {
	var election models.Election
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&election); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	elections = append(elections, election)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(election)
}

func GetElections(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(elections)
}
