package handlers

import (
	"ProjectAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateProject() {
	var project models.Project
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(project)
}

func GetProjects() {
	// Implementation for retrieving projects
}
func UpdateProject() {
	// Implementation for updating a project
}
func DeleteProject() {
	// Implementation for deleting a project
}
