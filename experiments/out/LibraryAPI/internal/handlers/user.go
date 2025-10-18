package handlers

import (
	"LibraryAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Add logic to save user to database
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Add logic to retrieve users from database
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode([]models.User{})
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Add logic to update user in database
	w.WriteHeader(http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Add logic to delete user from database
	w.WriteHeader(http.StatusNoContent)
}
