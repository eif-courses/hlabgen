package handlers_test

import (
	"AuctionAPI/internal/handlers"
	"AuctionAPI/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateBid(t *testing.T) {
	bid := models.Bid{
		AuctionID: 1,
		UserID:    1,
		Amount:    50.0,
	}
	body, _ := json.Marshal(bid)
	req := httptest.NewRequest("POST", "/bids", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateBid(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}
