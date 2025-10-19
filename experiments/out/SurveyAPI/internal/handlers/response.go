package handlers

import (
	"SurveyAPI/internal/models"
	"encoding/json"
	"net/http"
)

var responses []models.Response

func CreateResponse(w http.ResponseWriter, r *http.Request) {
	// Implementation for creating a response
}
func GetResponses(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting all responses
}
func GetResponse(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a single response
}
func UpdateResponse(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating a response
}
func DeleteResponse(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting a response
}
