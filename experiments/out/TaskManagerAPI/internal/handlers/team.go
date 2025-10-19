package handlers

import (
	"TaskManagerAPI/internal/models"
	"encoding/json"
	"net/http"
)

var teams []models.Team

func CreateTeam(w http.ResponseWriter, r *http.Request) {
	var team models.Team
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&team); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	teams = append(teams, team)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(team)
}
