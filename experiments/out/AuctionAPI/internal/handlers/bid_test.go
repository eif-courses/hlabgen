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
		Amount:    150.0,
	}
	body, _ := json.Marshal(bid)
	req := httptest.NewRequest("POST", "/bids", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateBid(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetBids(t *testing.T) {
	req := httptest.NewRequest("GET", "/bids", nil)
	w := httptest.NewRecorder()
	handlers.GetBids(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
