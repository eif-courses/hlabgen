// Package rules provides rule-based code generation for REST APIs
// This file consolidates simple and complex CRUD generation patterns
package rules

import (
	"fmt"
	"strings"
)

// =============================================================================
// SIMPLE CRUD GENERATION (No Business Logic)
// =============================================================================

// GenerateSimpleModel creates a basic model struct without business logic
func GenerateSimpleModel(entityName string) string {
	return fmt.Sprintf(`package models

import "time"

// %s represents a simple CRUD entity
type %s struct {
	ID          int       `+"`json:\"id\"`"+`
	Name        string    `+"`json:\"name\"`"+`
	Description string    `+"`json:\"description,omitempty\"`"+`
	CreatedAt   time.Time `+"`json:\"created_at\"`"+`
	UpdatedAt   time.Time `+"`json:\"updated_at\"`"+`
}
`, entityName, entityName)
}

// GenerateSimpleHandler creates basic CRUD handlers without business logic
func GenerateSimpleHandler(entityName string, moduleName string) string {
	entityLower := strings.ToLower(entityName)
	entityPlural := pluralize(entityName)

	return fmt.Sprintf(`package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"%s/internal/models"
)

// In-memory storage (simple storage, no business logic)
var %s []models.%s
var next%sID = 1

// Create%s handles POST /%s
func Create%s(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "missing request body", http.StatusBadRequest)
		return
	}

	var item models.%s
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validation (deterministic only)
	if item.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	if len(item.Name) < 3 {
		http.Error(w, "name must be at least 3 characters", http.StatusBadRequest)
		return
	}

	if len(item.Name) > 255 {
		http.Error(w, "name must be at most 255 characters", http.StatusBadRequest)
		return
	}

	// Persist
	item.ID = next%sID
	next%sID++
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()
	%s = append(%s, item)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

// GetAll%s handles GET /%s
func GetAll%s(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(%s)
}

// Get%s handles GET /%s/{id}
func Get%s(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "invalid ID format", http.StatusBadRequest)
		return
	}

	for _, item := range %s {
		if item.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
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
		http.Error(w, "invalid ID format", http.StatusBadRequest)
		return
	}

	if r.Body == nil {
		http.Error(w, "missing request body", http.StatusBadRequest)
		return
	}

	var updated models.%s
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	// Validation
	if updated.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	for i, item := range %s {
		if item.ID == id {
			updated.ID = id
			updated.CreatedAt = item.CreatedAt
			updated.UpdatedAt = time.Now()
			%s[i] = updated

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updated)
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
		http.Error(w, "invalid ID format", http.StatusBadRequest)
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
		moduleName,
		entityLower, entityName,
		entityName,
		entityName, entityPlural,
		entityName,
		entityName,
		entityName, entityName,
		entityLower, entityLower,
		entityName, entityPlural,
		entityName,
		entityLower,
		entityName, entityPlural,
		entityName,
		entityLower,
		entityName,
		entityName, entityPlural,
		entityName,
		entityName,
		entityLower,
		entityLower,
		entityName,
		entityName, entityPlural,
		entityName,
		entityLower,
		entityLower, entityLower, entityLower,
		entityLower)
}

// GenerateSimpleTest creates basic CRUD tests without business logic
func GenerateSimpleTest(entityName string, moduleName string) string {
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

// TestCreate%s tests basic creation
func TestCreate%s(t *testing.T) {
	item := models.%s{
		Name:        "Test %s",
		Description: "Test Description",
	}

	body, _ := json.Marshal(item)
	req := httptest.NewRequest("POST", "/%s", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handlers.Create%s(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %%d", w.Code)
	}

	var response models.%s
	json.NewDecoder(w.Body).Decode(&response)
	if response.ID == 0 {
		t.Error("Expected ID to be set")
	}
}

// TestGetAll%s tests simple retrieval
func TestGetAll%s(t *testing.T) {
	req := httptest.NewRequest("GET", "/%s", nil)
	w := httptest.NewRecorder()

	handlers.GetAll%s(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %%d", w.Code)
	}
}

// TestGet%s tests simple lookup
func TestGet%s(t *testing.T) {
	// Create test item first
	item := models.%s{Name: "Test"}
	body, _ := json.Marshal(item)
	createReq := httptest.NewRequest("POST", "/%s", bytes.NewBuffer(body))
	createReq.Header.Set("Content-Type", "application/json")
	createW := httptest.NewRecorder()
	handlers.Create%s(createW, createReq)

	// Now retrieve it
	req := httptest.NewRequest("GET", "/%s/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	w := httptest.NewRecorder()

	handlers.Get%s(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %%d", w.Code)
	}
}

// TestUpdate%s tests simple update
func TestUpdate%s(t *testing.T) {
	// Create test item
	item := models.%s{Name: "Original"}
	body, _ := json.Marshal(item)
	createReq := httptest.NewRequest("POST", "/%s", bytes.NewBuffer(body))
	createReq.Header.Set("Content-Type", "application/json")
	createW := httptest.NewRecorder()
	handlers.Create%s(createW, createReq)

	// Update it
	updated := models.%s{Name: "Updated"}
	updateBody, _ := json.Marshal(updated)
	req := httptest.NewRequest("PUT", "/%s/1", bytes.NewBuffer(updateBody))
	req.Header.Set("Content-Type", "application/json")
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	w := httptest.NewRecorder()

	handlers.Update%s(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %%d", w.Code)
	}
}

// TestDelete%s tests simple deletion
func TestDelete%s(t *testing.T) {
	// Create test item
	item := models.%s{Name: "ToDelete"}
	body, _ := json.Marshal(item)
	createReq := httptest.NewRequest("POST", "/%s", bytes.NewBuffer(body))
	createReq.Header.Set("Content-Type", "application/json")
	createW := httptest.NewRecorder()
	handlers.Create%s(createW, createReq)

	// Delete it
	req := httptest.NewRequest("DELETE", "/%s/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	w := httptest.NewRecorder()

	handlers.Delete%s(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("Expected 204, got %%d", w.Code)
	}
}
`,
		moduleName, moduleName,
		entityName, entityName,
		entityName, entityName,
		entityPlural,
		entityName,
		entityName,
		entityName, entityName,
		entityPlural,
		entityName,
		entityName, entityName,
		entityName,
		entityPlural,
		entityName,
		entityPlural,
		entityName,
		entityName, entityName,
		entityName,
		entityPlural,
		entityName,
		entityName,
		entityPlural,
		entityName,
		entityName, entityName,
		entityName,
		entityPlural,
		entityName,
		entityPlural,
		entityName)
}

