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
	var out []string
	inImportBlock := false
	seenImports := make(map[string]bool)

	for _, line := range lines {
		trim := strings.TrimSpace(line)

		if strings.HasPrefix(trim, "import (") {
			inImportBlock = true
			out = append(out, line)
			continue
		}

		if inImportBlock && trim == ")" {
			inImportBlock = false
			out = append(out, line)
			continue
		}

		if inImportBlock {
			if strings.Contains(trim, `"`) {
				start := strings.Index(trim, `"`)
				end := strings.LastIndex(trim, `"`)
				if start >= 0 && end > start {
					importPath := trim[start : end+1]
					if seenImports[importPath] {
						continue
					}
					seenImports[importPath] = true
				}
			}
		}

		out = append(out, line)
	}

	code = strings.Join(out, "\n")

	openCount := strings.Count(code, "{")
	closeCount := strings.Count(code, "}")
	if openCount > closeCount {
		code += "\n}"
	}

	return code
}

// cleanUnusedImportsInModels removes unused imports from model files
func cleanUnusedImportsInModels(code string) string {
	if strings.Contains(code, `"encoding/json"`) && !strings.Contains(code, "json.") {
		lines := strings.Split(code, "\n")
		var result []string
		for _, line := range lines {
			if strings.Contains(line, `"encoding/json"`) {
				continue
			}
			result = append(result, line)
		}
		code = strings.Join(result, "\n")
	}
	return code
}

// ensureTimeImport adds time import if time.Time is used but not imported
func ensureTimeImport(code string) string {
	if !strings.Contains(code, "time.Time") {
		return code
	}
	if strings.Contains(code, `"time"`) {
		return code
	}

	lines := strings.Split(code, "\n")
	var result []string
	importAdded := false

	for i, line := range lines {
		result = append(result, line)
		if !importAdded && strings.HasPrefix(strings.TrimSpace(line), "package ") {
			if i+1 < len(lines) && strings.Contains(lines[i+1], "import") {
				continue
			}
			result = append(result, "")
			result = append(result, `import "time"`)
			importAdded = true
		}
	}

	return strings.Join(result, "\n")
}

// removeUnusedModelsImport removes models import if not used in handlers
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
