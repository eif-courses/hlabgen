package handlers_test

import (
	"HealthAPI/internal/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateRecord() {
	req, err := http.NewRequest("POST", "/records", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	handlers.CreateRecord(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("expected status %d, got %d", http.StatusCreated, w.Code)
	}
}
