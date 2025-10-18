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
	promptText := `You are a Go code generator. Generate ONLY valid, compilable Go code for a REST API.

Project Requirements: %s

═══════════════════════════════════════════════════════════════
🚨 CRITICAL RULES - FOLLOW EXACTLY OR CODE WILL NOT COMPILE 🚨
═══════════════════════════════════════════════════════════════

1️⃣ MODULE AND IMPORT PATHS (ZERO TOLERANCE):
   • Module name is EXACTLY: %s
   • ALL imports MUST use: "%s/internal/..."
   • DO NOT use "github.com/yourusername/..."
   • DO NOT use "github.com/eif-courses/..."
   • DO NOT use "yourapp/..." or "your_project/..."
   
   ✅ CORRECT:
   import "%s/internal/models"
   import "%s/internal/handlers"
   import "%s/internal/routes"
   
   ❌ WRONG:
   import "github.com/yourusername/%s/internal/models"
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
       "%s/internal/handlers"
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

4️⃣ TEST FUNCTION SIGNATURES (ABSOLUTELY MANDATORY):
   ALL test functions MUST have EXACTLY this signature:
   
   ✅ CORRECT:
   package handlers_test
   
   import (
       "bytes"
       "encoding/json"
       "net/http"
       "net/http/httptest"
       "testing"
       "%s/internal/handlers"
       "%s/internal/models"
   )
   
   func TestCreateBook(t *testing.T) {
       book := models.Book{
           Title:  "Test Book",
           Author: "Test Author",
       }
       body, _ := json.Marshal(book)
       req := httptest.NewRequest("POST", "/books", bytes.NewBuffer(body))
       w := httptest.NewRecorder()
       handlers.CreateBook(w, req)
       if w.Code != http.StatusCreated {
           t.Errorf("Expected 201, got %%d", w.Code)
       }
   }
   
   ❌ WRONG:
   func TestCreateBook() {
   func TestCreateBook(t testing.T) {
   func TestCreateBook(t *testing.T, w http.ResponseWriter, r *http.Request) {

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
   • Handlers: "encoding/json", "net/http"
   • Tests: "bytes", "encoding/json", "net/http", "net/http/httptest", "testing"
   • Routes: "github.com/gorilla/mux"
   • DO NOT import "github.com/gorilla/mux" in handlers
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

═══════════════════════════════════════════════════════════════
📋 COMPLETE FILE EXAMPLES
═══════════════════════════════════════════════════════════════

EXAMPLE: internal/models/book.go
───────────────────────────────────
package models

type Book struct {
	ID     int    ` + "`json:\"id\"`" + `
	Title  string ` + "`json:\"title\"`" + `
	Author string ` + "`json:\"author\"`" + `
	ISBN   string ` + "`json:\"isbn\"`" + `
}

EXAMPLE: internal/handlers/book.go
───────────────────────────────────
package handlers

import (
	"encoding/json"
	"net/http"
	"%s/internal/models"
)

var books []models.Book

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
	books = append(books, book)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	// Implementation for getting a single book
	w.WriteHeader(http.StatusOK)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating a book
	w.WriteHeader(http.StatusOK)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting a book
	w.WriteHeader(http.StatusNoContent)
}

EXAMPLE: internal/handlers/book_test.go
───────────────────────────────────
package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"%s/internal/handlers"
	"%s/internal/models"
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
		t.Errorf("Expected 201, got %%d", w.Code)
	}
}

func TestGetBooks(t *testing.T) {
	req := httptest.NewRequest("GET", "/books", nil)
	w := httptest.NewRecorder()
	handlers.GetBooks(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %%d", w.Code)
	}
}

EXAMPLE: internal/routes/routes.go
───────────────────────────────────
package routes

import (
	"github.com/gorilla/mux"
	"%s/internal/handlers"
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
	"%s/internal/routes"
)

func main() {
	r := mux.NewRouter()
	routes.Register(r)
	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

EXAMPLE: tasks.md
───────────────────────────────────
# Lab Tasks

1. Implement the GetBook handler to return a single book by ID
2. Add validation for required fields in CreateBook handler
3. Write additional test cases for UpdateBook and DeleteBook handlers
4. Implement error handling for book not found scenarios
5. Add pagination support for the GetBooks endpoint

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
` + "```json" + `
[...]
` + "```" + `

The response must START with [ and END with ]

═══════════════════════════════════════════════════════════════
⚡ FINAL CHECKLIST - VERIFY BEFORE RESPONDING
═══════════════════════════════════════════════════════════════

Before generating output, verify:
☐ All handlers have (w http.ResponseWriter, r *http.Request)
☐ Register function has (r *mux.Router) parameter
☐ All test functions have (t *testing.T) parameter
☐ All multi-line struct literals have trailing commas
☐ All imports use "%s/internal/..." format
☐ Test package is "handlers_test" not "handlers"
☐ No gin imports or gin.Context usage
☐ Output is pure JSON array starting with [
☐ Generated at least one test file for each handler file
☐ All handlers check r.Body == nil before decoding

Generate the complete REST API code now. Return ONLY the JSON array.`

	// Apply the format with correct number of arguments
	formattedPrompt := fmt.Sprintf(promptText,
		string(b), // %s - Requirements JSON
		s.AppName, // %s - Module name EXACTLY
		s.AppName, // %s - ALL imports path
		s.AppName, // %s - import models
		s.AppName, // %s - import handlers
		s.AppName, // %s - import routes
		s.AppName, // %s - wrong github import example
		s.AppName, // %s - routes Register import handlers
		s.AppName, // %s - test import handlers
		s.AppName, // %s - test import models
		s.AppName, // %s - handlers book example
		s.AppName, // %s - test book example (handlers import)
		s.AppName, // %s - test book example (models import)
		s.AppName, // %s - routes example
		s.AppName, // %s - main example
		s.AppName) // %s - final checklist

	buf.WriteString(formattedPrompt)
	return buf.String()
}
