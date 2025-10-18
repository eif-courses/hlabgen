package handlers

import (
	"ResearchAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreatePaper() {
	var paper models.Paper
	if err := json.NewDecoder(r.Body).Decode(&paper); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(paper)
}

func GetPaper() {
	// Implementation here
}
func UpdatePaper() {
	// Implementation here
}
func DeletePaper() {
	// Implementation here
}
