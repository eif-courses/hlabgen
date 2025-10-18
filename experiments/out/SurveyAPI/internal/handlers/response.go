package handlers

import (
	"SurveyAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateResponse() {
	var response models.Response
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func GetResponses() {
	// Implementation for getting responses
}
func UpdateResponse() {
	// Implementation for updating a response
}
func DeleteResponse() {
	// Implementation for deleting a response
}
