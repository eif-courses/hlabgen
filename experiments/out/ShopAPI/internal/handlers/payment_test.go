package handlers_test

import (
	"ShopAPI/internal/handlers"
	"ShopAPI/internal/routes"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreatePayment(t *testing.T) {
	body := strings.NewReader("{\"id\":1}")
	req, _ := http.NewRequest("POST", "/payments", body)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handlers.CreatePayment(rr, req)
	if rr.Code != http.StatusCreated && rr.Code != http.StatusOK && rr.Code != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v", rr.Code)
	}
}

func TestGetPayment(t *testing.T) {
	body := strings.NewReader("{}")
	req, _ := http.NewRequest("GET", "/payments", body)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handlers.GetPayment(rr, req)
	if rr.Code != http.StatusCreated && rr.Code != http.StatusOK && rr.Code != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v", rr.Code)
	}
}

// Additional tests for other handlers
