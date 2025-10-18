package assemble

import (
	"fmt"
	"go/parser"
	"go/token"
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
	originalCode := code
	lines := strings.Split(code, "\n")

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if !strings.HasPrefix(trimmed, "func Test") {
			continue
		}

		funcNameMatch := regexp.MustCompile(`func (Test\w+)\s*\(`).FindStringSubmatch(trimmed)
		if len(funcNameMatch) < 2 {
			continue
		}
		funcName := funcNameMatch[1]

		if strings.Contains(trimmed, "func "+funcName+"(t *testing.T)") {
			continue
		}

		oldLine := line
		if strings.Contains(trimmed, "{") {
			lines[i] = regexp.MustCompile(`func `+regexp.QuoteMeta(funcName)+`\s*\([^)]*\)\s*\{`).
				ReplaceAllString(line, `func `+funcName+`(t *testing.T) {`)
		} else {
			lines[i] = regexp.MustCompile(`func `+regexp.QuoteMeta(funcName)+`\s*\([^)]*\)\s*$`).
				ReplaceAllString(line, `func `+funcName+`(t *testing.T)`)
		}

		if lines[i] != oldLine {
			fixed = true
			fmt.Printf("🔧 Fixed test signature for %s\n", funcName)
		}
	}

	code = strings.Join(lines, "\n")

	// Fix missing commas
	code, commaFixed := fixMissingCommas(code)
	if commaFixed {
		fixed = true
	}

	// Ensure testing import exists if we made changes
	if code != originalCode && strings.Contains(code, "func Test") {
		if !strings.Contains(code, `"testing"`) {
			code = ensureTestingImport(code)
			fixed = true
			fmt.Println("🔧 Added missing testing import")
		}
	}

	return code, fixed
}

// FixHandlerSignatures ensures all handler functions have the correct signature
func FixHandlerSignatures(code string) (string, bool) {
	fixed := false
	lines := strings.Split(code, "\n")

	// Pattern to match handler functions without parameters
	handlerPattern := regexp.MustCompile(`^func\s+(Create|Get|Update|Delete|List)\w+\(\s*\)\s*{`)

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)

		if handlerPattern.MatchString(trimmed) {
			// Extract function name
			match := regexp.MustCompile(`func\s+(\w+)`).FindStringSubmatch(trimmed)
			if len(match) > 1 {
				funcName := match[1]
				// Replace with correct signature
				oldLine := line
				lines[i] = regexp.MustCompile(`func\s+`+regexp.QuoteMeta(funcName)+`\s*\(\s*\)\s*{`).
					ReplaceAllString(line, `func `+funcName+`(w http.ResponseWriter, r *http.Request) {`)

				if lines[i] != oldLine {
					fixed = true
					fmt.Printf("🔧 Fixed handler signature for %s\n", funcName)
				}
			}
		}
	}

	code = strings.Join(lines, "\n")

	// Ensure net/http import exists if we made changes
	if fixed && !strings.Contains(code, `"net/http"`) {
		code = ensureNetHTTPImport(code)
		fmt.Println("🔧 Added missing net/http import")
	}

	return code, fixed
}

// ensureNetHTTPImport adds net/http import if missing
func ensureNetHTTPImport(code string) string {
	if strings.Contains(code, `"net/http"`) {
		return code
	}

	lines := strings.Split(code, "\n")
	for i, line := range lines {
		if strings.Contains(line, "import (") {
			// Check if net/http already exists
			hasHTTP := false
			for j := i + 1; j < len(lines) && j < i+15; j++ {
				if strings.Contains(lines[j], `"net/http"`) {
					hasHTTP = true
					break
				}
				if strings.TrimSpace(lines[j]) == ")" {
					break
				}
			}
			if !hasHTTP {
				lines = append(lines[:i+1], append([]string{"\t\"net/http\""}, lines[i+1:]...)...)
			}
			return strings.Join(lines, "\n")
		}
	}

	// No import block found, create one
	for i, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), "package ") {
			lines = append(lines[:i+1], append([]string{"", "import (", "\t\"net/http\"", ")"}, lines[i+1:]...)...)
			return strings.Join(lines, "\n")
		}
	}

	return code
}

