package handlers

import (
	"SurveyAPI/internal/models"
	"encoding/json"
	"net/http"
)

var surveys []models.Survey

func CreateSurvey(w http.ResponseWriter, r *http.Request) {
	var survey models.Survey
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&survey); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	surveys = append(surveys, survey)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(survey)
}

func GetSurveys(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(surveys)
}
