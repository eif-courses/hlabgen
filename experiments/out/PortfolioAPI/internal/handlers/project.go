package handlers

import (
	"PortfolioAPI/internal/models"
	"encoding/json"
	"net/http"
)

var projects []models.Project

func CreateProject(w http.ResponseWriter, r *http.Request) {
	var project models.Project
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	projects = append(projects, project)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(project)
}

func GetProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}
func GetProject(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a single project
	w.WriteHeader(http.StatusOK)
}
func UpdateProject(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating a project
	w.WriteHeader(http.StatusOK)
}
func DeleteProject(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting a project
	w.WriteHeader(http.StatusNoContent)
}
