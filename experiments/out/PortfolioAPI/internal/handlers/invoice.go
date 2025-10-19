package handlers

import (
	"PortfolioAPI/internal/models"
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

func GetInvoices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoices)
}
func GetInvoice(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a single invoice
	w.WriteHeader(http.StatusOK)
}
func UpdateInvoice(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating an invoice
	w.WriteHeader(http.StatusOK)
}
func DeleteInvoice(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting an invoice
	w.WriteHeader(http.StatusNoContent)
}
