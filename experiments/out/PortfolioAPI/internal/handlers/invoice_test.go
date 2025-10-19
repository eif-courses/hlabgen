package handlers_test

import (
	"PortfolioAPI/internal/handlers"
	"PortfolioAPI/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateInvoice(t *testing.T) {
	invoice := models.Invoice{
		Amount: 100.50,
		Paid:   false,
	}
	body, _ := json.Marshal(invoice)
	req := httptest.NewRequest("POST", "/invoices", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handlers.CreateInvoice(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetInvoices(t *testing.T) {
	req := httptest.NewRequest("GET", "/invoices", nil)
	w := httptest.NewRecorder()
	handlers.GetInvoices(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
