package handlers

import (
	"AuctionAPI/internal/models"
	"encoding/json"
	"net/http"
)

var users []models.User

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
	users = append(users, user)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a single user
	w.WriteHeader(http.StatusOK)
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating a user
	w.WriteHeader(http.StatusOK)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting a user
	w.WriteHeader(http.StatusNoContent)
}
