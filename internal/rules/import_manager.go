package rules

import (
	"fmt"
	"strings"
)

// EnsureImport adds an import to Go code if missing
// Handles both single-line and multi-line import blocks
func EnsureImport(code, importPath string) string {
	if strings.Contains(code, `"`+importPath+`"`) {
		return code // Already imported
	}

	lines := strings.Split(code, "\n")
	var result []string
	added := false

	for i, line := range lines {
		result = append(result, line)

		// Add to existing import block
		if !added && strings.Contains(line, "import (") {
			result = append(result, "\t\""+importPath+"\"")
			added = true
		}
	}

	// If no import block found, create one after package declaration
	if !added {
		for i, line := range result {
			if strings.HasPrefix(strings.TrimSpace(line), "package ") {
				newResult := make([]string, 0, len(result)+4)
				newResult = append(newResult, result[:i+1]...)
				newResult = append(newResult, "", "import (", "\t\""+importPath+"\"", ")")
				newResult = append(newResult, result[i+1:]...)
				return strings.Join(newResult, "\n")
			}
		}
	}

	return strings.Join(result, "\n")
}

// EnsureNetHTTPImport adds "net/http" if used but missing
func EnsureNetHTTPImport(code string) string {
	return EnsureImport(code, "net/http")
}

// EnsureTestingImport adds "testing" if used but missing
func EnsureTestingImport(code string) string {
	return EnsureImport(code, "testing")
}

// EnsureJSONImport adds "encoding/json" if used but missing
func EnsureJSONImport(code string) string {
	return EnsureImport(code, "encoding/json")
}

// EnsureHTTPTestImport adds "net/http/httptest" if used but missing
func EnsureHTTPTestImport(code string) string {
	return EnsureImport(code, "net/http/httptest")
}

// EnsureBytesImport adds "bytes" if used but missing
func EnsureBytesImport(code string) string {
	return EnsureImport(code, "bytes")
}

// EnsureMuxImport adds "github.com/gorilla/mux" if used but missing
func EnsureMuxImport(code string) string {
	return EnsureImport(code, "github.com/gorilla/mux")
}

// EnsureModelsImport adds module-specific models import
func EnsureModelsImport(code, moduleName string) string {
	return EnsureImport(code, moduleName+"/internal/models")
}

// EnsureHandlersImport adds module-specific handlers import
func EnsureHandlersImport(code, moduleName string) string {
	return EnsureImport(code, moduleName+"/internal/handlers")
}

// EnsureStrconvImport adds "strconv" if used but missing
func EnsureStrconvImport(code string) string {
	return EnsureImport(code, "strconv")
}

// EnsureTimeImport adds "time" if used but missing
func EnsureTimeImport(code string) string {
	return EnsureImport(code, "time")
}

// EnsureContextImport adds "context" if used but missing
func EnsureContextImport(code string) string {
	return EnsureImport(code, "context")
}

// EnsureStringsImport adds "strings" if used but missing
func EnsureStringsImport(code string) string {
	return EnsureImport(code, "strings")
}

// CleanDuplicateImports removes duplicate import lines
func CleanDuplicateImports(code string) string {
	lines := strings.Split(code, "\n")
	var out []string
	inImportBlock := false
	seenImports := make(map[string]bool)

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		trim := strings.TrimSpace(line)

		// Detect import block start
		if strings.HasPrefix(trim, "import (") {
			inImportBlock = true
			seenImports = make(map[string]bool)
			out = append(out, line)
			continue
		}

		// Detect import block end
		if inImportBlock && trim == ")" {
			inImportBlock = false
			out = append(out, line)
			continue
		}

		// Process imports inside the block
		if inImportBlock {
			if trim == "" {
				continue
			}

			importKey := strings.TrimSpace(trim)

			if seenImports[importKey] {
				continue
			}
			seenImports[importKey] = true
			out = append(out, line)
			continue
		}

		out = append(out, line)
	}

	code = strings.Join(out, "\n")

	// Ensure balanced braces
	openCount := strings.Count(code, "{")
	closeCount := strings.Count(code, "}")
	if openCount > closeCount {
		code += strings.Repeat("\n}", openCount-closeCount)
	}

	return code
}

// RemoveUnusedImport removes a specific import if it's not used in the code
func RemoveUnusedImport(code, importPath string) string {
	// Extract package name from import path
	parts := strings.Split(importPath, "/")
	pkgName := parts[len(parts)-1]

	// Check if package is actually used
	if !strings.Contains(code, pkgName+".") {
		lines := strings.Split(code, "\n")
		var result []string

		for _, line := range lines {
			if strings.Contains(line, `"`+importPath+`"`) {
				continue // Skip this import
			}
			result = append(result, line)
		}

		return strings.Join(result, "\n")
	}

	return code
}

// EnsureImportsForHandlers ensures all required imports for handler files
func EnsureImportsForHandlers(code, moduleName string) string {
	code = EnsureNetHTTPImport(code)
	code = EnsureJSONImport(code)
	code = EnsureStrconvImport(code)
	code = EnsureMuxImport(code)
	code = EnsureModelsImport(code, moduleName)
	return CleanDuplicateImports(code)
}

// EnsureImportsForTests ensures all required imports for test files
func EnsureImportsForTests(code, moduleName string) string {
	code = EnsureTestingImport(code)
	code = EnsureNetHTTPImport(code)
	code = EnsureHTTPTestImport(code)
	code = EnsureBytesImport(code)
	code = EnsureJSONImport(code)
	code = EnsureMuxImport(code)
	code = EnsureHandlersImport(code, moduleName)
	code = EnsureModelsImport(code, moduleName)
	return CleanDuplicateImports(code)
}

// EnsureImportsForRoutes ensures all required imports for routes files
func EnsureImportsForRoutes(code, moduleName string) string {
	code = EnsureMuxImport(code)
	code = EnsureHandlersImport(code, moduleName)
	return CleanDuplicateImports(code)
}

// FixImportPaths replaces placeholder import paths with correct module paths
func FixImportPaths(code, moduleName string) string {
	// Fix common placeholder patterns
	code = strings.ReplaceAll(code, `"yourapp/`, fmt.Sprintf(`"%s/`, moduleName))
	code = strings.ReplaceAll(code, `"your_project/`, fmt.Sprintf(`"%s/`, moduleName))
	code = strings.ReplaceAll(code, `"github.com/yourusername/`, fmt.Sprintf(`"%s/`, moduleName))

	return code
}
