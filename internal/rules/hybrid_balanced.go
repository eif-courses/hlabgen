package rules

import (
	"fmt"
	"strings"
)

// GenerateHybridBalancedHandler creates handler with business logic hooks for ML
func GenerateHybridBalancedHandler(entity string, features []string) string {
	entityLower := strings.ToLower(entity)

	return fmt.Sprintf(`package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"MODULENAME/internal/models"
)

// In-memory storage
var items[]models.%s
var nextID = 1

// Create%s handles POST /%s with hybrid business logic
func Create%s(w http.ResponseWriter, r *http.Request) {
	// Step 1: Decode and validate (RULES)
	if r.Body == nil {
		http.Error(w, "missing request body", http.StatusBadRequest)
		return
	}

	var item models.%s
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "invalid JSON: " + err.Error(), http.StatusBadRequest)
		return
	}

	// Step 2: Basic validation (RULES)
	if item.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	// Step 3: Apply business logic (ML CUSTOMIZES THIS SECTION)
	applyBusinessLogic(&item)

	// Step 4: Persist (RULES)
	item.ID = nextID
	nextID++
	item.CreatedAt = time.Now()
	items = append(items, item)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

// GetAll%s handles GET /%ss (RULES)
func GetAll%s(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// Get%s handles GET /%ss/{id} (RULES)
func Get%s(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}

	for _, item := range items {
		if item.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	http.Error(w, "%s not found", http.StatusNotFound)
}

// Update%s handles PUT /%ss/{id} with hybrid business logic
func Update%s(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}

	var updated models.%s
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	// ML CUSTOMIZES THIS SECTION (business logic)
	applyBusinessLogic(&updated)

	for i, item := range items {
		if item.ID == id {
			updated.ID = id
			updated.CreatedAt = item.CreatedAt
			updated.UpdatedAt = time.Now()
			items[i] = updated

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updated)
			return
		}
	}

	http.Error(w, "%s not found", http.StatusNotFound)
}

// Delete%s handles DELETE /%ss/{id} (RULES)
func Delete%s(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}

	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "%s not found", http.StatusNotFound)
}

// ============================================================================
// ML CUSTOMIZES THIS SECTION
// ============================================================================

// applyBusinessLogic implements the business rules
// ML will replace the placeholder implementation below with actual logic
func applyBusinessLogic(item *models.%s) {
	// ML implements this function based on features: %v
	// 
	// This function should:
	// 1. Calculate discounts if applicable
	// 2. Apply tax calculations if needed
	// 3. Validate state transitions if workflow exists
	// 4. Implement any domain-specific rules
	//
	// Example implementations for different feature sets:
%s
	
	_ = item // Placeholder - ML replaces this
}
`, entity, entity, entityLower,
		entity,
		entity,
		entity, entityLower,
		entity,
		entity, entityLower,
		entity,
		entityLower,
		entity, entityLower,
		entity,
		entity,
		entityLower,
		entity, entityLower,
		entityLower,
		entity, features,
		generateExampleImplementations(features))
}

// generateExampleImplementations creates placeholder implementations for ML to replace
func generateExampleImplementations(features []string) string {
	var examples strings.Builder

	hasDiscount := false
	hasTax := false
	hasState := false

	for _, f := range features {
		lower := strings.ToLower(f)
		if strings.Contains(lower, "discount") {
			hasDiscount = true
		}
		if strings.Contains(lower, "tax") {
			hasTax = true
		}
		if strings.Contains(lower, "state") || strings.Contains(lower, "workflow") {
			hasState = true
		}
	}

	examples.WriteString("\t// Placeholder implementations for ML:\n")

	if hasDiscount {
		examples.WriteString(`	// Discount calculation (ML: replace with your logic):
	// if item.CustomerType == "premium" && item.Subtotal > 100 {
	//     item.Discount = item.Subtotal * 0.10
	// } else if item.Subtotal > 200 {
	//     item.Discount = item.Subtotal * 0.05
	// }
`)
	}

	if hasTax {
		examples.WriteString(`	// Tax calculation (ML: replace with your logic):
	// if item.Subtotal > 0 {
	//     item.Tax = (item.Subtotal - item.Discount) * 0.08
	// }
	// item.Total = item.Subtotal - item.Discount + item.Tax
`)
	}

	if hasState {
		examples.WriteString(`	// State transition validation (ML: replace with your logic):
	// if item.Status != "pending" && item.Status != "approved" && item.Status != "completed" {
	//     item.Status = "pending"
	// }
`)
	}

	if !hasDiscount && !hasTax && !hasState {
		examples.WriteString("\t// No special business logic needed - basic CRUD only\n")
	}

	return examples.String()
}

