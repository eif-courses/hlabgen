package tests

import (
	"LibraryAPI/internal/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateLoan(t *testing.T) {
	req, err := http.NewRequest("POST", "/loans", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handlers.CreateLoan(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}
