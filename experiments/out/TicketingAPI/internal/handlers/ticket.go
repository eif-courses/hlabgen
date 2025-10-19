package handlers

import (
	"TicketingAPI/internal/models"
	"encoding/json"
	"net/http"
)

var tickets []models.Ticket

func CreateTicket(w http.ResponseWriter, r *http.Request) {
	var ticket models.Ticket
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&ticket); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tickets = append(tickets, ticket)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ticket)
}
