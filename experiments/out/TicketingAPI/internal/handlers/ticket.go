package handlers

import (
	"TicketingAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateTicket() {
	var ticket models.Ticket
	if err := json.NewDecoder(r.Body).Decode(&ticket); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ticket)
}