// =============================================================================
// COMPLEX CRUD GENERATION (With Business Logic)
// =============================================================================

// ComplexHandlerGenerator generates handlers with business logic patterns
type ComplexHandlerGenerator struct {
	EntityName    string
	ModuleName    string
	Features      []string
	HasDiscount   bool
	HasTax        bool
	HasState      bool
	HasAuth       bool
	HasValidation bool
}

// NewComplexHandler analyzes features and creates appropriate handler
func NewComplexHandler(entityName, moduleName string, features []string) *ComplexHandlerGenerator {
	gen := &ComplexHandlerGenerator{
		EntityName: entityName,
		ModuleName: moduleName,
		Features:   features,
	}

	// Analyze features to detect patterns
	for _, feature := range features {
		lower := strings.ToLower(feature)

		if strings.Contains(lower, "discount") {
			gen.HasDiscount = true
		}
		if strings.Contains(lower, "tax") {
			gen.HasTax = true
		}
		if strings.Contains(lower, "state") || strings.Contains(lower, "status") ||
			strings.Contains(lower, "workflow") || strings.Contains(lower, "transition") {
			gen.HasState = true
		}
		if strings.Contains(lower, "auth") || strings.Contains(lower, "permission") ||
			strings.Contains(lower, "jwt") {
			gen.HasAuth = true
		}
		if strings.Contains(lower, "validate") || strings.Contains(lower, "required") ||
			strings.Contains(lower, "min") || strings.Contains(lower, "max") {
			gen.HasValidation = true
		}
	}

	return gen
}

// GenerateComplexHandler creates handler with business logic
func (g *ComplexHandlerGenerator) GenerateComplexHandler() string {
	var buf strings.Builder

	buf.WriteString(fmt.Sprintf(`package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
	
	"%s/internal/models"
	"github.com/gorilla/mux"
)

`, g.ModuleName))

	// Generate Create handler with business logic
	buf.WriteString(g.generateCreateHandler())
	buf.WriteString("\n\n")

	// Generate other CRUD operations
	buf.WriteString(g.generateGetAllHandler())
	buf.WriteString("\n\n")
	buf.WriteString(g.generateGetByIDHandler())
	buf.WriteString("\n\n")
	buf.WriteString(g.generateUpdateHandler())
	buf.WriteString("\n\n")
	buf.WriteString(g.generateDeleteHandler())

	return buf.String()
}

