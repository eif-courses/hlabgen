package handlers_test

import (
	"VotingAPI/internal/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateVote() {
	req, err := http.NewRequest("POST", "/votes", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.CreateVote)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}
