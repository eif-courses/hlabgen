// internal/rules/templates.go
package rules

import (
	"fmt"
	"strings"
)

// --- helpers ---

func pluralize(s string) string {
	l := strings.ToLower(s)
	// very light pluralization helpers
	switch {
	case strings.HasSuffix(l, "s") || strings.HasSuffix(l, "x") || strings.HasSuffix(l, "z") ||
		strings.HasSuffix(l, "ch") || strings.HasSuffix(l, "sh"):
		return l + "es"
	case strings.HasSuffix(l, "y") && len(l) > 1 && !isVowel(l[len(l)-2]):
		return l[:len(l)-1] + "ies"
	default:
		return l + "s"
	}
}

func isVowel(b byte) bool {
	switch b {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}

// --- generators ---

// GenerateModel creates a basic model struct for an entity
func GenerateModel(entityName string) string {
	return fmt.Sprintf(`package models

import "fmt"

// %s represents a %s entity
type %s struct {
	ID          int    `+"`json:\"id\"`"+`
	Name        string `+"`json:\"name\"`"+`
	Description string `+"`json:\"description,omitempty\"`"+`
	CreatedAt   string `+"`json:\"created_at,omitempty\"`"+`
	UpdatedAt   string `+"`json:\"updated_at,omitempty\"`"+`
}

// generate%sID generates a unique ID for %s
func generate%sID() int {
	return fmt.Sprintf("%%d", %sCounter)
}

var %sCounter = 1
`, entityName, strings.ToLower(entityName), entityName,
		entityName, entityName, entityName, entityName, entityName)
}

// GenerateHandler creates a basic CRUD handler for an entity
func GenerateHandler(entityName string, moduleName string) string {
	entityPlural := pluralize(entityName) // lower-case plural for variable & route
	entityVar := entityPlural             // e.g. "loans"
	entityType := entityName              // e.g. "Loan"

	return fmt.Sprintf(`package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"%s/internal/models"
)

// In-memory storage for %s
var %s []models.%s
var next%sID = 1

// Create%s handles POST /%s
func Create%s(w http.ResponseWriter, r *http.Request) {
	var item models.%s

	if r.Body == nil {
		http.Error(w, "missing body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	item.ID = next%sID
	next%sID++
	%s = append(%s, item)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(item)
}

// Get%ss handles GET /%s
func Get%ss(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(%s)
}

// Get%s handles GET /%s/{id}
func Get%s(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	for _, item := range %s {
		if item.ID == id {
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(item)
			return
		}
	}

	http.Error(w, "%s not found", http.StatusNotFound)
}

// Update%s handles PUT /%s/{id}
func Update%s(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var updated models.%s
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, item := range %s {
		if item.ID == id {
			updated.ID = id
			%s[i] = updated
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(updated)
			return
		}
	}

	http.Error(w, "%s not found", http.StatusNotFound)
}

// Delete%s handles DELETE /%s/{id}
func Delete%s(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	for i, item := range %s {
		if item.ID == id {
			%s = append(%s[:i], %s[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "%s not found", http.StatusNotFound)
}
`,
		// --- arguments for fmt.Sprintf placeholders ---
		moduleName,            // import path
		entityType,            // comment: "In-memory storage for X"
		entityVar, entityType, // var <plural> []models.<Entity>
		entityType,               // next<Entity>ID
		entityType, entityPlural, // comment and route
		entityType,             // func Create<Entity>
		entityType,             // var item models.<Entity>
		entityType, entityType, // ID handling
		entityVar, entityVar, // append
		entityType, entityPlural, // Get<Entity>s comment
		entityType,               // func Get<Entity>s
		entityVar,                // encode plural
		entityType, entityPlural, // Get<Entity> comment
		entityType,               // func Get<Entity>
		entityVar,                // loop over plural
		entityType,               // "Entity not found"
		entityType, entityPlural, // Update<Entity> comment
		entityType,           // func Update<Entity>
		entityType,           // models.<Entity>
		entityVar, entityVar, // loop and update slice
		entityType,               // "Entity not found"
		entityType, entityPlural, // Delete<Entity> comment
		entityType,                                 // func Delete<Entity>
		entityVar, entityVar, entityVar, entityVar, // slice manipulation
		entityType, // "Entity not found"
	)
}

// GenerateTest creates a basic test file for an entity
func GenerateTest(entityName string, moduleName string) string {
	entityPlural := pluralize(entityName)

	return fmt.Sprintf(`package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"%s/internal/handlers"
	"%s/internal/models"
)

func TestCreate%s(t *testing.T) {
	item := models.%s{
		Name:        "Test %s",
		Description: "Test Description",
	}

	body, err := json.Marshal(item)
	if err != nil {
		t.Fatalf("Failed to marshal: %%v", err)
	}

	req := httptest.NewRequest("POST", "/%s", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handlers.Create%s(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %%d", w.Code)
	}

	var response models.%s
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %%v", err)
	}

	if response.ID == 0 {
		t.Error("Expected ID to be set")
	}

	if response.Name != item.Name {
		t.Errorf("Expected name %%s, got %%s", item.Name, response.Name)
	}
}

func TestGet%ss(t *testing.T) {
	// Always OK to return an array (possibly empty)
	req := httptest.NewRequest("GET", "/%s", nil)
	w := httptest.NewRecorder()

	handlers.Get%ss(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %%d", w.Code)
	}

	var items []models.%s
	if err := json.NewDecoder(w.Body).Decode(&items); err != nil {
		t.Fatalf("Failed to decode response: %%v", err)
	}
}

func TestGet%s(t *testing.T) {
	// Create a test item first
	item := models.%s{
		Name: "Test Item",
	}
	body, _ := json.Marshal(item)
	createReq := httptest.NewRequest("POST", "/%s", bytes.NewBuffer(body))
	createReq.Header.Set("Content-Type", "application/json")
	createW := httptest.NewRecorder()
	handlers.Create%s(createW, createReq)

	var created models.%s
	_ = json.NewDecoder(createW.Body).Decode(&created)

	// Now test Get
	req := httptest.NewRequest("GET", "/%s/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	w := httptest.NewRecorder()

	handlers.Get%s(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %%d", w.Code)
	}
}

func TestUpdate%s(t *testing.T) {
	// Create a test item first
	item := models.%s{
		Name: "Original Name",
	}
	body, _ := json.Marshal(item)
	createReq := httptest.NewRequest("POST", "/%s", bytes.NewBuffer(body))
	createReq.Header.Set("Content-Type", "application/json")
	createW := httptest.NewRecorder()
	handlers.Create%s(createW, createReq)

	// Update the item
	updated := models.%s{
		Name: "Updated Name",
	}
	updateBody, _ := json.Marshal(updated)
	req := httptest.NewRequest("PUT", "/%s/1", bytes.NewBuffer(updateBody))
	req.Header.Set("Content-Type", "application/json")
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	w := httptest.NewRecorder()

	handlers.Update%s(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %%d", w.Code)
	}
}

func TestDelete%s(t *testing.T) {
	// Create a test item first
	item := models.%s{
		Name: "To Delete",
	}
	body, _ := json.Marshal(item)
	createReq := httptest.NewRequest("POST", "/%s", bytes.NewBuffer(body))
	createReq.Header.Set("Content-Type", "application/json")
	createW := httptest.NewRecorder()
	handlers.Create%s(createW, createReq)

	// Delete the item
	req := httptest.NewRequest("DELETE", "/%s/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	w := httptest.NewRecorder()

	handlers.Delete%s(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("Expected status 204, got %%d", w.Code)
	}
}
`,
		// imports
		moduleName, moduleName,
		// TestCreate<Entity>
		entityName,
		entityName, entityName,
		entityPlural,
		entityName,
		entityName,
		// TestGet<Entities>
		entityName,
		entityPlural,
		entityName,
		entityName,
		// TestGet<Entity>
		entityName,
		entityName,
		entityPlural,
		entityName,
		entityName,
		entityPlural,
		entityName,
		// TestUpdate<Entity>
		entityName,
		entityName,
		entityPlural,
		entityName,
		entityName,
		entityPlural,
		entityName,
		// TestDelete<Entity>
		entityName,
		entityName,
		entityPlural,
		entityName,
		entityPlural,
		entityName,
	)
}

// GenerateRoutes creates the routes registration with all entities
func GenerateRoutes(entities []string, moduleName string) string {
	var routesList strings.Builder

	routesList.WriteString(`package routes

import (
	"github.com/gorilla/mux"
	"` + moduleName + `/internal/handlers"
)

// Register sets up all API routes
func Register(r *mux.Router) {
`)

	for _, entity := range entities {
		entityPlural := pluralize(entity)

		routesList.WriteString(fmt.Sprintf(`
	// %s routes
	r.HandleFunc("/%s", handlers.Create%s).Methods("POST")
	r.HandleFunc("/%s", handlers.Get%ss).Methods("GET")
	r.HandleFunc("/%s/{id}", handlers.Get%s).Methods("GET")
	r.HandleFunc("/%s/{id}", handlers.Update%s).Methods("PUT")
	r.HandleFunc("/%s/{id}", handlers.Delete%s).Methods("DELETE")
`,
			entity,
			entityPlural, entity,
			entityPlural, entity,
			entityPlural, entity,
			entityPlural, entity,
			entityPlural, entity))
	}

	routesList.WriteString("}\n")
	return routesList.String()
}

// GenerateTasksMarkdown creates a tasks.md file with lab instructions
func GenerateTasksMarkdown(entities []string) string {
	var tasks strings.Builder

	tasks.WriteString("# Lab Tasks\n\n")
	tasks.WriteString("## Overview\n\n")
	tasks.WriteString("This project implements a REST API with the following entities:\n\n")

	for _, entity := range entities {
		tasks.WriteString(fmt.Sprintf("- %s\n", entity))
	}

	tasks.WriteString("\n## Tasks to Complete\n\n")

	taskNum := 1
	for _, entity := range entities {
		entityLower := strings.ToLower(entity)
		tasks.WriteString(fmt.Sprintf("%d. Test all CRUD operations for %s\n", taskNum, entity))
		taskNum++
		tasks.WriteString(fmt.Sprintf("%d. Add validation for %s fields\n", taskNum, entityLower))
		taskNum++
	}

	tasks.WriteString(fmt.Sprintf("%d. Implement database persistence (currently using in-memory storage)\n", taskNum))
	taskNum++
	tasks.WriteString(fmt.Sprintf("%d. Add authentication middleware\n", taskNum))
	taskNum++
	tasks.WriteString(fmt.Sprintf("%d. Implement pagination for list endpoints\n", taskNum))
	taskNum++
	tasks.WriteString(fmt.Sprintf("%d. Add error handling and logging\n", taskNum))
	taskNum++
	tasks.WriteString(fmt.Sprintf("%d. Write integration tests\n", taskNum))
	taskNum++
	tasks.WriteString(fmt.Sprintf("%d. Add API documentation (Swagger/OpenAPI)\n", taskNum))

	tasks.WriteString("\n## Running the Application\n\n")
	tasks.WriteString("```bash\n")
	tasks.WriteString("# Install dependencies\n")
	tasks.WriteString("go mod tidy\n\n")
	tasks.WriteString("# Run the server\n")
	tasks.WriteString("go run cmd/main.go\n\n")
	tasks.WriteString("# Run tests\n")
	tasks.WriteString("go test ./...\n\n")
	tasks.WriteString("# Check coverage\n")
	tasks.WriteString("go test -cover ./...\n")
	tasks.WriteString("```\n\n")

	tasks.WriteString("## API Endpoints\n\n")

	for _, entity := range entities {
		entityLower := strings.ToLower(entity)
		entityPlural := pluralize(entity)

		tasks.WriteString(fmt.Sprintf("### %s\n\n", entity))
		tasks.WriteString(fmt.Sprintf("- `POST /%s` - Create a new %s\n", entityPlural, entityLower))
		tasks.WriteString(fmt.Sprintf("- `GET /%s` - Get all %s\n", entityPlural, entityPlural))
		tasks.WriteString(fmt.Sprintf("- `GET /%s/{id}` - Get a specific %s\n", entityPlural, entityLower))
		tasks.WriteString(fmt.Sprintf("- `PUT /%s/{id}` - Update a %s\n", entityPlural, entityLower))
		tasks.WriteString(fmt.Sprintf("- `DELETE /%s/{id}` - Delete a %s\n\n", entityPlural, entityLower))
	}

	return tasks.String()
}
