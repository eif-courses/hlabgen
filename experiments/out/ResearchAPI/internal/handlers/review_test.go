package handlers_test

import (
	"ResearchAPI/internal/handlers"
	"ResearchAPI/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateReview(t *testing.T) {
	review := models.Review{
		PaperID:  1,
		Reviewer: "Test Reviewer",
		Rating:   5,
		Comments: "Great paper",
	}
	body, _ := json.Marshal(review)
	req := httptest.NewRequest("POST", "/reviews", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateReview(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}
