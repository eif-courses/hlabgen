package handlers_test

import (
	"EcommerceAPI/internal/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateProduct() {
	req, err := http.NewRequest("POST", "/products", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.CreateProduct)
	handler.ServeHTTP(rr, req)
	// Check response code and other assertions
}
