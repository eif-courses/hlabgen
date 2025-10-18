package assemble

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/eif-courses/hlabgen/internal/ml"
	"github.com/eif-courses/hlabgen/internal/rules"
)

// File represents a generated file with path and content.
type File struct {
	Filename string
	Content  string
}

// ValidateAndFixTestFunctions checks for invalid test signatures and fixes them
func ValidateAndFixTestFunctions(code string) (string, bool) {
	fixed := false

	// Pattern 1: func TestXxx() { -> func TestXxx(t *testing.T) {
	pattern1 := regexp.MustCompile(`func (Test\w+)\(\) {`)
	if pattern1.MatchString(code) {
		code = pattern1.ReplaceAllString(code, `func $1(t *testing.T) {`)
		fixed = true
		fmt.Println("🔧 Fixed test signature: added missing (t *testing.T) parameter")
	}

	// Pattern 2: func TestXxx(t testing.T) { -> func TestXxx(t *testing.T) {
	pattern2 := regexp.MustCompile(`func (Test\w+)\(t testing\.T\) {`)
	if pattern2.MatchString(code) {
		code = pattern2.ReplaceAllString(code, `func $1(t *testing.T) {`)
		fixed = true
		fmt.Println("🔧 Fixed test signature: added missing pointer *")
	}

	// Pattern 3: Remove extra parameters from test functions
	pattern3 := regexp.MustCompile(`func (Test\w+)\(t \*testing\.T[^)]+\) {`)
	if pattern3.MatchString(code) {
		code = pattern3.ReplaceAllString(code, `func $1(t *testing.T) {`)
		fixed = true
		fmt.Println("🔧 Fixed test signature: removed extra parameters")
	}

	return code, fixed
}

// WriteMany writes multiple generated files to disk,
// applies rule-based safety fixes, and auto-fixes import paths.
func WriteMany(base string, files []File, metrics *ml.GenerationMetrics) error {
	if metrics == nil {
		metrics = &ml.GenerationMetrics{}
	}

	// Detect Go module name from go.mod
	moduleName, err := detectModule(base)
	if err != nil {
		fmt.Printf("⚠️  Could not detect module name: %v\n", err)
	}

	for _, f := range files {
		content := f.Content
		filename := f.Filename

		// ✅ Move tests next to handlers and adjust package name
		filename, content = rules.PlaceTestsWithHandlers(filename, content)
		metrics.RuleFixes++

		// ✅ Apply safety rule for handlers (decode + type mismatch fix)
		if strings.Contains(filename, "handlers/") && !strings.HasSuffix(filename, "_test.go") {
			before := content
			content = rules.SafeDecode(content)
			content = rules.FixIDTypeMismatch(content)
			content = removeUnusedModelsImport(content)
			if content != before {
				metrics.RuleFixes++
			}
		}

		// ✅ Remove unused imports from models
		if strings.Contains(filename, "models/") && !strings.HasSuffix(filename, "_test.go") {
			before := content
			content = cleanUnusedImportsInModels(content)
			content = ensureTimeImport(content)
			if content != before {
				metrics.RuleFixes++
			}
		}

		// ✅ Normalize routes: rename RegisterRoutes → Register
		if strings.Contains(filename, "routes.go") {
			before := content
			content = rules.FixRegisterFunction(content)
			if content != before {
				metrics.RuleFixes++
			}
		}

		// ✅ Apply test-specific cleanups
		if strings.HasSuffix(filename, "_test.go") {
			before := content

			// 🔧 NEW: Validate and fix test function signatures FIRST
			content, wasFixed := ValidateAndFixTestFunctions(content)
			if wasFixed {
				metrics.RuleFixes++
				fmt.Printf("✅ Auto-fixed test signatures in %s\n", filename)
			}

			// Then apply other test fixes
			content = rules.FixTestImports(content)
			content = rules.FixTestBodies(content)
			content = CleanDuplicateImports(content)
			if content != before {
				metrics.RuleFixes++
			}
		}

		// ✅ Remove unnecessary mux imports in handlers
		if strings.Contains(filename, "handlers/") && strings.Contains(content, `"github.com/gorilla/mux"`) {
			content = strings.ReplaceAll(content, "\t\"github.com/gorilla/mux\"\n", "")
			metrics.RuleFixes++
		}

		// ✅ Auto-fix placeholder import paths like "yourapp/", "your_project/", etc.
		if moduleName != "" {
			before := content
			content = strings.ReplaceAll(content, `"yourapp/routes"`, fmt.Sprintf(`"%s/internal/routes"`, moduleName))
			content = strings.ReplaceAll(content, `"yourapp/handlers"`, fmt.Sprintf(`"%s/internal/handlers"`, moduleName))
			content = strings.ReplaceAll(content, `"yourapp/models"`, fmt.Sprintf(`"%s/internal/models"`, moduleName))

			content = strings.ReplaceAll(content, `"your_project/routes"`, fmt.Sprintf(`"%s/internal/routes"`, moduleName))
			content = strings.ReplaceAll(content, `"your_project/handlers"`, fmt.Sprintf(`"%s/internal/handlers"`, moduleName))
			content = strings.ReplaceAll(content, `"your_project/models"`, fmt.Sprintf(`"%s/internal/models"`, moduleName))

			re := regexp.MustCompile(`"your[^"]+/`)
			content = re.ReplaceAllString(content, fmt.Sprintf(`"%s/internal/`, moduleName))

			if content != before {
				metrics.RuleFixes++
			}
		}

		// ✅ Final deduplication and brace fixes
		before := content
		content = CleanDuplicateImports(content)
		if content != before {
			metrics.RuleFixes++
		}

		// ✅ Normalize output paths (ensure /internal/ structure)
		fullPath := filepath.Join(base, rules.NormalizePath(filename))
		if err := os.MkdirAll(filepath.Dir(fullPath), 0o755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", filepath.Dir(fullPath), err)
		}

		// ✅ Write file to disk
		if err := os.WriteFile(fullPath, []byte(content), 0o644); err != nil {
			return fmt.Errorf("failed to write file %s: %w", fullPath, err)
		}
	}

	return nil
}

