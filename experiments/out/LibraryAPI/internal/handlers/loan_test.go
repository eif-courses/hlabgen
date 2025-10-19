package handlers_test

import (
	"LibraryAPI/internal/handlers"
	"LibraryAPI/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateLoan(t *testing.T) {
	loan := models.Loan{
		BookID:   1,
		MemberID: 1,
		DueDate:  "2022-12-31",
	}
	body, _ := json.Marshal(loan)
	req := httptest.NewRequest("POST", "/loans", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handlers.CreateLoan(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetLoans(t *testing.T) {
	req := httptest.NewRequest("GET", "/loans", nil)
	w := httptest.NewRecorder()
	handlers.GetLoans(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
