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
		for _, dir := range testDirs {
			cmd := exec.Command("go", "test", "-cover")
			cmd.Dir = dir
			var out bytes.Buffer
			cmd.Stdout, cmd.Stderr = &out, &out
			_ = cmd.Run()
			combinedOut.WriteString(out.String() + "\n")
		}

		out := combinedOut.String()
		m.TestsPass = strings.Contains(out, "PASS")
		m.CoveragePct = parseCoverage(out)
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
		if err != nil || info.IsDir() {
			return nil
		}
		if strings.HasSuffix(info.Name(), "_test.go") {
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

func countLines(s string) int {
	if strings.TrimSpace(s) == "" {
		return 0
	}
	return len(strings.Split(strings.TrimSpace(s), "\n"))
}

func parseCoverage(out string) float64 {
	for _, line := range strings.Split(out, "\n") {
		if strings.Contains(line, "coverage:") && strings.Contains(line, "%") {
			f := strings.Fields(line)
			for _, t := range f {
				if strings.HasSuffix(t, "%") {
					val := strings.TrimSuffix(t, "%")
					if v, err := strconv.ParseFloat(val, 64); err == nil {
						return v
					}
				}
			}
		}
	}
	return 0
}

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
