package handlers_test

import (
	"HealthAPI/internal/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateDoctor() {
	req, err := http.NewRequest("POST", "/doctors", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	handlers.CreateDoctor(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("expected status %d, got %d", http.StatusCreated, w.Code)
	}
}
