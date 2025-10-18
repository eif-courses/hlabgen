package validate

import (
	"bytes"
	"fmt"
	"github.com/eif-courses/hlabgen/internal/metrics"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func Run(projectPath string) (metrics.Result, error) {
	m := metrics.Result{}

	// --- go fmt (non-fatal)
	_ = exec.Command("go", "fmt", "./...").Run()

	// --- go build
	build := exec.Command("go", "build", "./...")
	build.Dir = projectPath
	err := build.Run()
	m.BuildSuccess = (err == nil)

	// --- go vet
	vet := exec.Command("go", "vet", "./...")
	vet.Dir = projectPath
	var vetOut bytes.Buffer
	vet.Stdout, vet.Stderr = &vetOut, &vetOut
	_ = vet.Run()
	m.VetWarnings = countLines(vetOut.String())

	// --- golangci-lint (optional)
	lint := exec.Command("golangci-lint", "run", "--out-format=tab")
	lint.Dir = projectPath
	var lintOut bytes.Buffer
	lint.Stdout = &lintOut
	_ = lint.Run()
	m.LintWarnings = countLines(lintOut.String())

	// --- go test -cover (smart detection)
	testDirs := findTestDirs(projectPath)
	if len(testDirs) == 0 {
		fmt.Println("⚠️  No test files detected — skipping tests.")
		m.TestsPass = false
		m.CoveragePct = 0
	} else {
		var combinedOut strings.Builder
		var coverValues []float64
		passed := false

		for _, dir := range testDirs {
			cmd := exec.Command("go", "test", "-v", "-cover")
			cmd.Dir = dir
			var out bytes.Buffer
			cmd.Stdout, cmd.Stderr = &out, &out
			_ = cmd.Run()
			output := out.String()
			combinedOut.WriteString(fmt.Sprintf("\n--- %s ---\n%s", dir, output))

			if strings.Contains(output, "PASS") {
				passed = true
			}
			if cov := parseCoverage(output); cov > 0 {
				coverValues = append(coverValues, cov)
			}
		}

		// ✅ Mark as passed if *any* directory reports PASS
		m.TestsPass = passed

		// ✅ Average coverage across testable packages
		m.CoveragePct = average(coverValues)

		// Optional debug print
		fmt.Println("\n--- go test summary ---")
		fmt.Println(combinedOut.String())
		fmt.Println("------------------------")
	}

	// --- gocyclo (if available)
	cyclo := exec.Command("gocyclo", "-over", "0", ".")
	cyclo.Dir = projectPath
	var cycloOut bytes.Buffer
	cyclo.Stdout = &cycloOut
	_ = cyclo.Run()
	m.CyclomaticAvg = avgCyclo(cycloOut.String())

	return m, nil
}

// findTestDirs scans for directories containing *_test.go files.
func findTestDirs(root string) []string {
	var testDirs []string
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), "_test.go") {
			dir := filepath.Dir(path)
			if !contains(testDirs, dir) {
				testDirs = append(testDirs, dir)
			}
		}
		return nil
	})
	return testDirs
}

// Helper to check slice membership.
func contains(list []string, val string) bool {
	for _, v := range list {
		if v == val {
			return true
		}
	}
	return false
}

// Counts non-empty lines (used for vet/lint warnings).
func countLines(s string) int {
	if strings.TrimSpace(s) == "" {
		return 0
	}
	return len(strings.Split(strings.TrimSpace(s), "\n"))
}

// Extracts coverage percentage (e.g. "coverage: 38.5% of statements").
func parseCoverage(out string) float64 {
	for _, line := range strings.Split(out, "\n") {
		if strings.Contains(line, "coverage:") && strings.Contains(line, "%") {
			fields := strings.Fields(line)
			for _, f := range fields {
				if strings.HasSuffix(f, "%") {
					val := strings.TrimSuffix(f, "%")
					if v, err := strconv.ParseFloat(val, 64); err == nil {
						return v
					}
				}
			}
		}
	}
	return 0
}

// Calculates the average of float slice values.
func average(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	var sum float64
	for _, v := range values {
		sum += v
	}
	return sum / float64(len(values))
}

// Computes average cyclomatic complexity (if gocyclo is available).
func avgCyclo(out string) float64 {
	var sum, n float64
	for _, line := range strings.Split(strings.TrimSpace(out), "\n") {
		fs := strings.Fields(line)
		if len(fs) > 0 {
			if v, err := strconv.ParseFloat(fs[0], 64); err == nil {
				sum += v
				n++
			}
		}
	}
	if n == 0 {
		return 0
	}
	return sum / n
}
