package ml

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Schema struct {
	AppName    string
	Database   string
	APIPattern string
	Difficulty string
	Entities   []string
	Features   []string
	Objectives []string
}

func BuildPrompt(s Schema) string {
	var buf bytes.Buffer
	b, _ := json.Marshal(s)

	fmt.Fprintf(&buf, `Generate Go REST API files as a JSON array of objects with "filename" and "code" fields.
Requirements: %s

CRITICAL IMPORT RULES - READ CAREFULLY:
- Module name is EXACTLY: %s
- ALL imports must use the LOCAL module path: "%s/internal/..."
- DO NOT use "github.com/yourusername/..." or any GitHub paths
- DO NOT use "github.com/eif-courses/..." paths
- Use ONLY gorilla/mux for routing (import "github.com/gorilla/mux" only in routes.go)
- Use standard net/http handlers: func(w http.ResponseWriter, r *http.Request)
- DO NOT use gin.Context or any gin-specific types

CORRECT import examples for this project:
✅ import "%s/internal/models"
✅ import "%s/internal/handlers"
❌ import "github.com/yourusername/%s/internal/models"
❌ import "yourapp/internal/models"

Include:
- models with JSON tags in internal/models/
- handlers with standard http.ResponseWriter signature in internal/handlers/
- routes using mux.Router in internal/routes/routes.go (append to existing Register function)
- tests in internal/handlers/ using package handlers_test
- tasks.md with 3 lab tasks

Example handler structure:
func CreateBook(w http.ResponseWriter, r *http.Request) {
    var book models.Book
    if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(book)
}

CRITICAL TEST FUNCTION RULES:
- ALL test functions MUST have EXACTLY this signature: func TestName(t *testing.T)
- The parameter MUST be: t *testing.T (pointer to testing.T)
- Test functions MUST start with "Test" followed by capitalized name
- Import "testing" package in all test files

CORRECT test function examples:
✅ func TestCreateBook(t *testing.T) { ... }
✅ func TestGetBook(t *testing.T) { ... }
✅ func TestUpdateBook(t *testing.T) { ... }

WRONG test function examples (DO NOT USE):
❌ func TestCreateBook() { ... }                                    // Missing parameter
❌ func TestCreateBook(t testing.T) { ... }                         // Missing pointer *
❌ func TestCreateBook(t *testing.T, w http.ResponseWriter, r *http.Request) { ... }  // Extra parameters

Example test structure:
package handlers_test

import (
    "testing"
    "net/http"
    "net/http/httptest"
    "bytes"
    "encoding/json"
    "%s/internal/handlers"
    "%s/internal/models"
)

func TestCreateBook(t *testing.T) {
    book := models.Book{Title: "Test Book", Author: "Test Author"}
    body, _ := json.Marshal(book)
    req := httptest.NewRequest("POST", "/books", bytes.NewBuffer(body))
    w := httptest.NewRecorder()
    
    handlers.CreateBook(w, req)
    
    if w.Code != http.StatusCreated {
        t.Errorf("Expected status 201, got %%d", w.Code)
    }
}

CRITICAL SYNTAX RULES:
- Always add trailing commas in multi-line struct literals
- Close all braces and parentheses properly
- Use proper Go formatting

Example of CORRECT struct initialization:
✅ item := models.Item{
    Name: "test",
    Price: 100,     // Trailing comma required
}

Example of WRONG struct initialization:
❌ item := models.Item{
    Name: "test"
    Price: 100      // Missing comma - will cause compile error
}

CRITICAL: Return ONLY a valid JSON array. Do NOT wrap in markdown. Start with [ and end with ].`,
		string(b), s.AppName, s.AppName, s.AppName, s.AppName, s.AppName, s.AppName, s.AppName)
	return buf.String()
}
