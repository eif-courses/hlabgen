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

func TestCreateAccount(t *testing.T) {
	account := models.Account{
		Name:    "Test Account",
		Balance: 100.0,
	}
	body, _ := json.Marshal(account)
	req := httptest.NewRequest("POST", "/accounts", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateAccount(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}
