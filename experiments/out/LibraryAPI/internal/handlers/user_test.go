package handlers_test

import (
	"LibraryAPI/internal/handlers"
    "net/http"
    "net/http/httptest"
    "testing"
    "LibraryAPI/internal/handlers"
)

func TestCreateUser(t *testing.T) {
	body := strings.NewReader("{\"id\":1}")
	req, _ := http.NewRequest("POST", "/users", body)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handlers.CreateUser(rr, req)
	if rr.Code != http.StatusCreated && rr.Code != http.StatusOK && rr.Code != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v", rr.Code)
	}
}
    req, err := http.NewRequest("POST", "/users", nil)
    if err != nil {
        t.Fatal(err)
    }
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(handlers.CreateUser)
    handler.ServeHTTP(rr, req)
    if status := rr.Code; status != http.StatusCreated {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
    }
}
