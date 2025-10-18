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

	// Write the prompt in sections for clarity
	buf.WriteString(fmt.Sprintf(`Generate Go REST API files as a JSON array of objects with "filename" and "code" fields.
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

CRITICAL HANDLER STRUCTURE RULES:
All HTTP handlers MUST have this exact signature:
func HandlerName(w http.ResponseWriter, r *http.Request) {

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

CORRECT handler examples:
✅ func CreateProduct(w http.ResponseWriter, r *http.Request) { ... }
✅ func GetProduct(w http.ResponseWriter, r *http.Request) { ... }
✅ func UpdateProduct(w http.ResponseWriter, r *http.Request) { ... }

WRONG handler examples (DO NOT USE):
❌ func CreateProduct() { ... }                                    // Missing parameters
❌ func CreateProduct(c *gin.Context) { ... }                      // Wrong framework
❌ func CreateProduct(w ResponseWriter, r Request) { ... }         // Missing http package

CRITICAL TEST FUNCTION RULES - ABSOLUTELY MANDATORY:
- ALL test functions MUST have EXACTLY this signature: func TestName(t *testing.T) {
- The parameter MUST be: t *testing.T (pointer to testing.T with asterisk)
- Opening brace { MUST be on the same line as the function signature
- Test functions MUST start with "Test" followed by capitalized name
- ALWAYS import "testing" package in all test files
- Use net/http/httptest for HTTP testing

CORRECT test function examples (COPY THESE EXACTLY):
✅ func TestCreateBook(t *testing.T) {
    // test code here
}

✅ func TestGetBook(t *testing.T) {
    // test code here
}

✅ func TestUpdateBook(t *testing.T) {
    req := httptest.NewRequest("POST", "/books", nil)
    w := httptest.NewRecorder()
    handlers.CreateBook(w, req)
    if w.Code != http.StatusOK {
        t.Errorf("got %%d, want %%d", w.Code, http.StatusOK)
    }
}

WRONG test function examples (NEVER GENERATE THESE):
❌ func TestCreateBook() {
    // Missing parameter completely
}

❌ func TestCreateBook(t testing.T) {
    // Missing pointer * before testing.T
}

❌ func TestCreateBook(t *testing.T, w http.ResponseWriter, r *http.Request) {
    // Extra parameters - test functions only take t *testing.T
}

❌ func TestCreateBook(w http.ResponseWriter, r *http.Request) {
    // Wrong parameters - this is a handler signature, not a test
}

❌ func TestCreateBook(t *testing.T)
{
    // Opening brace on wrong line
}

CRITICAL: Every test function MUST look EXACTLY like this:
func TestSomething(t *testing.T) {
    // implementation
}

Example COMPLETE test file structure:
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
    book := models.Book{
        Title: "Test Book",
        Author: "Test Author",
    }
    body, _ := json.Marshal(book)
    req := httptest.NewRequest("POST", "/books", bytes.NewBuffer(body))
    w := httptest.NewRecorder()
    
    handlers.CreateBook(w, req)
    
    if w.Code != http.StatusCreated {
        t.Errorf("Expected status 201, got %%d", w.Code)
    }
}

CRITICAL SYNTAX RULES - MUST FOLLOW:
1. Always add trailing commas in multi-line struct literals
2. Close all braces and parentheses properly
3. Use proper Go formatting
4. Every field in a multi-line struct initialization MUST end with a comma
5. The last field before closing brace MUST also have a comma

Example of CORRECT struct initialization:
✅ item := models.Item{
    Name: "test",
    Price: 100,
}

✅ cart := models.Cart{
    UserID: 1,
    Items: []models.Item{
        {Name: "item1", Price: 10},
        {Name: "item2", Price: 20},
    },
    Total: 30,
}

✅ user := models.User{
    ID: 1,
    Name: "John",
    Email: "john@example.com",
}

✅ Complex nested example:
product := models.Product{
    ID: 1,
    Name: "Laptop",
    Details: models.Details{
        Brand: "Dell",
        Model: "XPS 13",
        Year: 2024,
    },
    Tags: []string{
        "electronics",
        "computers",
    },
}

Example of WRONG struct initialization (DO NOT GENERATE THESE):
❌ item := models.Item{
    Name: "test"
    Price: 100
}

❌ cart := models.Cart{
    UserID: 1
    Items: []models.Item{}
}

❌ user := models.User{
    ID: 1,
    Name: "John"
    Email: "john@example.com"
}

❌ product := models.Product{
    ID: 1,
    Name: "Laptop",
    Details: models.Details{
        Brand: "Dell"
        Model: "XPS 13"
    }
}

CRITICAL: Every field in a multi-line struct MUST end with a comma, including the last field before the closing brace.

ROUTES FUNCTION SIGNATURE:
The Register function in routes.go MUST have this signature:
func Register(r *mux.Router) {

Example routes structure:
package routes

import (
    "github.com/gorilla/mux"
    "%s/internal/handlers"
)

func Register(r *mux.Router) {
    r.HandleFunc("/books", handlers.CreateBook).Methods("POST")
    r.HandleFunc("/books/{id:[0-9]+}", handlers.GetBook).Methods("GET")
    r.HandleFunc("/books/{id:[0-9]+}", handlers.UpdateBook).Methods("PUT")
    r.HandleFunc("/books/{id:[0-9]+}", handlers.DeleteBook).Methods("DELETE")
}

MAIN.GO STRUCTURE:
The main.go file should import and use the routes package:
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

MODELS STRUCTURE:
All models should use proper JSON tags and Go types:
package models

type Book struct {
    ID     int    `+"`json:\"id\"`"+`
    Title  string `+"`json:\"title\"`"+`
    Author string `+"`json:\"author\"`"+`
    ISBN   string `+"`json:\"isbn\"`"+`
    Year   int    `+"`json:\"year\"`"+`
}

TESTING BEST PRACTICES:
1. Use table-driven tests when appropriate
2. Test both success and error cases
3. Use httptest.NewRequest and httptest.NewRecorder
4. Verify status codes and response bodies
5. Keep tests independent and isolated

Example comprehensive test:
func TestCreateProduct(t *testing.T) {
    tests := []struct {
        name       string
        product    models.Product
        wantStatus int
    }{
        {
            name: "valid product",
            product: models.Product{
                Name:  "Laptop",
                Price: 999,
            },
            wantStatus: http.StatusCreated,
        },
        {
            name: "product with special chars",
            product: models.Product{
                Name:  "Product & Name",
                Price: 50,
            },
            wantStatus: http.StatusCreated,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            body, _ := json.Marshal(tt.product)
            req := httptest.NewRequest("POST", "/products", bytes.NewBuffer(body))
            w := httptest.NewRecorder()
            
            handlers.CreateProduct(w, req)
            
            if w.Code != tt.wantStatus {
                t.Errorf("Expected status %%d, got %%d", tt.wantStatus, w.Code)
            }
        })
    }
}

ERROR HANDLING:
Always include proper error handling in handlers:
✅ Good error handling:
func GetBook(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    
    // Validate input
    if id == "" {
        http.Error(w, "ID is required", http.StatusBadRequest)
        return
    }
    
    // Process request
    // ...
    
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(book)
}

RESPONSE FORMATS:
Always use proper HTTP status codes:
- 200 OK: Successful GET, PUT
- 201 Created: Successful POST
- 204 No Content: Successful DELETE
- 400 Bad Request: Invalid input
- 404 Not Found: Resource not found
- 500 Internal Server Error: Server error

FINAL CRITICAL REMINDERS BEFORE GENERATING CODE:
1. ✅ ALL test functions: func TestName(t *testing.T) { with opening brace on same line
2. ✅ ALL handler functions: func HandlerName(w http.ResponseWriter, r *http.Request) {
3. ✅ ALL struct fields in multi-line literals: MUST end with comma, including last field
4. ✅ Routes Register function: func Register(r *mux.Router) {
5. ✅ Import paths: Use "%s/internal/..." format only
6. ✅ NO markdown formatting, NO code blocks, ONLY valid JSON array

If you generate even ONE test function without (t *testing.T), the entire output is invalid.
If you generate even ONE struct without trailing commas, the code will not compile.

CRITICAL: Return ONLY a valid JSON array. Do NOT wrap in markdown. Start with [ and end with ].
The JSON must be parseable by encoding/json.
Each object must have "filename" (string) and "code" fields.

Example output format:
[
    {
        "filename": "internal/models/book.go",
        "code": "package models\n\ntype Book struct {\n\tID int `+"`json:\"id\"`"+`\n}"
    },
    {
        "filename": "internal/handlers/book.go",
        "code": "package handlers\n\nfunc CreateBook(w http.ResponseWriter, r *http.Request) {\n\t// implementation\n}"
    }
]`,
		string(b),  // 1
		s.AppName,  // 2
		s.AppName,  // 3
		s.AppName,  // 4
		s.AppName,  // 5
		s.AppName,  // 6
		s.AppName,  // 7
		s.AppName,  // 8
		s.AppName,  // 9
		s.AppName,  // 10
		s.AppName)) // 11

	return buf.String()
}
