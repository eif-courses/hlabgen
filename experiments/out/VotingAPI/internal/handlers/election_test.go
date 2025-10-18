package handlers_test

import (
	"VotingAPI/internal/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateElection() {
	req, err := http.NewRequest("POST", "/elections", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.CreateElection)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}
