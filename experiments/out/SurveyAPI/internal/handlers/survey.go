package handlers

import (
	"SurveyAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateSurvey() {
	var survey models.Survey
	if err := json.NewDecoder(r.Body).Decode(&survey); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(survey)
}

func GetSurveys() {
	// Implementation for getting surveys
}
func UpdateSurvey() {
	// Implementation for updating a survey
}
func DeleteSurvey() {
	// Implementation for deleting a survey
}
