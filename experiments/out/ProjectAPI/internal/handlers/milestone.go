package handlers

import (
	"ProjectAPI/internal/models"
	"encoding/json"
	"net/http"
)

var milestones []models.Milestone

func CreateMilestone(w http.ResponseWriter, r *http.Request) {
	var milestone models.Milestone
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&milestone); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	milestones = append(milestones, milestone)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(milestone)
}

func GetMilestones(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(milestones)
}
func GetMilestone(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a single milestone
	w.WriteHeader(http.StatusOK)
}
func UpdateMilestone(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating a milestone
	w.WriteHeader(http.StatusOK)
}
func DeleteMilestone(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting a milestone
	w.WriteHeader(http.StatusNoContent)
}
