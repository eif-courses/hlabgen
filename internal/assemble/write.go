package assemble

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/eif-courses/hlabgen/internal/rules"
)

// File represents a generated file with path and content.
type File struct {
	Filename string
	Content  string
}

// WriteMany writes multiple generated files to disk,
// applies rule-based safety fixes, and auto-fixes import paths.
func WriteMany(base string, files []File) error {
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

		// ✅ Apply safety rule for handlers (decode + type mismatch fix)
		if strings.Contains(filename, "handlers/") {
			content = rules.SafeDecode(content)
			content = rules.FixIDTypeMismatch(content)
		}

		// ✅ Normalize routes: rename RegisterRoutes → Register
		if strings.Contains(filename, "routes.go") {
			content = rules.FixRegisterFunction(content)
		}

		// ✅ Apply test fixes (imports + JSON body + cleanup)
		if strings.HasSuffix(filename, "_test.go") {
			content = rules.FixTestImports(content)
			content = rules.FixTestBodies(content)
			content = CleanDuplicateImports(content)
			content = FixUnbalancedBraces(content)
		}

		// ✅ Remove unnecessary mux imports in handlers
		if strings.Contains(filename, "handlers/") && strings.Contains(content, `"github.com/gorilla/mux"`) {
			content = strings.ReplaceAll(content, "\t\"github.com/gorilla/mux\"\n", "")
		}

		// ✅ Auto-fix placeholder import paths like "yourapp/", "your_project/", etc.
		if moduleName != "" {
			content = strings.ReplaceAll(content, `"yourapp/routes"`, fmt.Sprintf(`"%s/internal/routes"`, moduleName))
			content = strings.ReplaceAll(content, `"yourapp/handlers"`, fmt.Sprintf(`"%s/internal/handlers"`, moduleName))
			content = strings.ReplaceAll(content, `"yourapp/models"`, fmt.Sprintf(`"%s/internal/models"`, moduleName))

			content = strings.ReplaceAll(content, `"your_project/routes"`, fmt.Sprintf(`"%s/internal/routes"`, moduleName))
			content = strings.ReplaceAll(content, `"your_project/handlers"`, fmt.Sprintf(`"%s/internal/handlers"`, moduleName))
			content = strings.ReplaceAll(content, `"your_project/models"`, fmt.Sprintf(`"%s/internal/models"`, moduleName))

			// Generic fallback regex for any other "your..." placeholder
			re := regexp.MustCompile(`"your[^"]+/`)
			content = re.ReplaceAllString(content, fmt.Sprintf(`"%s/internal/`, moduleName))
		}

		// ✅ Final deduplication pass
		content = CleanDuplicateImports(content)

		// ✅ Normalize output paths (ensure /internal/ structure)
		fullPath := filepath.Join(base, rules.NormalizePath(filename))
		if err := os.MkdirAll(filepath.Dir(fullPath), 0o755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", filepath.Dir(fullPath), err)
		}

		// ✅ Write file to disk
		if err := os.WriteFile(fullPath, []byte(content), 0o644); err != nil {
			return fmt.Errorf("failed to write file %s: %w", fullPath, err)
		}

		// ✅ Log fixes for visibility
		if strings.Contains(f.Content, "yourapp/") || strings.Contains(f.Content, "your_project/") {
			fmt.Printf("🔧 Fixed imports in: %s\n", filename)
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
	var out []string
	importBlock := false

	for _, line := range lines {
		trim := strings.TrimSpace(line)

		// Detect start/end of import block
		if strings.HasPrefix(trim, "import (") {
			importBlock = true
			out = append(out, line)
			continue
		}
		if importBlock && strings.HasPrefix(trim, ")") {
			importBlock = false
			out = append(out, line)
			continue
		}

		// Skip duplicate imports (normalize spacing)
		if importBlock && strings.HasPrefix(trim, `"`) {
			key := strings.ReplaceAll(strings.ReplaceAll(trim, "\t", ""), " ", "")
			if seen[key] {
				continue
			}
			seen[key] = true
		}

		out = append(out, line)
	}

	code = strings.Join(out, "\n")

	// ✅ Ensure exactly one closing brace at EOF if braces are unbalanced
	openCount := strings.Count(code, "{")
	closeCount := strings.Count(code, "}")
	if openCount > closeCount {
		code += "\n}"
	}

	return code
}

// FixUnbalancedBraces ensures generated Go files end with properly balanced braces.
func FixUnbalancedBraces(code string) string {
	openCount := strings.Count(code, "{")
	closeCount := strings.Count(code, "}")
	diff := openCount - closeCount

	if diff > 0 {
		code += strings.Repeat("\n}", diff)
	}
	return code
}
