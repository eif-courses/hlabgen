package handlers_test

import (
	"SurveyAPI/internal/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateSurvey() {
	req, err := http.NewRequest("POST", "/surveys", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.CreateSurvey)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}

func TestGetSurveys() {
	req, err := http.NewRequest("GET", "/surveys", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetSurveys)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
