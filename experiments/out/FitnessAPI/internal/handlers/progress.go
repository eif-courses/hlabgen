package handlers

import (
	"FitnessAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateProgress(w http.ResponseWriter, r *http.Request) {
	var progress models.Progress
	if err := json.NewDecoder(r.Body).Decode(&progress); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(progress)
}
