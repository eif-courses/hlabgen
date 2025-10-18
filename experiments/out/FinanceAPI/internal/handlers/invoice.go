package handlers

import (
	"FinanceAPI/internal/models"
	"encoding/json"
	"net/http"
)

func CreateInvoice() {
	var invoice models.Invoice
	if err := json.NewDecoder(r.Body).Decode(&invoice); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(invoice)
}

func GetInvoices() {
	// Implementation for fetching invoices
}
func UpdateInvoice() {
	// Implementation for updating an invoice
}
func DeleteInvoice() {
	// Implementation for deleting an invoice
}
