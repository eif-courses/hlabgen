package validate

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/eif-courses/hlabgen/internal/rules"
)

// =============================================================================
// SYNTAX FIXING FUNCTIONS (delegates to rules package where possible)
// =============================================================================

// FixTestFunctions ensures all test functions have correct signature: (t *testing.T)
func FixTestFunctions(code string) (string, bool) {
	fixed := false
	lines := strings.Split(code, "\n")

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if !strings.HasPrefix(trimmed, "func Test") {
			continue
		}

		// Extract function name
		funcNameMatch := regexp.MustCompile(`func (Test\w+)\s*\(`).FindStringSubmatch(trimmed)
		if len(funcNameMatch) < 2 {
			continue
		}
		funcName := funcNameMatch[1]

		// Skip if already correct
		if strings.Contains(trimmed, "func "+funcName+"(t *testing.T)") {
			continue
		}

		// Replace with correct signature
		oldLine := line
		pattern := regexp.MustCompile(`func ` + regexp.QuoteMeta(funcName) + `\s*\([^)]*\)\s*\{`)
		lines[i] = pattern.ReplaceAllString(line, `func `+funcName+`(t *testing.T) {`)

		if lines[i] != oldLine {
			fixed = true
			fmt.Printf("🔧 Fixed test signature for %s\n", funcName)
		}
	}

	return strings.Join(lines, "\n"), fixed
}

// FixHandlerSignatures ensures handlers have (w http.ResponseWriter, r *http.Request)
func FixHandlerSignatures(code string) (string, bool) {
	fixed := false
	lines := strings.Split(code, "\n")

	pattern := regexp.MustCompile(`^func\s+(Create|Get|Update|Delete|List)\w+\(\s*\)\s*{`)

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)

		if pattern.MatchString(trimmed) {
			match := regexp.MustCompile(`func\s+(\w+)`).FindStringSubmatch(trimmed)
			if len(match) > 1 {
				funcName := match[1]
				oldLine := line
				newLine := regexp.MustCompile(`func\s+`+regexp.QuoteMeta(funcName)+`\s*\(\s*\)\s*{`).
					ReplaceAllString(line, `func `+funcName+`(w http.ResponseWriter, r *http.Request) {`)

				if newLine != oldLine {
					lines[i] = newLine
					fixed = true
					fmt.Printf("🔧 Fixed handler signature for %s\n", funcName)
				}
			}
		}
	}

	return strings.Join(lines, "\n"), fixed
}

// FixMissingCommas adds trailing commas in struct literals
func FixMissingCommas(code string) (string, bool) {
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

		// Detect struct literal start
		if (strings.Contains(trimmed, "= models.") ||
			strings.Contains(trimmed, ":= models.") ||
			strings.Contains(trimmed, "= []models.")) &&
			strings.HasSuffix(trimmed, "{") {
			inStructLiteral = true
			structBraceCount = 1
			continue
		}

		// Track brace depth
		if inStructLiteral {
			structBraceCount += strings.Count(trimmed, "{")
			structBraceCount -= strings.Count(trimmed, "}")

			if structBraceCount == 0 {
				inStructLiteral = false
				continue
			}

			// Only fix commas INSIDE struct literals
			if structBraceCount > 0 {
				// Skip if already has comma
				if strings.HasSuffix(trimmed, ",") || strings.HasSuffix(trimmed, "{") {
					continue
				}

				// Check if next line is closing brace or another field
				if (nextLineTrimmed == "}" || nextLineTrimmed == "}," ||
					strings.Contains(nextLineTrimmed, ":")) &&
					strings.Contains(trimmed, ":") {
					lines[i] = currentLine + ","
					fixed = true
					fmt.Printf("🔧 Fixed missing comma at line %d\n", i+1)
				}
			}
		}
	}

	if fixed {
		fmt.Println("✅ Fixed missing commas in struct literals")
	}

	return strings.Join(lines, "\n"), fixed
}

// =============================================================================
// WRAPPERS TO RULES PACKAGE (avoid duplication)
// =============================================================================

// FixIDTypeMismatch delegates to rules.FixIDTypeMismatch
func FixIDTypeMismatch(code string) string {
	return rules.FixIDTypeMismatch(code)
}

// FixRegisterParameter delegates to rules.FixRegisterParameter
func FixRegisterParameter(code string) string {
	return rules.FixRegisterParameter(code)
}