// detectModule reads the go.mod file and extracts the module name.
func detectModule(base string) (string, error) {
	data, err := os.ReadFile(filepath.Join(base, "go.mod"))
	if err != nil {
		return "", err
	}
	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module ")), nil
		}
	}
	return "", nil
}

// CleanDuplicateImports removes duplicate import lines (even if spacing differs)
// and ensures the file ends with balanced braces.
func CleanDuplicateImports(code string) string {
	lines := strings.Split(code, "\n")
	seen := make(map[string]bool)
	result := []string{}

	inImport := false
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Track when we enter/exit import blocks
		if strings.HasPrefix(trimmed, "import (") {
			inImport = true
			result = append(result, line)
			continue
		}
		if inImport && trimmed == ")" {
			inImport = false
			result = append(result, line)
			continue
		}

		// Deduplicate import statements
		if inImport && trimmed != "" {
			if seen[trimmed] {
				continue // Skip duplicate
			}
			seen[trimmed] = true
		}

		result = append(result, line)
	}

	return strings.Join(result, "\n")
}

// removeUnusedModelsImport removes "/internal/models" import if it's not used
func removeUnusedModelsImport(code string) string {
	if !strings.Contains(code, "models.") {
		lines := strings.Split(code, "\n")
		var result []string
		for _, line := range lines {
			if strings.Contains(line, "/internal/models\"") {
				continue
			}
			result = append(result, line)
		}
		return strings.Join(result, "\n")
	}
	return code
}

// cleanUnusedImportsInModels removes unused imports in model files
func cleanUnusedImportsInModels(code string) string {
	// Remove "time" import if not used
	if !strings.Contains(code, "time.") && !strings.Contains(code, "Time") {
		lines := strings.Split(code, "\n")
		var result []string
		for _, line := range lines {
			if strings.Contains(line, `"time"`) {
				continue
			}
			result = append(result, line)
		}
		code = strings.Join(result, "\n")
	}
	return code
}

// ensureTimeImport adds "time" import if time.Time is used but not imported
func ensureTimeImport(code string) string {
	if strings.Contains(code, "time.Time") && !strings.Contains(code, `"time"`) {
		// Find import block and add time
		lines := strings.Split(code, "\n")
		var result []string
		added := false
		for i, line := range lines {
			result = append(result, line)
			if !added && strings.Contains(line, "import (") {
				// Add time import after "import ("
				if i+1 < len(lines) {
					result = append(result, "\t\"time\"")
					added = true
				}
			}
		}
		if added {
			code = strings.Join(result, "\n")
		}
	}
	return code
}
