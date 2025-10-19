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

func TestCreateTransaction(t *testing.T) {
	transaction := models.Transaction{
		AccountID:   1,
		Amount:      50.0,
		Description: "Test Transaction",
	}
	body, _ := json.Marshal(transaction)
	req := httptest.NewRequest("POST", "/transactions", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateTransaction(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}
