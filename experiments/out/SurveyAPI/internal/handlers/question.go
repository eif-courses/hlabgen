package handlers

import (
	"SurveyAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateQuestion() {
	var question models.Question
	if err := json.NewDecoder(r.Body).Decode(&question); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(question)
}

func GetQuestions() {
	// Implementation for getting questions
}
func UpdateQuestion() {
	// Implementation for updating a question
}
func DeleteQuestion() {
	// Implementation for deleting a question
}
