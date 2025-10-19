package handlers_test

import (
	"MovieAPI/internal/handlers"
	"MovieAPI/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateReview(t *testing.T) {
	review := models.Review{
		MovieID: 1,
		Rating:  5,
		Comment: "Great movie!",
	}
	body, _ := json.Marshal(review)
	req := httptest.NewRequest("POST", "/reviews", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateReview(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetReviews(t *testing.T) {
	req := httptest.NewRequest("GET", "/reviews", nil)
	w := httptest.NewRecorder()
	handlers.GetReviews(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
