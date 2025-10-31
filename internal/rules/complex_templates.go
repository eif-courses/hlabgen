// Enhanced rule-based template generator for complex APIs
// Place this in: internal/rules/complex_templates.go

package rules

import (
	"fmt"
	"strings"
)

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
		http.Error(w, "invalid JSON: " + err.Error(), http.StatusBadRequest)
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
	// Rule-based approach: try to implement common discount patterns
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
		"draft":     {"submitted"},
		"submitted": {"under_review"},
		"under_review": {"approved", "rejected"},
		"rejected":  {"draft"},
		"approved":  {}, // Terminal state
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
				http.Error(w, "invalid state transition from " + currentState + " to " + newState, http.StatusBadRequest)
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

// Helper function (keep existing pluralize)
//func pluralize(s string) string {
//	l := strings.ToLower(s)
//	switch {
//	case strings.HasSuffix(l, "s") || strings.HasSuffix(l, "x") || strings.HasSuffix(l, "z"):
//		return l + "es"
//	case strings.HasSuffix(l, "y") && len(l) > 1:
//		return l[:len(l)-1] + "ies"
//	default:
//		return l + "s"
//	}
//}
