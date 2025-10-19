package routes_test

import (
	"LibraryAPI/internal/handlers"
	"LibraryAPI/internal/routes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterRoutes() {
	router := mux.NewRouter()
	routes.Register(router)
	req := httptest.NewRequest("GET", "/books", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
