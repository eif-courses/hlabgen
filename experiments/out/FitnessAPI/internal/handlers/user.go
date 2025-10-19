package handlers

import (
	"FitnessAPI/internal/models"
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
