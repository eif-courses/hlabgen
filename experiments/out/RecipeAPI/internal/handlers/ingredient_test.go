package handlers_test

import (
	"RecipeAPI/internal/handlers"
	"RecipeAPI/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateIngredient(t *testing.T) {
	ingredient := models.Ingredient{
		Name: "Test Ingredient",
	}
	body, _ := json.Marshal(ingredient)
	req := httptest.NewRequest("POST", "/ingredients", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handlers.CreateIngredient(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetIngredients(t *testing.T) {
	req := httptest.NewRequest("GET", "/ingredients", nil)
	w := httptest.NewRecorder()
	handlers.GetIngredients(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