// fixMissingCommas adds missing trailing commas in struct literals ONLY
func fixMissingCommas(code string) (string, bool) {
	fixed := false
	lines := strings.Split(code, "\n")

	inStructLiteral := false
	structBraceCount := 0

	for i := 0; i < len(lines)-1; i++ {
		currentLine := lines[i]
		trimmed := strings.TrimSpace(currentLine)
		nextLineTrimmed := strings.TrimSpace(lines[i+1])

		// Skip comments and empty lines
		if strings.HasPrefix(trimmed, "//") || trimmed == "" {
			continue
		}

		// Detect struct literal start: = SomeType{ or := SomeType{
		if (strings.Contains(trimmed, "= models.") ||
			strings.Contains(trimmed, ":= models.") ||
			strings.Contains(trimmed, "= []models.")) &&
			strings.HasSuffix(trimmed, "{") {
			inStructLiteral = true
			structBraceCount = 1
			continue
		}

		// Track brace depth inside struct literals
		if inStructLiteral {
			structBraceCount += strings.Count(trimmed, "{")
			structBraceCount -= strings.Count(trimmed, "}")

			// Exit struct literal when braces are balanced
			if structBraceCount == 0 {
				inStructLiteral = false
				continue
			}

			// Only fix commas INSIDE struct literals
			if structBraceCount > 0 {
				// Already has comma
				if strings.HasSuffix(trimmed, ",") {
					continue
				}

				// Opening braces - don't add comma
				if strings.HasSuffix(trimmed, "{") {
					continue
				}

				// Check if next line is closing brace
				if nextLineTrimmed == "}" || nextLineTrimmed == "}," ||
					nextLineTrimmed == "})," || nextLineTrimmed == "})" {
					// This is a field that needs a comma before closing brace
					if strings.Contains(trimmed, ":") ||
						(strings.Contains(trimmed, "{") && strings.Contains(trimmed, "}")) {
						lines[i] = currentLine + ","
						fixed = true
						fmt.Printf("🔧 Fixed missing comma in struct literal at line %d\n", i+1)
					}
				}

				// Check if next line is another field (contains : and is not a label)
				if strings.Contains(nextLineTrimmed, ":") &&
					!strings.HasPrefix(nextLineTrimmed, "//") &&
					!strings.HasSuffix(nextLineTrimmed, ":") { // exclude labels like "default:"
					if strings.Contains(trimmed, ":") {
						lines[i] = currentLine + ","
						fixed = true
						fmt.Printf("🔧 Fixed missing comma between fields at line %d\n", i+1)
					}
				}

				// Special case: inline struct literal followed by closing brace
				if strings.Contains(currentLine, "{") &&
					strings.Contains(currentLine, "}") &&
					!strings.HasSuffix(trimmed, "},") &&
					(nextLineTrimmed == "}" || nextLineTrimmed == "},") {
					if !strings.HasSuffix(trimmed, ",") {
						lines[i] = currentLine + ","
						fixed = true
						fmt.Printf("🔧 Fixed missing comma after inline struct at line %d\n", i+1)
					}
				}
			}
		}
	}

	if fixed {
		fmt.Println("✅ Fixed missing commas in struct literals")
	}

	return strings.Join(lines, "\n"), fixed
}

// ensureTestingImport adds "testing" import if missing
func ensureTestingImport(code string) string {
	if strings.Contains(code, `"testing"`) {
		return code
	}

	lines := strings.Split(code, "\n")
	for i, line := range lines {
		if strings.Contains(line, "import (") {
			hasTestingImport := false
			for j := i + 1; j < len(lines) && j < i+10; j++ {
				if strings.Contains(lines[j], `"testing"`) {
					hasTestingImport = true
					break
				}
				if strings.TrimSpace(lines[j]) == ")" {
					break
				}
			}
			if !hasTestingImport {
				lines = append(lines[:i+1], append([]string{"\t\"testing\""}, lines[i+1:]...)...)
			}
			return strings.Join(lines, "\n")
		}
	}

	for i, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), "package ") {
			lines = append(lines[:i+1], append([]string{"", "import (", "\t\"testing\"", ")"}, lines[i+1:]...)...)
			return strings.Join(lines, "\n")
		}
	}

	return code
}

