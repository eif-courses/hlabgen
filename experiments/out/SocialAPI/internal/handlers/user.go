package handlers

import (
	"SocialAPI/internal/models"
	"encoding/json"
	"net/http"
)

// CreateUser handles the creation of a new user.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GetUser handles fetching a user by ID.
func GetUser(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}

// UpdateUser handles updating user information.
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}

// DeleteUser handles deleting a user.
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}
