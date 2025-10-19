package handlers_test

import (
	"ShopAPI/internal/handlers"
	"ShopAPI/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreatePayment(t *testing.T) {
	payment := models.Payment{
		OrderID:     1,
		Amount:      9.99,
		PaymentDate: "2023-01-01",
	}
	body, _ := json.Marshal(payment)
	req := httptest.NewRequest("POST", "/payments", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreatePayment(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}
