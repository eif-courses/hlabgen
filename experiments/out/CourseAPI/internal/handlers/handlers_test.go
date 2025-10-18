package handlers_test

import (
	"CourseAPI/internal/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateCourse() {
	req, err := http.NewRequest("POST", "/courses", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	handlers.CreateCourse(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %v", w.Code)
	}
}

// Additional tests for other handlers would go here...