// FixRegisterFunction delegates to rules.FixRegisterFunction
func FixRegisterFunction(code string) string {
	return rules.FixRegisterFunction(code)
}

// =============================================================================
// IMPORT MANAGEMENT (unique to validate package)
// =============================================================================

// RemoveUnusedModelsImport removes "/internal/models" import if it's not used
func RemoveUnusedModelsImport(code string) string {
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

// CleanUnusedImportsInModels removes unused imports in model files
func CleanUnusedImportsInModels(code string) string {
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

// EnsureTimeImport adds "time" import if time.Time is used but not imported
func EnsureTimeImport(code string) string {
	if strings.Contains(code, "time.Time") && !strings.Contains(code, `"time"`) {
		lines := strings.Split(code, "\n")
		var result []string
		added := false
		for i, line := range lines {
			result = append(result, line)
			if !added && strings.Contains(line, "import (") {
				if i+1 < len(lines) {
					result = append(result, "\t\"time\"")
					added = true
				}
			}
		}
		if added {
			code = strings.Join(result, "\n")
			fmt.Println("🔧 Added missing time import")
		}
	}
	return code
}

// RemoveUnusedImports removes common unused imports from code
func RemoveUnusedImports(code string) string {
	lines := strings.Split(code, "\n")
	var result []string
	inImportBlock := false

	for _, line := range lines {
		trim := strings.TrimSpace(line)

		if strings.HasPrefix(trim, "import (") {
			inImportBlock = true
			result = append(result, line)
			continue
		}

		if inImportBlock && trim == ")" {
			inImportBlock = false
			result = append(result, line)
			continue
		}

		if inImportBlock && trim != "" && !strings.HasPrefix(trim, "//") {
			// Never skip handlers import in test files
			if strings.Contains(trim, "/internal/handlers\"") {
				result = append(result, line)
				continue
			}

			shouldSkip := false

			// Skip unused bytes
			if strings.Contains(trim, `"bytes"`) && !strings.Contains(code, "bytes.") {
				shouldSkip = true
			}

			// Skip unused encoding/json
			if strings.Contains(trim, `"encoding/json"`) && !strings.Contains(code, "json.") {
				shouldSkip = true
			}

			// Skip unused models
			if strings.Contains(trim, "/internal/models\"") && !strings.Contains(code, "models.") {
				shouldSkip = true
			}

			if !shouldSkip {
				result = append(result, line)
			}
			continue
		}

		result = append(result, line)
	}

	return strings.Join(result, "\n")
}

// =============================================================================
// CONVENIENCE FUNCTIONS
// =============================================================================

// ValidateAndFixAllSyntax runs all syntax fixes on code
func ValidateAndFixAllSyntax(code string, fileType string) (string, int) {
	fixCount := 0

	switch fileType {
	case "handler":
		var wasFixed bool
		code, wasFixed = FixHandlerSignatures(code)
		if wasFixed {
			fixCount++
		}
		code = FixIDTypeMismatch(code)

	case "test":
		var wasFixed bool
		code, wasFixed = FixTestFunctions(code)
		if wasFixed {
			fixCount++
		}
		code, wasFixed = FixMissingCommas(code)
		if wasFixed {
			fixCount++
		}

	case "routes":
		oldCode := code
		code = FixRegisterParameter(code)
		code = FixRegisterFunction(code)
		if code != oldCode {
			fixCount++
		}
	}

	return code, fixCount
}

// FixAllSyntaxIssues applies all available syntax fixes
func FixAllSyntaxIssues(code string) (string, int) {
	totalFixes := 0

	// Fix test functions
	var wasFixed bool
	code, wasFixed = FixTestFunctions(code)
	if wasFixed {
		totalFixes++
	}

	// Fix handler signatures
	code, wasFixed = FixHandlerSignatures(code)
	if wasFixed {
		totalFixes++
	}

	// Fix missing commas
	code, wasFixed = FixMissingCommas(code)
	if wasFixed {
		totalFixes++
	}

	// Fix ID type mismatch (delegates to rules package)
	oldCode := code
	code = FixIDTypeMismatch(code)
	if code != oldCode {
		totalFixes++
	}

	// Fix Register parameter (delegates to rules package)
	oldCode = code
	code = FixRegisterParameter(code)
	if code != oldCode {
		totalFixes++
	}

	return code, totalFixes
}