func (g *ComplexHandlerGenerator) generateCreateHandler() string {
	entityLower := strings.ToLower(g.EntityName)
	entityVar := entityLower

	var buf strings.Builder

	buf.WriteString(fmt.Sprintf(`func Create%s(w http.ResponseWriter, r *http.Request) {
	// Request validation
	if r.Body == nil {
		http.Error(w, "missing request body", http.StatusBadRequest)
		return
	}
	
	var %s models.%s
	if err := json.NewDecoder(r.Body).Decode(&%s); err != nil {
		http.Error(w, "invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
`, g.EntityName, entityVar, g.EntityName, entityVar))

	// Add validation logic if needed
	if g.HasValidation {
		buf.WriteString(g.generateValidationLogic(entityVar))
	}

	// Add discount calculation if needed
	if g.HasDiscount {
		buf.WriteString(g.generateDiscountLogic(entityVar))
	}

	// Add tax calculation if needed
	if g.HasTax {
		buf.WriteString(g.generateTaxLogic(entityVar))
	}

	// Add state initialization if needed
	if g.HasState {
		buf.WriteString(g.generateStateInitialization(entityVar))
	}

	buf.WriteString(fmt.Sprintf(`
	// Set metadata
	%s.ID = generateID()
	%s.CreatedAt = time.Now()
	
	// Persist (in-memory for now)
	store%s = append(store%s, %s)
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(%s)
}`, entityVar, entityVar, g.EntityName, g.EntityName, entityVar, entityVar))

	return buf.String()
}

func (g *ComplexHandlerGenerator) generateValidationLogic(varName string) string {
	return fmt.Sprintf(`
	// Business rule validation
	if %s.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}
	
	// Additional field validation based on features
	// Note: Rule-based approach uses pattern matching, may not capture all nuances
`, varName)
}

func (g *ComplexHandlerGenerator) generateDiscountLogic(varName string) string {
	return fmt.Sprintf(`
	// Business rule: Discount calculation (rule-based pattern matching)
	// Attempting to implement discount logic based on common patterns
	var discount float64
	
	// Pattern 1: Premium customer discount
	if %s.CustomerType == "premium" && %s.Subtotal > 100 {
		discount = %s.Subtotal * 0.10 // 10%% discount
	}
	
	// Pattern 2: Volume discount
	if %s.Subtotal > 200 {
		discount = %s.Subtotal * 0.05 // 5%% discount
	}
	
	%s.Discount = discount
	%s.Total = %s.Subtotal - discount
`, varName, varName, varName, varName, varName, varName, varName, varName)
}

func (g *ComplexHandlerGenerator) generateTaxLogic(varName string) string {
	return fmt.Sprintf(`
	// Business rule: Tax calculation (rule-based fixed rate)
	// Assumes standard tax rate of 8%%
	taxableAmount := %s.Subtotal - %s.Discount
	%s.Tax = taxableAmount * 0.08
	%s.Total = taxableAmount + %s.Tax
`, varName, varName, varName, varName, varName)
}

func (g *ComplexHandlerGenerator) generateStateInitialization(varName string) string {
	return fmt.Sprintf(`
	// Business rule: State machine initialization
	// Rule-based approach sets initial state
	if %s.Status == "" {
		%s.Status = "pending" // Default initial state
	}
	
	// Validate state transition (basic pattern)
	validStates := map[string]bool{
		"pending": true, "draft": true, "submitted": true,
		"approved": true, "rejected": true, "completed": true,
	}
	if !validStates[%s.Status] {
		http.Error(w, "invalid status value", http.StatusBadRequest)
		return
	}
`, varName, varName, varName)
}

func (g *ComplexHandlerGenerator) generateGetAllHandler() string {
	return fmt.Sprintf(`func GetAll%s(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(store%s)
}`, g.EntityName, g.EntityName)
}

func (g *ComplexHandlerGenerator) generateGetByIDHandler() string {
	entityLower := strings.ToLower(g.EntityName)

	return fmt.Sprintf(`func Get%sByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	for _, item := range store%s {
		if item.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	
	http.Error(w, "%s not found", http.StatusNotFound)
}`, g.EntityName, g.EntityName, entityLower)
}

