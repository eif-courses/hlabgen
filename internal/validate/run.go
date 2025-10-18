package validate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/eif-courses/hlabgen/internal/metrics"
)

// Run performs validation and metrics extraction for a generated Go project.
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
		perPkg := make(map[string]float64)

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
				perPkg[filepath.Base(dir)] = cov
			}
		}

		m.TestsPass = passed
		m.CoveragePct = average(coverValues)

		// ✅ Save per-package coverage as JSON
		covPath := filepath.Join(projectPath, "coverage.json")
		if data, err := json.MarshalIndent(perPkg, "", "  "); err == nil {
			_ = os.WriteFile(covPath, data, 0o644)
			fmt.Printf("📁 Saved per-package coverage → %s\n", covPath)
		}

		// ✅ Append global coverage summary to CSV (with ML metrics)
		_ = appendCoverageCSV(projectPath, m)

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

// appendCoverageCSV appends project metrics + ML metrics to experiments/logs/coverage.csv.
func appendCoverageCSV(projectPath string, m metrics.Result) error {
	appName := filepath.Base(projectPath)
	logDir := filepath.Join("experiments", "logs")
	_ = os.MkdirAll(logDir, 0o755)
	filePath := filepath.Join(logDir, "coverage.csv")

	// Try reading gen_metrics.json if available
	type genMetrics struct {
		Duration       float64 `json:"duration_sec"`
		RepairAttempts int     `json:"repair_attempts"`
	}
	var g genMetrics
	genFile := filepath.Join(projectPath, "gen_metrics.json")
	if data, err := os.ReadFile(genFile); err == nil {
		_ = json.Unmarshal(data, &g)
	}

	// If new file, write header
	_, err := os.Stat(filePath)
	isNew := os.IsNotExist(err)
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()

	if isNew {
		header := "AppName,BuildSuccess,TestsPass,CoveragePct,MLDurationSec,RepairAttempts\n"
		if _, err := f.WriteString(header); err != nil {
			return err
		}
	}

	row := fmt.Sprintf("%s,%v,%v,%.1f,%.2f,%d\n",
		appName,
		m.BuildSuccess,
		m.TestsPass,
		m.CoveragePct,
		g.Duration,
		g.RepairAttempts,
	)
	_, _ = f.WriteString(row)
	fmt.Printf("🧾 Added summary row (with ML metrics) → experiments/logs/coverage.csv\n")
	return nil
}

// --- Helpers ---

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
