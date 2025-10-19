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

func TestCreateAuction(t *testing.T) {
	auction := models.Auction{
		Title:       "Test Auction",
		Description: "Test Description",
		UserID:      1,
		StartPrice:  100.0,
		EndTime:     "2023-12-31T23:59:59Z",
	}
	body, _ := json.Marshal(auction)
	req := httptest.NewRequest("POST", "/auctions", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateAuction(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetAuctions(t *testing.T) {
	req := httptest.NewRequest("GET", "/auctions", nil)
	w := httptest.NewRecorder()
	handlers.GetAuctions(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
