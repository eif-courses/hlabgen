package handlers

import (
	"ProjectAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateMilestone() {
	var milestone models.Milestone
	if err := json.NewDecoder(r.Body).Decode(&milestone); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(milestone)
}

func GetMilestones() {
	// Implementation for retrieving milestones
}
func UpdateMilestone() {
	// Implementation for updating a milestone
}
func DeleteMilestone() {
	// Implementation for deleting a milestone
}
