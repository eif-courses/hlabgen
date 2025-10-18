package handlers

import (
	"PortfolioAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateClient() {
	var client models.Client
	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(client)
}

func GetClients() {
	// Implementation for retrieving clients
}
func UpdateClient() {
	// Implementation for updating a client
}
func DeleteClient() {
	// Implementation for deleting a client
}
