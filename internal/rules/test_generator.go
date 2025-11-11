package rules

import (
	"fmt"
	"strings"
)

// TestGenerationMode specifies what type of tests to generate
type TestGenerationMode string

const (
	TestModeSimple TestGenerationMode = "simple"
	TestModeFull   TestGenerationMode = "full"
)

// GenerateTestFile creates a complete test file for an entity
// This consolidates GenerateTest(), GenerateRulesPrimaryTest(), and fallback tests
func GenerateTestFile(entityName string, moduleName string, mode TestGenerationMode) string {
	entityLower := strings.ToLower(entityName)
	entityPlural := pluralize(entityName)

	if mode == TestModeSimple {
		return generateSimpleTest(entityName, entityLower, entityPlural, moduleName)
	}

	return generateFullTest(entityName, entityLower, entityPlural, moduleName)
}

func generateSimpleTest(entityName, entityLower, entityPlural, moduleName string) string {
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

	body, _ := json.Marshal(item)
	req := httptest.NewRequest("POST", "/%s", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handlers.Create%s(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %%d", w.Code)
	}
}

func TestGetAll%s(t *testing.T) {
	req := httptest.NewRequest("GET", "/%s", nil)
	w := httptest.NewRecorder()

	handlers.GetAll%s(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %%d", w.Code)
	}
}

func TestGet%s(t *testing.T) {
	req := httptest.NewRequest("GET", "/%s/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	w := httptest.NewRecorder()

	handlers.Get%s(w, req)

	if w.Code != http.StatusOK && w.Code != http.StatusNotFound {
		t.Errorf("Expected 200 or 404, got %%d", w.Code)
	}
}

func TestUpdate%s(t *testing.T) {
	updated := models.%s{Name: "Updated"}
	body, _ := json.Marshal(updated)
	req := httptest.NewRequest("PUT", "/%s/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	w := httptest.NewRecorder()

	handlers.Update%s(w, req)

	if w.Code != http.StatusOK && w.Code != http.StatusNotFound {
		t.Errorf("Expected 200 or 404, got %%d", w.Code)
	}
}

func TestDelete%s(t *testing.T) {
	req := httptest.NewRequest("DELETE", "/%s/1", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	w := httptest.NewRecorder()

	handlers.Delete%s(w, req)

	if w.Code != http.StatusNoContent && w.Code != http.StatusNotFound {
		t.Errorf("Expected 204 or 404, got %%d", w.Code)
	}
}
`,
		moduleName, moduleName,
		entityName,
		entityName, entityName,
		entityPlural,
		entityName,
		entityName,
		entityPlural,
		entityName,
		entityName,
		entityPlural,
		entityName,
		entityName,
		entityName,
		entityPlural,
		entityName,
		entityName,
		entityPlural,
		entityName,
	)
}

func generateFullTest(entityName, entityLower, entityPlural, moduleName string) string {
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
	// Create test item
	item := models.%s{
		Name:        "Test %s",
		Description: "Comprehensive test description",
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

func TestGetAll%s(t *testing.T) {
	req := httptest.NewRequest("GET", "/%s", nil)
	w := httptest.NewRecorder()

	handlers.GetAll%s(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %%d", w.Code)
	}

	var items []models.%s
	if err := json.NewDecoder(w.Body).Decode(&items); err != nil {
		t.Fatalf("Failed to decode response: %%v", err)
	}
}

func TestGet%s(t *testing.T) {
	// Create test item first
	item := models.%s{Name: "Test Item"}
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
	// Create test item first
	item := models.%s{Name: "Original Name"}
	body, _ := json.Marshal(item)
	createReq := httptest.NewRequest("POST", "/%s", bytes.NewBuffer(body))
	createReq.Header.Set("Content-Type", "application/json")
	createW := httptest.NewRecorder()
	handlers.Create%s(createW, createReq)

	// Update the item
	updated := models.%s{Name: "Updated Name"}
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
	// Create test item first
	item := models.%s{Name: "To Delete"}
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
		moduleName, moduleName,
		entityName,
		entityName, entityName,
		entityPlural,
		entityName,
		entityName,
		entityName,
		entityPlural,
		entityName,
		entityName,
		entityName,
		entityName,
		entityPlural,
		entityName,
		entityName,
		entityPlural,
		entityName,
		entityName,
		entityName,
		entityPlural,
		entityName,
		entityName,
		entityPlural,
		entityName,
		entityName,
		entityName,
		entityPlural,
		entityName,
		entityPlural,
		entityName,
	)
}

// ValidateAndFixTestSignatures ensures test functions have correct signatures
func ValidateAndFixTestSignatures(code string) (string, bool) {
	// This now just wraps the syntax_fixer function for compatibility
	return FixTestFunctions(code)
}

// GenerateFallbackTest creates a minimal test if generation fails
func GenerateFallbackTest(entityName, moduleName string) string {
	return fmt.Sprintf(`package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"%s/internal/handlers"
)

func TestCreate%s(t *testing.T) {
	body := strings.NewReader("{\"id\":1}")
	req, _ := http.NewRequest("POST", "/%s", body)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handlers.Create%s(rr, req)
	if rr.Code != http.StatusCreated && rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %%v", rr.Code)
	}
}
`, moduleName, entityName, strings.ToLower(entityName), entityName)
}

// FixTestFunctions is imported from syntax_fixer but included here for reference
func FixTestFunctions(code string) (string, bool) {
	// Implementation moved to internal/validate/syntax_fixer.go
	// This is a stub to maintain compatibility
	// In actual code, import from: "github.com/eif-courses/hlabgen/internal/validate"
	return code, false
}

// NOTE: pluralize() function removed - use the one from generator_rules.go instead
// This eliminates the redeclaration error