// ValidateGoSyntax checks if the code is valid Go
func ValidateGoSyntax(code string) error {
	fset := token.NewFileSet()
	_, err := parser.ParseFile(fset, "temp.go", code, parser.AllErrors)
	return err
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

		fmt.Printf("📝 Processing: %s\n", filename)

		// ✅ Move tests next to handlers and adjust package name
		filename, content = rules.PlaceTestsWithHandlers(filename, content)
		metrics.RuleFixes++

		// ✅ Fix handler signatures FIRST (missing w, r parameters)
		if strings.Contains(filename, "handlers/") && !strings.HasSuffix(filename, "_test.go") {
			before := content
			var wasFixed bool
			content, wasFixed = FixHandlerSignatures(content)
			if wasFixed {
				metrics.RuleFixes++
			}

			// Apply safety rule for handlers (decode + type mismatch fix)
			content = rules.SafeDecode(content)
			content = rules.FixIDTypeMismatch(content)
			content = removeUnusedModelsImport(content)
			if content != before {
				metrics.RuleFixes++
			}
		}

		// ✅ Fix routes Register function parameter
		if strings.Contains(filename, "routes") && !strings.HasSuffix(filename, "_test.go") {
			before := content
			content = fixRegisterParameter(content)
			if content != before {
				metrics.RuleFixes++
				fmt.Println("🔧 Fixed Register function parameter")
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

			// 🔧 Validate and fix test function signatures FIRST
			content, wasFixed := ValidateAndFixTestFunctions(content)
			if wasFixed {
				metrics.RuleFixes++
				fmt.Printf("✅ Auto-fixed test signatures in %s\n", filename)
			}

			// Then apply other test fixes
			content = rules.FixTestImports(content)
			content = rules.FixTestBodies(content)
			content = CleanDuplicateImports(content)

			// Ensure net/http import for httptest
			if !strings.Contains(content, `"net/http"`) && strings.Contains(content, "httptest") {
				content = ensureNetHTTPImport(content)
				metrics.RuleFixes++
			}

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

		// ✅ Validate syntax before writing
		if err := ValidateGoSyntax(content); err != nil {
			fmt.Printf("⚠️  Syntax validation failed for %s: %v\n", filename, err)
			fmt.Println("🔧 Attempting additional fixes...")

			// Try one more comma fix pass
			content, _ = fixMissingCommas(content)

			// Re-validate
			if err := ValidateGoSyntax(content); err != nil {
				fmt.Printf("❌ Could not auto-fix syntax errors in %s\n", filename)
				// Continue anyway - go build will catch it
			} else {
				fmt.Printf("✅ Auto-fixed syntax errors in %s\n", filename)
				metrics.RuleFixes++
			}
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

		fmt.Printf("✅ Written: %s\n", fullPath)
	}

	fmt.Printf("\n🔧 Total rule-based fixes applied: %d\n", metrics.RuleFixes)
	return nil
}

// FixAllGeneratedFiles applies all fixes to already generated files
func FixAllGeneratedFiles(projectDir string) error {
	fmt.Printf("\n🔧 Auto-fixing all files in: %s\n", projectDir)

	fixes := 0

	// Fix 1: Test function signatures
	fmt.Println("\n📝 Fixing test function signatures...")
	err := filepath.Walk(filepath.Join(projectDir, "internal", "handlers"), func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(path, "_test.go") {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		original := string(content)
		fixed := original

		// Fix: func TestXxx() { -> func TestXxx(t *testing.T) {
		re := regexp.MustCompile(`func (Test\w+)\s*\(\s*\)\s*\{`)
		fixed = re.ReplaceAllString(fixed, `func $1(t *testing.T) {`)

		if fixed != original {
			err = os.WriteFile(path, []byte(fixed), 0644)
			if err == nil {
				fixes++
				fmt.Printf("  ✅ Fixed test signatures in %s\n", filepath.Base(path))
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	// Fix 2: Handler function signatures
	fmt.Println("\n📝 Fixing handler function signatures...")
	err = filepath.Walk(filepath.Join(projectDir, "internal", "handlers"), func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || strings.HasSuffix(path, "_test.go") || !strings.HasSuffix(path, ".go") {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		original := string(content)
		fixed := original

		// Fix: func CreateXxx() { -> func CreateXxx(w http.ResponseWriter, r *http.Request) {
		re := regexp.MustCompile(`func (Create|Get|Update|Delete|List)\w+\s*\(\s*\)\s*\{`)
		fixed = re.ReplaceAllString(fixed, `func $1(w http.ResponseWriter, r *http.Request) {`)

		// Ensure net/http import
		if fixed != original {
			if !strings.Contains(fixed, `"net/http"`) {
				fixed = ensureNetHTTPImport(fixed)
			}
			if !strings.Contains(fixed, `"encoding/json"`) {
				fixed = ensureJSONImport(fixed)
			}

			err = os.WriteFile(path, []byte(fixed), 0644)
			if err == nil {
				fixes++
				fmt.Printf("  ✅ Fixed handler signatures in %s\n", filepath.Base(path))
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	// Fix 3: Routes Register function
	fmt.Println("\n📝 Fixing routes Register function...")
	routesPath := filepath.Join(projectDir, "internal", "routes", "routes.go")
	if _, err := os.Stat(routesPath); err == nil {
		content, err := os.ReadFile(routesPath)
		if err == nil {
			original := string(content)
			fixed := fixRegisterParameter(original)

			if fixed != original {
				err = os.WriteFile(routesPath, []byte(fixed), 0644)
				if err == nil {
					fixes++
					fmt.Println("  ✅ Fixed Register function parameter")
				}
			}
		}
	}

	// Fix 4: Missing commas in test files
	fmt.Println("\n📝 Fixing missing commas in struct literals...")
	err = filepath.Walk(filepath.Join(projectDir, "internal", "handlers"), func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(path, "_test.go") {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		original := string(content)
		fixed, wasFixed := fixMissingCommas(original)

		if wasFixed {
			err = os.WriteFile(path, []byte(fixed), 0644)
			if err == nil {
				fixes++
				fmt.Printf("  ✅ Fixed missing commas in %s\n", filepath.Base(path))
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	// Fix 5: Ensure all test files have required imports
	fmt.Println("\n📝 Ensuring test file imports...")
	err = filepath.Walk(filepath.Join(projectDir, "internal", "handlers"), func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(path, "_test.go") {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		original := string(content)
		fixed := original

		// Ensure required imports
		if !strings.Contains(fixed, `"testing"`) {
			fixed = ensureTestingImport(fixed)
		}
		if strings.Contains(fixed, "httptest.") && !strings.Contains(fixed, `"net/http/httptest"`) {
			fixed = ensureHTTPTestImport(fixed)
		}
		if strings.Contains(fixed, "bytes.") && !strings.Contains(fixed, `"bytes"`) {
			fixed = ensureBytesImport(fixed)
		}

		if fixed != original {
			err = os.WriteFile(path, []byte(fixed), 0644)
			if err == nil {
				fixes++
				fmt.Printf("  ✅ Fixed imports in %s\n", filepath.Base(path))
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	fmt.Printf("\n✅ Applied %d fixes total!\n", fixes)
	return nil
}

// ensureJSONImport adds encoding/json import if missing
func ensureJSONImport(code string) string {
	if strings.Contains(code, `"encoding/json"`) {
		return code
	}

	lines := strings.Split(code, "\n")
	for i, line := range lines {
		if strings.Contains(line, "import (") {
			lines = append(lines[:i+1], append([]string{"\t\"encoding/json\""}, lines[i+1:]...)...)
			return strings.Join(lines, "\n")
		}
	}
	return code
}

// ensureHTTPTestImport adds net/http/httptest import if missing
func ensureHTTPTestImport(code string) string {
	if strings.Contains(code, `"net/http/httptest"`) {
		return code
	}

	lines := strings.Split(code, "\n")
	for i, line := range lines {
		if strings.Contains(line, "import (") {
			lines = append(lines[:i+1], append([]string{"\t\"net/http/httptest\""}, lines[i+1:]...)...)
			return strings.Join(lines, "\n")
		}
	}
	return code
}

// ensureBytesImport adds bytes import if missing
func ensureBytesImport(code string) string {
	if strings.Contains(code, `"bytes"`) {
		return code
	}

	lines := strings.Split(code, "\n")
	for i, line := range lines {
		if strings.Contains(line, "import (") {
			lines = append(lines[:i+1], append([]string{"\t\"bytes\""}, lines[i+1:]...)...)
			return strings.Join(lines, "\n")
		}
	}
	return code
}

// fixRegisterParameter ensures Register function has the correct mux.Router parameter
func fixRegisterParameter(code string) string {
	registerPattern := regexp.MustCompile(`func\s+Register\s*\(\s*\)\s*{`)
	if registerPattern.MatchString(code) {
		code = registerPattern.ReplaceAllString(code, `func Register(r *mux.Router) {`)
	}
	return code
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
