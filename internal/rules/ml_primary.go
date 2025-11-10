package rules

import (
	"fmt"
	"strings"
)

// GenerateMLPrimaryHandler creates handler structure for ML to implement
func GenerateMLPrimaryHandler(entity string, features []string) string {
	entityLower := strings.ToLower(entity)

	return fmt.Sprintf(`package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"MODULENAME/internal/business"
	"MODULENAME/internal/models"
)

// Create%s handles POST /%s with ML-implemented business logic
func Create%s(w http.ResponseWriter, r *http.Request) {
	var req models.Create%sRequest
	
	// Validation (rules)
	if r.Body == nil {
		http.Error(w, "missing request body", http.StatusBadRequest)
		return
	}
	
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON: " + err.Error(), http.StatusBadRequest)
		return
	}

	// Validation (rules)
	if err := validate%s(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Business logic (ML implements this)
	result, err := business.Create%s(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

// Get%ss handles GET /%ss
func Get%ss(w http.ResponseWriter, r *http.Request) {
	items, err := business.GetAll%s(r.Context())
	if err != nil {
		http.Error(w, "failed to retrieve items", http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// Get%s handles GET /%ss/{id}
func Get%s(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}

	item, err := business.Get%s(r.Context(), id)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

// Update%s handles PUT /%ss/{id}
func Update%s(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}

	var req models.Update%sRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	result, err := business.Update%s(r.Context(), id, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// Delete%s handles DELETE /%ss/{id}
func Delete%s(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}

	if err := business.Delete%s(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ============================================================================
// Validation Rules (Rules-based - deterministic patterns)
// ============================================================================

func validate%s(req models.Create%sRequest) error {
	if req.Name == "" {
		return fmt.Errorf("name is required")
	}
	
	if len(req.Name) < 3 {
		return fmt.Errorf("name must be at least 3 characters")
	}
	
	if len(req.Name) > 255 {
		return fmt.Errorf("name must be at most 255 characters")
	}
	
	return nil
}
`, entity, strings.ToLower(entity), entity, entity,
		entity, entity,
		entity, entity,
		entity, strings.ToLower(entity),
		entity,
		entity, strings.ToLower(entity),
		entity, strings.ToLower(entity),
		entity,
		entity, strings.ToLower(entity),
		entity,
		entity,
		entity,
		entity, entity)
}

// BuildBusinessLogicPrompt creates ML prompt for implementing business logic
func BuildBusinessLogicPrompt(entityName string, features []string) string {
	return fmt.Sprintf(`You are a Go backend specialist. Implement the business logic package for this API:

Entity: %s
Features: %v

Generate COMPLETE, PRODUCTION-READY code for: internal/business/logic.go

This package should contain these functions:
1. Create%s(ctx context.Context, req models.Create%sRequest) (*models.%s, error)
2. GetAll%s(ctx context.Context) ([]models.%s, error)
3. Get%s(ctx context.Context, id int) (*models.%s, error)
4. Update%s(ctx context.Context, id int, req models.Update%sRequest) (*models.%s, error)
5. Delete%s(ctx context.Context, id int) error

Business Logic Requirements:
- Use context for cancellation and timeouts
- Validate all inputs thoroughly
- Return descriptive errors for invalid states
- Handle edge cases properly
- Implement any calculations (discount, tax, pricing) based on features: %v
- Implement state transitions if workflow/status features exist
- Use in-memory storage (slice) for now

Calculations to Implement:
%s

Key Rules:
- ALWAYS validate inputs before processing
- Return proper error messages
- Handle empty/nil cases gracefully
- Use consistent error handling pattern

Return ONLY the Go code for the business package. Start with 'package business' and include all imports.
NO explanations, NO markdown blocks, ONLY valid Go code.`,
		entityName, features,
		entityName, entityName, entityName,
		entityName, entityName,
		entityName, entityName,
		entityName, entityName, entityName,
		entityName,
		features,
		extractCalculationRules(features))
}

// extractCalculationRules extracts specific rules from features
func extractCalculationRules(features []string) string {
	var rules []string

	for _, f := range features {
		lower := strings.ToLower(f)
		if strings.Contains(lower, "discount") {
			rules = append(rules, "- Apply discount based on customer type (premium: 10%, standard: 5%)")
		}
		if strings.Contains(lower, "tax") {
			rules = append(rules, "- Calculate tax as 8% of subtotal minus discount")
		}
		if strings.Contains(lower, "pricing") {
			rules = append(rules, "- Ensure price > 0, calculate total = base + tax - discount")
		}
		if strings.Contains(lower, "workflow") || strings.Contains(lower, "state") {
			rules = append(rules, "- Validate state transitions: pending -> approved -> completed")
		}
	}

	if len(rules) == 0 {
		rules = append(rules, "- No special calculations, just CRUD operations")
	}

	return strings.Join(rules, "\n")
}

// GenerateMLPrimaryModel creates model with business logic fields
func GenerateMLPrimaryModel(entityName string, features []string) string {
	var fields strings.Builder

	fields.WriteString(fmt.Sprintf(`package models

import (
	"context"
	"time"
)

// %s represents the main entity
type %s struct {
	ID        int       `+"`json:\"id\"`"+`
	Name      string    `+"`json:\"name\"`"+`
	CreatedAt time.Time `+"`json:\"created_at\"`"+`
	UpdatedAt time.Time `+"`json:\"updated_at\"`"+`
`, entityName, entityName))

	// Add business logic fields based on features
	for _, f := range features {
		lower := strings.ToLower(f)
		if strings.Contains(lower, "discount") {
			fields.WriteString(`	Subtotal     float64 ` + "`json:\"subtotal\"`" + `
	Discount     float64 ` + "`json:\"discount\"`" + `
	CustomerType string  ` + "`json:\"customer_type,omitempty\"`" + `
`)
		}
		if strings.Contains(lower, "tax") {
			fields.WriteString(`	Tax   float64 ` + "`json:\"tax\"`" + `
	Total float64 ` + "`json:\"total\"`" + `
`)
		}
		if strings.Contains(lower, "state") || strings.Contains(lower, "workflow") {
			fields.WriteString(`	Status string ` + "`json:\"status\"`" + `
`)
		}
	}

	fields.WriteString(`}

// Create` + entityName + `Request is the input for creation
type Create` + entityName + `Request struct {
	Name string ` + "`json:\"name\"`" + `
`)

	for _, f := range features {
		lower := strings.ToLower(f)
		if strings.Contains(lower, "discount") {
			fields.WriteString(`	Subtotal     float64 ` + "`json:\"subtotal\"`" + `
	CustomerType string  ` + "`json:\"customer_type,omitempty\"`" + `
`)
		}
	}

	fields.WriteString(`}

// Update` + entityName + `Request is the input for updates
type Update` + entityName + `Request struct {
	Name string ` + "`json:\"name\"`" + `
`)

	for _, f := range features {
		lower := strings.ToLower(f)
		if strings.Contains(lower, "state") || strings.Contains(lower, "workflow") {
			fields.WriteString(`	Status string ` + "`json:\"status\"`" + `
`)
		}
	}

	fields.WriteString(`}
`)

	return fields.String()
}
