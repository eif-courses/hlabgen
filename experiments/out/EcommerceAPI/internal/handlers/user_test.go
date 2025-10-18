package handlers_test

import (
    "bytes"
    "encoding/json"
    "net/http/httptest"
    "testing"
    "EcommerceAPI/internal/handlers"
    "EcommerceAPI/internal/models"
)

func TestCreateUser() {
user := models.User{Username: "testuser", Email: "test@example.com", Password: "password"
}
    body, _ := json.Marshal(user)
    req := httptest.NewRequest("POST", "/users", bytes.NewBuffer(body))
    w := httptest.NewRecorder()
    handlers.CreateUser(w, req)
    if w.Code != http.StatusCreated {
        t.Errorf("Expected 201, got %d", w.Code)
    }
}
