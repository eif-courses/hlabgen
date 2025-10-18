package handlers

import (
	"FitnessAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateUser() {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func GetUser() {
	// Implementation for getting a user
}
func UpdateUser() {
	// Implementation for updating a user
}
func DeleteUser() {
	// Implementation for deleting a user
}
