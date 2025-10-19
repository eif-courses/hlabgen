package handlers_test

import (
	"CarRentalAPI/internal/handlers"
	"CarRentalAPI/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateCar(t *testing.T) {
	car := models.Car{
		Make:   "Toyota",
		Model:  "Camry",
		Year:   2020,
		Status: "available",
	}
	body, _ := json.Marshal(car)
	req := httptest.NewRequest("POST", "/cars", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	handlers.CreateCar(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}
