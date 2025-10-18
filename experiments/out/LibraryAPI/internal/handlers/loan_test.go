package handlers_test

import (
	"LibraryAPI/internal/handlers"
    "net/http"
    "net/http/httptest"
    "testing"
    "LibraryAPI/internal/handlers"
)

func TestCreateLoan(t *testing.T) {
	body := strings.NewReader("{\"id\":1}")
	req, _ := http.NewRequest("POST", "/loans", body)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handlers.CreateLoan(rr, req)
	if rr.Code != http.StatusCreated && rr.Code != http.StatusOK && rr.Code != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v", rr.Code)
	}
}
    req, err := http.NewRequest("POST", "/loans", nil)
    if err != nil {
        t.Fatal(err)
    }
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(handlers.CreateLoan)
    handler.ServeHTTP(rr, req)
    if status := rr.Code; status != http.StatusCreated {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
    }
}
