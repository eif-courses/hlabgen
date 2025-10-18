package handlers

import (
	"EcommerceAPI/internal/models"
	"encoding/json"
	"net/http"
)

// CreateUser handles the creation of a new user.
func CreateUser() {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GetUsers handles fetching all users.
func GetUsers() {
	// Implementation for fetching users
}