// GenerateHybridBalancedModel creates model with business logic fields
func GenerateHybridBalancedModel(entityName string, features []string) string {
	var fields strings.Builder

	fields.WriteString(fmt.Sprintf(`package models

import "time"

// %s represents the main entity with business logic fields
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
		if strings.Contains(lower, "priority") {
			fields.WriteString(`	Priority int ` + "`json:\"priority\"`" + `
`)
		}
	}

	fields.WriteString(`}
`)

	return fields.String()
}

// GenerateHybridBalancedLogicPrompt creates prompt for ML to implement just the business logic function
func GenerateHybridBalancedLogicPrompt(entityName string, features []string) string {
	return fmt.Sprintf(`Implement ONLY the applyBusinessLogic function for this entity:

Entity: %s
Features: %v

Generate the Go function body (NOT the function signature, just the body code):

Current function signature:
func applyBusinessLogic(item *models.%s) {
    // YOUR CODE HERE
}

Requirements:
1. Implement calculations: %s
2. Use item.* fields to access data
3. Modify item fields directly
4. DO NOT return anything (void function)
5. Handle nil/zero values gracefully
6. NO imports or package declaration needed

Example calculations:
%s

Return ONLY the function body code (everything between the braces).
NO function signature, NO explanations, ONLY Go code.`,
		entityName, features, entityName,
		extractCalculations(features),
		extractCalculationExamples(features))
}

// extractCalculations extracts what should be implemented
func extractCalculations(features []string) string {
	var calcs []string
	for _, f := range features {
		lower := strings.ToLower(f)
		if strings.Contains(lower, "discount") {
			calcs = append(calcs, "discount calculation")
		}
		if strings.Contains(lower, "tax") {
			calcs = append(calcs, "tax calculation")
		}
		if strings.Contains(lower, "state") || strings.Contains(lower, "workflow") {
			calcs = append(calcs, "state transition")
		}
	}
	if len(calcs) == 0 {
		calcs = append(calcs, "none - basic CRUD")
	}
	return strings.Join(calcs, ", ")
}

// extractCalculationExamples creates example implementations
func extractCalculationExamples(features []string) string {
	var examples strings.Builder

	for _, f := range features {
		lower := strings.ToLower(f)
		if strings.Contains(lower, "discount") {
			examples.WriteString(`// Discount: premium customers get 10%, others 5% for orders > 200
if item.CustomerType == "premium" {
    item.Discount = item.Subtotal * 0.10
} else if item.Subtotal > 200 {
    item.Discount = item.Subtotal * 0.05
}

`)
		}
		if strings.Contains(lower, "tax") {
			examples.WriteString(`// Tax: 8% of subtotal minus discount
taxableAmount := item.Subtotal - item.Discount
if taxableAmount > 0 {
    item.Tax = taxableAmount * 0.08
}
item.Total = item.Subtotal - item.Discount + item.Tax

`)
		}
		if strings.Contains(lower, "state") || strings.Contains(lower, "workflow") {
			examples.WriteString(`// State machine: pending -> approved -> completed
if item.Status == "" {
    item.Status = "pending"
}

`)
		}
	}

	return examples.String()
}
