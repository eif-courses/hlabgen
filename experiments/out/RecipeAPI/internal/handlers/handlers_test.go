package handlers_test

import (
	"RecipeAPI/internal/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateRecipe() {
	req, err := http.NewRequest("POST", "/recipes", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.CreateRecipe)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}

// Additional tests for other handlers...
