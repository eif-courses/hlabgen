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

func TestCreateRecipe(t *testing.T) {
	recipe := models.Recipe{
		Title:       "Test Recipe",
		Description: "Test Description",
	}
	body, _ := json.Marshal(recipe)
	req := httptest.NewRequest("POST", "/recipes", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handlers.CreateRecipe(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetRecipes(t *testing.T) {
	req := httptest.NewRequest("GET", "/recipes", nil)
	w := httptest.NewRecorder()
	handlers.GetRecipes(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
