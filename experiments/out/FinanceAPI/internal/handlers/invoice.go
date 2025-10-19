package handlers

import (
	"FinanceAPI/internal/models"
	"encoding/json"
	"net/http"
)

var invoices []models.Invoice

func CreateInvoice(w http.ResponseWriter, r *http.Request) {
	var invoice models.Invoice
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&invoice); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	invoices = append(invoices, invoice)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(invoice)
}
