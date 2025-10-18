package handlers_test

import (
	"BlogAPI/internal/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreatePost() {
	req, err := http.NewRequest("POST", "/posts", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.CreatePost)
	handler.ServeHTTP(rr, req)
	// Add assertions here
}

// Additional tests would go here
