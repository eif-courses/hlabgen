package validate

import (
	"bytes"
	"github.com/eif-courses/hlabgen/internal/metrics"
	"os/exec"
	"strconv"
	"strings"
)

func Run(projectPath string) (metrics.Result, error) {
	m := metrics.Result{}

	// go fmt (non-fatal)
	_ = exec.Command("go", "fmt", "./...").Run()

	// go build
	build := exec.Command("go", "build", "./...")
	build.Dir = projectPath
	err := build.Run()
	m.BuildSuccess = (err == nil)

	// go vet (count warnings)
	vet := exec.Command("go", "vet", "./...")
	vet.Dir = projectPath
	var vetOut bytes.Buffer
	vet.Stderr = &vetOut
	vet.Stdout = &vetOut
	_ = vet.Run()
	m.VetWarnings = countLines(vetOut.String())

	// golangci-lint (optional)
	lint := exec.Command("golangci-lint", "run", "--out-format=tab")
	lint.Dir = projectPath
	var lintOut bytes.Buffer
	lint.Stdout = &lintOut
	_ = lint.Run()
	m.LintWarnings = countLines(lintOut.String())

	// go test -cover
	test := exec.Command("go", "test", "./...", "-cover")
	test.Dir = projectPath
	var testOut bytes.Buffer
	test.Stdout = &testOut
	_ = test.Run()
	out := testOut.String()
	m.TestsPass = strings.Contains(out, "ok\t")
	m.CoveragePct = parseCoverage(out)

	// gocyclo average (if installed)
	cyclo := exec.Command("gocyclo", "-over", "0", ".")
	cyclo.Dir = projectPath
	var cycloOut bytes.Buffer
	cyclo.Stdout = &cycloOut
	_ = cyclo.Run()
	m.CyclomaticAvg = avgCyclo(cycloOut.String())

	return m, nil
}

func countLines(s string) int {
	if strings.TrimSpace(s) == "" {
		return 0
	}
	return len(strings.Split(strings.TrimSpace(s), "\n"))
}
func parseCoverage(out string) float64 {
	// find last 'coverage: XX.X% of statements'
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
	// gocyclo prints lines like: "  5 some/file.go:funcName"
	// very simple average by first field if numeric
	var sum float64
	var n float64
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
