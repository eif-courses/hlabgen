package handlers

import (
	"LibraryAPI/internal/models"
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
	// Save user to database (omitted)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GetUsers handles fetching all users.
func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Fetch users from database (omitted)
	var users []models.User
	json.NewEncoder(w).Encode(users)
}

// UpdateUser handles updating a user.
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Update user logic (omitted)
}

// DeleteUser handles deleting a user.
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Delete user logic (omitted)
}
