package handlers

import (
	"PortfolioAPI/internal/models"
	"encoding/json"
	"net/http"
)

var clients []models.Client

func CreateClient(w http.ResponseWriter, r *http.Request) {
	var client models.Client
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	clients = append(clients, client)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(client)
}

func GetClients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clients)
}
func GetClient(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a single client
	w.WriteHeader(http.StatusOK)
}
func UpdateClient(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating a client
	w.WriteHeader(http.StatusOK)
}
func DeleteClient(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting a client
	w.WriteHeader(http.StatusNoContent)
}
