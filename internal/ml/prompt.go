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

	// Build prompt with proper argument count
	promptText := `Generate Go REST API files as a JSON array of objects with "filename" and "code" fields.
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
- routes using mux.Router in internal/routes/routes.go
- tests in internal/handlers/ using package handlers_test
- tasks.md with 3 lab tasks

CRITICAL HANDLER STRUCTURE RULES:
All HTTP handlers MUST have this exact signature:
func HandlerName(w http.ResponseWriter, r *http.Request) {

Example handler:
func CreateBook(w http.ResponseWriter, r *http.Request) {
    var book models.Book
    if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(book)
}

CRITICAL TEST FUNCTION RULES - ABSOLUTELY MANDATORY:
- ALL test functions MUST have EXACTLY this signature: func TestName(t *testing.T) {
- The parameter MUST be: t *testing.T (pointer to testing.T with asterisk)
- Opening brace { MUST be on the same line as the function signature
- ALWAYS import "testing" package in all test files

CORRECT test examples:
✅ func TestCreateBook(t *testing.T) {
    book := models.Book{Title: "Test", Author: "Author"}
    body, _ := json.Marshal(book)
    req := httptest.NewRequest("POST", "/books", bytes.NewBuffer(body))
    w := httptest.NewRecorder()
    handlers.CreateBook(w, req)
    if w.Code != http.StatusCreated {
        t.Errorf("Expected 201, got %%d", w.Code)
    }
}

WRONG test examples (NEVER GENERATE):
❌ func TestCreateBook() {
❌ func TestCreateBook(t testing.T) {
❌ func TestCreateBook(t *testing.T, w http.ResponseWriter, r *http.Request) {

CRITICAL SYNTAX RULES:
1. Always add trailing commas in multi-line struct literals
2. Every field in multi-line struct MUST end with comma

CORRECT struct examples:
✅ item := models.Item{
    Name: "test",
    Price: 100,
}

✅ cart := models.Cart{
    UserID: 1,
    Items: []models.Item{
        {Name: "item1", Price: 10},
    },
    Total: 30,
}

WRONG struct examples:
❌ item := models.Item{
    Name: "test"
    Price: 100
}

ROUTES FUNCTION:
package routes
import (
    "github.com/gorilla/mux"
    "%s/internal/handlers"
)
func Register(r *mux.Router) {
    r.HandleFunc("/books", handlers.CreateBook).Methods("POST")
}

MAIN.GO:
package main
import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "%s/internal/routes"
)
func main() {
    r := mux.NewRouter()
    routes.Register(r)
    log.Fatal(http.ListenAndServe(":8080", r))
}

FINAL CRITICAL REMINDERS:
1. ✅ ALL test functions: func TestName(t *testing.T) {
2. ✅ ALL handler functions: func HandlerName(w http.ResponseWriter, r *http.Request) {
3. ✅ ALL struct fields: MUST end with comma
4. ✅ Import paths: Use "%s/internal/..." format

⚠️ ABSOLUTELY CRITICAL OUTPUT FORMAT:
Your ENTIRE response must be ONLY the JSON array.
Do NOT include any text before or after the JSON.
Do NOT use markdown code blocks.
START your response with [ and END with ]

WRONG:
Here is the code:
[...]

CORRECT:
[...]

Return ONLY valid JSON starting with [ now.`

	// Apply the format with correct number of arguments
	formattedPrompt := fmt.Sprintf(promptText,
		string(b), // %s - Requirements JSON
		s.AppName, // %s - Module name EXACTLY
		s.AppName, // %s - LOCAL module path
		s.AppName, // %s - import internal/models
		s.AppName, // %s - import internal/handlers
		s.AppName, // %s - github.com/yourusername/AppName
		s.AppName, // %s - routes import handlers
		s.AppName, // %s - main import routes
		s.AppName) // %s - final import format reminder

	buf.WriteString(formattedPrompt)
	return buf.String()
}
