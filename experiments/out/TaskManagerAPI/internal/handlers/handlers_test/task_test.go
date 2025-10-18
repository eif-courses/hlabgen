package handlers_test

import (
	"TaskManagerAPI/internal/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateTask() {
	req, err := http.NewRequest("POST", "/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.CreateTask)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}
