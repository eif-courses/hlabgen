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
	// Logic to save survey to database
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(survey)
}

func GetSurveys() {
	// Logic to retrieve surveys from database
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode([]models.Survey{})
}
func UpdateSurvey() {
	// Logic to update survey in database
}
func DeleteSurvey() {
	// Logic to delete survey from database
}
