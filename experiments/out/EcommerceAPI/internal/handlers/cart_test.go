package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"EcommerceAPI/internal/handlers"
	"EcommerceAPI/internal/models"
)

func TestCreateCart(t *testing.T) {
cart := models.Cart{
		UserID: 1,
		Items: []struct {
			ProductID int `json:"product_id"`,
			Quantity  int `json:"quantity"`
		}{
			{
				ProductID: 1,
				Quantity:  2,
			},
		},
}
	body, _ := json.Marshal(cart)
	req := httptest.NewRequest("POST", "/carts", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateCart(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}