func (g *ComplexHandlerGenerator) generateUpdateHandler() string {
	entityLower := strings.ToLower(g.EntityName)
	entityVar := entityLower

	var stateValidation string
	if g.HasState {
		stateValidation = `
	// State transition validation (rule-based patterns)
	validTransitions := map[string][]string{
		"draft":        {"submitted"},
		"submitted":    {"under_review"},
		"under_review": {"approved", "rejected"},
		"rejected":     {"draft"},
		"approved":     {}, // Terminal state
	}
	
	for i, item := range store` + g.EntityName + ` {
		if item.ID == id {
			currentState := item.Status
			newState := ` + entityVar + `.Status
			
			// Check if transition is valid
			allowed := validTransitions[currentState]
			validTransition := false
			for _, state := range allowed {
				if state == newState {
					validTransition = true
					break
				}
			}
			
			if !validTransition && currentState != newState {
				http.Error(w, "invalid state transition from "+currentState+" to "+newState, http.StatusBadRequest)
				return
			}
			
			store` + g.EntityName + `[i] = ` + entityVar + `
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(` + entityVar + `)
			return
		}
	}`
	} else {
		stateValidation = `
	for i, item := range store` + g.EntityName + ` {
		if item.ID == id {
			store` + g.EntityName + `[i] = ` + entityVar + `
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(` + entityVar + `)
			return
		}
	}`
	}

	return fmt.Sprintf(`func Update%s(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	if r.Body == nil {
		http.Error(w, "missing request body", http.StatusBadRequest)
		return
	}
	
	var %s models.%s
	if err := json.NewDecoder(r.Body).Decode(&%s); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	
	%s.ID = id
	%s
	
	http.Error(w, "%s not found", http.StatusNotFound)
}`, g.EntityName, entityVar, g.EntityName, entityVar, entityVar, stateValidation, entityLower)
}

func (g *ComplexHandlerGenerator) generateDeleteHandler() string {
	entityLower := strings.ToLower(g.EntityName)

	return fmt.Sprintf(`func Delete%s(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	for i, item := range store%s {
		if item.ID == id {
			store%s = append(store%s[:i], store%s[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	
	http.Error(w, "%s not found", http.StatusNotFound)
}`, g.EntityName, g.EntityName, g.EntityName, g.EntityName, g.EntityName, entityLower)
}

// GenerateComplexModel creates model with business logic fields
func GenerateComplexModel(entityName string, features []string) string {
	gen := NewComplexHandler(entityName, "", features)

	var buf strings.Builder
	buf.WriteString(fmt.Sprintf(`package models

import "time"

// %s represents a %s entity with business logic
type %s struct {
	ID          string    `+"`json:\"id\"`"+`
	Name        string    `+"`json:\"name\"`"+`
	Description string    `+"`json:\"description,omitempty\"`"+`
`, entityName, strings.ToLower(entityName), entityName))

	// Add business logic fields based on features
	if gen.HasDiscount {
		buf.WriteString(`	Subtotal     float64   ` + "`json:\"subtotal\"`" + `
	Discount     float64   ` + "`json:\"discount\"`" + `
	CustomerType string    ` + "`json:\"customer_type,omitempty\"`" + `
`)
	}

	if gen.HasTax {
		buf.WriteString(`	Tax   float64 ` + "`json:\"tax\"`" + `
	Total float64 ` + "`json:\"total\"`" + `
`)
	}

	if gen.HasState {
		buf.WriteString(`	Status string ` + "`json:\"status\"`" + `
`)
	}

	buf.WriteString(`	CreatedAt time.Time ` + "`json:\"created_at,omitempty\"`" + `
	UpdatedAt time.Time ` + "`json:\"updated_at,omitempty\"`" + `
}

// In-memory storage
var store` + entityName + ` []` + entityName + `

// Helper function
func generateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
`)

	return buf.String()
}

// =============================================================================
// ROUTES GENERATION
// =============================================================================

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
	r.HandleFunc("/%s", handlers.GetAll%s).Methods("GET")
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

// =============================================================================
// DOCUMENTATION GENERATION
// =============================================================================

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

// =============================================================================
// HELPER FUNCTIONS
// =============================================================================

// pluralize converts singular entity names to plural
func pluralize(s string) string {
	l := strings.ToLower(s)
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

// isVowel checks if a byte is a vowel
func isVowel(b byte) bool {
	switch b {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}
