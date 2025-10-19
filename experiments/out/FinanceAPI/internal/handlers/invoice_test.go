package handlers_test

import (
	"FinanceAPI/internal/handlers"
	"FinanceAPI/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateInvoice(t *testing.T) {
	invoice := models.Invoice{
		Client: "Test Client",
		Amount: 75.0,
		Paid:   false,
	}
	body, _ := json.Marshal(invoice)
	req := httptest.NewRequest("POST", "/invoices", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateInvoice(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}
