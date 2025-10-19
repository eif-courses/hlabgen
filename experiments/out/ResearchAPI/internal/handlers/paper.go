package handlers

import (
	"ResearchAPI/internal/models"
	"encoding/json"
	"net/http"
)

var papers []models.Paper

func CreatePaper(w http.ResponseWriter, r *http.Request) {
	var paper models.Paper
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&paper); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	papers = append(papers, paper)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(paper)
}
