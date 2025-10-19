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
		Title:    "Test Auction",
		Category: "Test Category",
		Price:    100.0,
	}
	body, _ := json.Marshal(auction)
	req := httptest.NewRequest("POST", "/auctions", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateAuction(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}
