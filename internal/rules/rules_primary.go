package rules

import (
	"fmt"
	"strings"
)

// GenerateRulesPrimaryModel creates a simple CRUD model without business logic
func GenerateRulesPrimaryModel(entityName string) string {
	return fmt.Sprintf(`package models

import "time"

// %s represents a simple CRUD entity
type %s struct {
	ID        int       `+"`json:\"id\"`"+`
	Name      string    `+"`json:\"name\"`"+`
	Description string `+"`json:\"description,omitempty\"`"+`
	CreatedAt time.Time `+"`json:\"created_at\"`"+`
	UpdatedAt time.Time `+"`json:\"updated_at\"`"+`
}
`, entityName, entityName)
}

// GenerateRulesPrimaryHandler creates simple CRUD handlers (no business logic)
func GenerateRulesPrimaryHandler(entityName string, moduleName string) string {
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

// In-memory storage (RULES-PRIMARY: simple storage, no business logic)
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
		http.Error(w, "invalid JSON: " + err.Error(), http.StatusBadRequest)
		return
	}

	// Validation (RULES-PRIMARY: deterministic only)
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

// GetAll%s handles GET /%s (RULES-PRIMARY: simple retrieval)
func GetAll%s(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(%s)
}

// Get%s handles GET /%s/{id} (RULES-PRIMARY: simple lookup)
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

// Update%s handles PUT /%s/{id} (RULES-PRIMARY: simple update)
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

	// Validation (RULES-PRIMARY)
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

// Delete%s handles DELETE /%s/{id} (RULES-PRIMARY: simple deletion)
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
		entityName,
		entityPlural, entityName,
		entityName,
		entityLower,
		entityName, entityPlural,
		entityName,
		entityLower,
		entityName, entityLower,
		entityLower,
		entityName, entityPlural,
		entityName,
		entityLower,
		entityName, entityPlural,
		entityName,
		entityLower,
		entityName, entityPlural,
		entityName,
		entityLower,
		entityLower, entityLower, entityLower,
		entityLower)
}

// GenerateRulesPrimaryTest creates simple CRUD tests (no business logic testing)
func GenerateRulesPrimaryTest(entityName string, moduleName string) string {
	entityLower := strings.ToLower(entityName)
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

// TestCreate%s (RULES-PRIMARY: basic creation)
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

// TestGetAll%s (RULES-PRIMARY: simple retrieval)
func TestGetAll%s(t *testing.T) {
	req := httptest.NewRequest("GET", "/%s", nil)
	w := httptest.NewRecorder()

	handlers.GetAll%s(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %%d", w.Code)
	}
}

// TestGet%s (RULES-PRIMARY: simple lookup)
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

// TestUpdate%s (RULES-PRIMARY: simple update)
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

// TestDelete%s (RULES-PRIMARY: simple deletion)
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
		entityName,
		entityPlural, entityName,
		entityPlural,
		entityName,
		entityName,
		entityName,
		entityPlural,
		entityName,
		entityPlural,
		entityName,
		entityName, entityPlural,
		entityName,
		entityName,
		entityPlural,
		entityName,
		entityName, entityPlural,
		entityName,
		entityName,
		entityName,
		entityPlural,
		entityName,
		entityPlural,
		entityName)
}
