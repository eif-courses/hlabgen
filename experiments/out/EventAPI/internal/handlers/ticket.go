package handlers

import (
	"EventAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateTicket(w http.ResponseWriter, r *http.Request) {
	var ticket models.Ticket
	if err := json.NewDecoder(r.Body).Decode(&ticket); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ticket)
}

func GetTicket(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}

func UpdateTicket(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}

func DeleteTicket(w http.ResponseWriter, r *http.Request) {
	// Implementation here
}
