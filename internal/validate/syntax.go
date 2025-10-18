package validate

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

// ValidateGoSyntax checks all .go files for syntax errors
func ValidateGoSyntax(projectPath string) []string {
	var errors []string

	filepath.Walk(projectPath, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}

		fset := token.NewFileSet()
		_, parseErr := parser.ParseFile(fset, path, nil, parser.AllErrors)

		if parseErr != nil {
			relPath, _ := filepath.Rel(projectPath, path)
			errors = append(errors, fmt.Sprintf("%s: %v", relPath, parseErr))
		}

		return nil
	})

	return errors
}
