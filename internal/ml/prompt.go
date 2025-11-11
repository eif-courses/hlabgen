package ml

import (
	"encoding/json"
	"fmt"
	"strings"
	"text/template"
)

type Schema struct {
	AppName           string   `json:"appName"`
	Database          string   `json:"database"`
	APIPattern        string   `json:"apiPattern"`
	Difficulty        string   `json:"difficulty"`
	Entities          []string `json:"entities"`
	Features          []string `json:"features"`
	Objectives        []string `json:"objectives"`
	AllowMuxInHandler bool     `json:"allowMuxInHandler"` // set true to allow mux in handlers
}

func (s Schema) Validate() error {
	if strings.TrimSpace(s.AppName) == "" {
		return fmt.Errorf("AppName is required")
	}
	return nil
}

func BuildPrompt(s Schema) (string, error) {
	if err := s.Validate(); err != nil {
		return "", err
	}

	reqJSON, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", fmt.Errorf("marshal schema: %w", err)
	}

	importBase := fmt.Sprintf("%s/internal", s.AppName)
	bt := "`"   // for struct tags like `json:"id"`
	cf := "```" // for markdown code-fences injected at render time

	// Handler GetByID example + optional mux import bits
	handlerMuxImportRule := "• DO NOT import \"github.com/gorilla/mux\" in handlers"
	handlerMuxImportList := "• Handlers: \"encoding/json\", \"net/http\", \"strconv\""
	handlerMuxVarBlock := `
func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	for _, book := range books {
		if book.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}`
	handlerMuxImport := `	"github.com/gorilla/mux"`

	if !s.AllowMuxInHandler {
		handlerMuxImportRule = "• Prefer not to import \"github.com/gorilla/mux\" in handlers; extract the \"id\" path variable via the router, or set AllowMuxInHandler=true to permit mux.Vars."
		handlerMuxImport = ""
		handlerMuxVarBlock = `
func GetBook(w http.ResponseWriter, r *http.Request) {
	// NOTE: If using gorilla/mux, set AllowMuxInHandler=true to use mux.Vars.
	// Otherwise, ensure the router passes the "id" with request context or use a consistent path parser.
	// For safety in a template, we demonstrate a simple fallback that expects /books/{id}.
	path := strings.TrimPrefix(r.URL.Path, "/books/")
	id, err := strconv.Atoi(path)
	if err != nil || path == "" {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	for _, book := range books {
		if book.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}`
	}

	const promptTmpl = `
You are a Go code generator. Generate ONLY valid, compilable Go code for a REST API with COMPLETE, PRODUCTION-READY implementations.

Project Requirements: {{.Requirements}}

═══════════════════════════════════════════════════════════════
🚨 CRITICAL RULES - FOLLOW EXACTLY OR CODE WILL NOT COMPILE 🚨
═══════════════════════════════════════════════════════════════

1️⃣ MODULE AND IMPORT PATHS (ZERO TOLERANCE):
   • Module name is EXACTLY: {{.AppName}}
   • ALL imports MUST use: "{{.AppName}}/internal/..."
   • DO NOT use "github.com/yourusername/..."
   • DO NOT use "github.com/eif-courses/..."
   • DO NOT use "yourapp/..." or "your_project/..."
   
   ✅ CORRECT:
   import "{{.ImportBase}}/models"
   import "{{.ImportBase}}/handlers"
   import "{{.ImportBase}}/routes"
   
   ❌ WRONG:
   import "github.com/yourusername/{{.AppName}}/internal/models"
   import "yourapp/internal/models"
   import "your_project/internal/handlers"

2️⃣ HANDLER FUNCTION SIGNATURES (MANDATORY):
   ALL handler functions MUST have EXACTLY this signature:
   
   ✅ CORRECT:
   func CreateBook(w http.ResponseWriter, r *http.Request) {
       var book models.Book
       if r.Body == nil {
           http.Error(w, "missing body", http.StatusBadRequest)
           return
       }
       if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
           http.Error(w, err.Error(), http.StatusBadRequest)
           return
       }
       w.WriteHeader(http.StatusCreated)
       json.NewEncoder(w).Encode(book)
   }
   
   ❌ WRONG:
   func CreateBook() {
   func CreateBook(c *gin.Context) {
   func CreateBook(w http.ResponseWriter) {

3️⃣ ROUTES REGISTER FUNCTION (MANDATORY):
   The Register function MUST have this signature:
   
   ✅ CORRECT:
   package routes
   
   import (
       "github.com/gorilla/mux"
       "{{.ImportBase}}/handlers"
   )
   
   func Register(r *mux.Router) {
       r.HandleFunc("/books", handlers.CreateBook).Methods("POST")
       r.HandleFunc("/books", handlers.GetBooks).Methods("GET")
       r.HandleFunc("/books/{id}", handlers.GetBook).Methods("GET")
       r.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
       r.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")
   }
   
   ❌ WRONG:
   func Register() {
   func RegisterRoutes(router *mux.Router) {

4️⃣ TEST FUNCTION SIGNATURES (ABSOLUTELY MANDATORY - READ CAREFULLY):
   EVERY test function MUST have EXACTLY ONE parameter: t *testing.T
   
   ✅ CORRECT - ONLY ONE PARAMETER:
   func TestCreateBook(t *testing.T) {
   func TestGetBooks(t *testing.T) {
   func TestUpdateBook(t *testing.T) {
   
   ❌ WRONG - THESE WILL FAIL:
   func TestCreateBook() {                                    // ❌ NO parameters
   func TestCreateBook(t testing.T) {                         // ❌ Missing *
   func TestCreateBook(w http.ResponseWriter, r *http.Request) { // ❌ Wrong parameters
   func TestCreateBook(t *testing.T, w http.ResponseWriter) { // ❌ Extra parameters
   func TestCreateBook(ctx context.Context, t *testing.T) {   // ❌ Extra parameters
   
   🚨 CRITICAL: Tests are NOT handlers!
   • Handlers get: (w http.ResponseWriter, r *http.Request)
   • Tests get: (t *testing.T) ONLY
   
   DO NOT confuse them. Tests create mock requests like this:
   
   func TestCreateBook(t *testing.T) {  // ← ONLY ONE PARAMETER
       // Create mock request inside the test
       req := httptest.NewRequest("POST", "/books", body)
       w := httptest.NewRecorder()
       
       // Call the handler (which has w, r parameters)
       handlers.CreateBook(w, req)
       
       // Assert results
       if w.Code != http.StatusCreated {
           t.Errorf("Expected 201, got %d", w.Code)
       }
   }

5️⃣ STRUCT LITERAL SYNTAX (CRITICAL):
   Every field in multi-line struct literals MUST end with a comma:
   
   ✅ CORRECT:
   user := models.User{
       ID:       1,
       Username: "testuser",
       Email:    "test@example.com",
       Password: "password",
   }
   
   order := models.Order{
       UserID: 1,
       Products: []models.Product{
           {
               ID:    1,
               Name:  "Product",
               Price: 10.0,
               Stock: 100,
           },
       },
       Total: 10.0,
   }
   
   ❌ WRONG:
   user := models.User{
       ID: 1,
       Username: "testuser",
       Email: "test@example.com"
   }
   
   order := models.Order{UserID: 1, Products: []models.Product{{ID: 1, Name: "Product", Price: 10.0, Stock: 100}}, Total: 10
   }

6️⃣ IMPORT REQUIREMENTS:
   {{.HandlerImportList}}
   • Tests: "bytes", "encoding/json", "net/http", "net/http/httptest", "testing"
   • Routes: "github.com/gorilla/mux"
   {{.HandlerMuxRule}}
   • DO NOT import gin or any other frameworks

7️⃣ PACKAGE NAMES:
   • Handlers: package handlers
   • Models: package models
   • Routes: package routes
   • Tests: package handlers_test (NOT handlers)
   • Main: package main

8️⃣ FILE STRUCTURE:
   Generate files with these exact paths:
   • internal/models/<entity>.go
   • internal/handlers/<entity>.go
   • internal/handlers/<entity>_test.go
   • internal/routes/routes.go
   • cmd/main.go
   • tasks.md

9️⃣ COMPLETE IMPLEMENTATIONS REQUIRED:
   ⚠️  DO NOT generate placeholder comments or empty functions!
   ⚠️  ALL CRUD operations must be FULLY IMPLEMENTED!
   
   YOU MUST IMPLEMENT:
   • Create: Add item to in-memory slice, return 201 Created
   • GetAll: Return entire collection, 200 OK
   • GetByID: Extract ID from URL, search slice, return item or 404
   • Update: Extract ID, find item, update fields, return 200 OK
   • Delete: Extract ID, remove from slice, return 204 No Content
   
   ❌ WRONG (placeholder):
   func GetBook(w http.ResponseWriter, r *http.Request) {
       // Implementation for getting a single book
       w.WriteHeader(http.StatusOK)
   }
   
   ✅ CORRECT (complete):{{.HandlerGetByID}}

═══════════════════════════════════════════════════════════════
📋 COMPLETE FILE EXAMPLES (FULL IMPLEMENTATIONS)
═══════════════════════════════════════════════════════════════

EXAMPLE: internal/models/book.go
───────────────────────────────────
package models

type Book struct {
	ID     int    {{.BT}}json:"id"{{.BT}}
	Title  string {{.BT}}json:"title"{{.BT}}
	Author string {{.BT}}json:"author"{{.BT}}
	ISBN   string {{.BT}}json:"isbn"{{.BT}}
}

EXAMPLE: internal/handlers/book.go (COMPLETE IMPLEMENTATION)
───────────────────────────────────
package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"{{.HandlerMuxImport}}
	"{{.ImportBase}}/models"
{{- if not .AllowMuxInHandler}}
	"strings"
{{- end}}
)

var books []models.Book
var nextBookID = 1

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	book.ID = nextBookID
	nextBookID++
	books = append(books, book)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

{{.HandlerGetByID}}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
{{- if .AllowMuxInHandler}}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
{{- else}}
	path := strings.TrimPrefix(r.URL.Path, "/books/")
	id, err := strconv.Atoi(path)
{{- end}}
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var updatedBook models.Book
	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for i, book := range books {
		if book.ID == id {
			updatedBook.ID = id
			books[i] = updatedBook
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedBook)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
{{- if .AllowMuxInHandler}}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
{{- else}}
	path := strings.TrimPrefix(r.URL.Path, "/books/")
	id, err := strconv.Atoi(path)
{{- end}}
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

EXAMPLE: internal/handlers/book_test.go (COMPREHENSIVE TESTS)
───────────────────────────────────
package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gorilla/mux"
	"{{.ImportBase}}/handlers"
	"{{.ImportBase}}/models"
)

func TestCreateBook(t *testing.T) {
	book := models.Book{
		Title:  "Test Book",
		Author: "Test Author",
		ISBN:   "1234567890",
	}
	body, _ := json.Marshal(book)
	req := httptest.NewRequest("POST", "/books", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handlers.CreateBook(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}
}

func TestGetBooks(t *testing.T) {
	req := httptest.NewRequest("GET", "/books", nil)
	w := httptest.NewRecorder()
	handlers.GetBooks(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}

func TestGetBook(t *testing.T) {
	req := httptest.NewRequest("GET", "/books/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	w := httptest.NewRecorder()
	handlers.GetBook(w, req)
	if w.Code != http.StatusOK && w.Code != http.StatusNotFound {
		t.Errorf("Expected 200 or 404, got %d", w.Code)
	}
}

func TestUpdateBook(t *testing.T) {
	book := models.Book{
		Title:  "Updated Book",
		Author: "Updated Author",
		ISBN:   "0987654321",
	}
	body, _ := json.Marshal(book)
	req := httptest.NewRequest("PUT", "/books/1", bytes.NewBuffer(body))
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handlers.UpdateBook(w, req)
	if w.Code != http.StatusOK && w.Code != http.StatusNotFound {
		t.Errorf("Expected 200 or 404, got %d", w.Code)
	}
}

func TestDeleteBook(t *testing.T) {
	req := httptest.NewRequest("DELETE", "/books/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	w := httptest.NewRecorder()
	handlers.DeleteBook(w, req)
	if w.Code != http.StatusNoContent && w.Code != http.StatusNotFound {
		t.Errorf("Expected 204 or 404, got %d", w.Code)
	}
}

EXAMPLE: internal/routes/routes.go
───────────────────────────────────
package routes

import (
	"github.com/gorilla/mux"
	"{{.ImportBase}}/handlers"
)

func Register(r *mux.Router) {
	// Book routes
	r.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	r.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", handlers.GetBook).Methods("GET")
	r.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")
}

EXAMPLE: cmd/main.go
───────────────────────────────────
package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"{{.ImportBase}}/routes"
)

func main() {
	r := mux.NewRouter()
	routes.Register(r)
	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

EXAMPLE: tasks.md
───────────────────────────────────
# Lab Tasks - Advanced Features

## Completed Implementation
✅ All CRUD operations are fully implemented
✅ Complete test suite for all handlers
✅ Proper error handling and status codes

## Optional Enhancements (Student Tasks)
1. Add database persistence using SQLite or PostgreSQL
2. Implement authentication and authorization middleware
3. Add request validation using a validation library
4. Implement pagination for GetBooks endpoint
5. Add filtering and sorting capabilities
6. Create OpenAPI/Swagger documentation
7. Implement rate limiting middleware
8. Add logging middleware for all requests

═══════════════════════════════════════════════════════════════
🎯 OUTPUT FORMAT REQUIREMENTS
═══════════════════════════════════════════════════════════════

Your response MUST be ONLY a JSON array. No explanations, no markdown, no text.

✅ CORRECT OUTPUT:
[
  {
    "filename": "internal/models/book.go",
    "code": "package models\n\ntype Book struct {\n..."
  },
  {
    "filename": "internal/handlers/book.go",
    "code": "package handlers\n\nimport (\n..."
  }
]

❌ WRONG OUTPUT:
Here is the generated code:
{{.CF}}json
[...]
{{.CF}}

The response must START with [ and END with ]

═══════════════════════════════════════════════════════════════
⚡ FINAL CHECKLIST - VERIFY BEFORE RESPONDING
═══════════════════════════════════════════════════════════════

Before generating output, verify:
☐ ALL handlers have COMPLETE implementations (no placeholders!)
☐ ALL CRUD operations are FULLY functional (Create, GetAll, GetByID, Update, Delete)
☐ All handlers have (w http.ResponseWriter, r *http.Request)
☐ Register function has (r *mux.Router) parameter
☐ All test functions have (t *testing.T) parameter
☐ All multi-line struct literals have trailing commas
☐ All imports use "{{.AppName}}/internal/..." format
☐ Test package is "handlers_test" not "handlers"
☐ No gin imports or gin.Context usage
☐ Output is pure JSON array starting with [
☐ Generated complete test files with ALL 5 test functions per entity
☐ All handlers check r.Body == nil before decoding
☐ GetByID, Update, Delete handlers extract and validate ID from URL

Generate the complete REST API code now with FULL IMPLEMENTATIONS. Return ONLY the JSON array.
`

	data := map[string]any{
		"AppName":           s.AppName,
		"ImportBase":        importBase,
		"Requirements":      string(reqJSON),
		"BT":                bt,
		"CF":                cf,
		"AllowMuxInHandler": s.AllowMuxInHandler,
		"HandlerMuxImport":  handlerMuxImport,
		"HandlerMuxRule":    handlerMuxImportRule,
		"HandlerImportList": handlerMuxImportList,
		"HandlerGetByID":    handlerMuxVarBlock,
	}

	tmpl, err := template.New("prompt").Parse(promptTmpl)
	if err != nil {
		return "", fmt.Errorf("parse template: %w", err)
	}

	var b strings.Builder
	if err := tmpl.Execute(&b, data); err != nil {
		return "", fmt.Errorf("execute template: %w", err)
	}
	return b.String(), nil
}
