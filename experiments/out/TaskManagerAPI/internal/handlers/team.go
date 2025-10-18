package handlers

import (
	"TaskManagerAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateTeam() {
	var team models.Team
	if err := json.NewDecoder(r.Body).Decode(&team); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Logic to save team to database
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(team)
}

func GetTeams() {
	// Logic to retrieve teams from database
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode([]models.Team{})
}
