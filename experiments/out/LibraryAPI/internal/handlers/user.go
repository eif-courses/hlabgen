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
	// Logic to save user to database
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GetUsers handles fetching all users.
func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Logic to fetch users from database
	var users []models.User
	json.NewEncoder(w).Encode(users)
}

// UpdateUser handles updating an existing user.
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Logic to update user in database
}

// DeleteUser handles deleting a user.
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Logic to delete user from database
}
