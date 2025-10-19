package handlers

import (
	"SurveyAPI/internal/models"
	"encoding/json"
	"net/http"
)

var questions []models.Question

func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	// Implementation for creating a question
}
func GetQuestions(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting all questions
}
func GetQuestion(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a single question
}
func UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating a question
}
func DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting a question
}
