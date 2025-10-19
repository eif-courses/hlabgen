package metrics

// Result captures all build, lint, and test metrics for a single experiment run.
type Result struct {
	BuildSuccess  bool    `json:"build_success"`  // Whether 'go build ./...' succeeded
	VetWarnings   int     `json:"vet_warnings"`   // Number of 'go vet' warnings
	LintWarnings  int     `json:"lint_warnings"`  // Number of 'golangci-lint' warnings
	TestsPass     bool    `json:"tests_pass"`     // True if all tests passed
	CoveragePct   float64 `json:"coverage_pct"`   // Code coverage percentage
	CyclomaticAvg float64 `json:"cyclomatic_avg"` // Average cyclomatic complexity per function
	GenTimeSec    float64 `json:"gen_time_sec"`   // Total validation duration in seconds
	Timestamp     string  `json:"timestamp"`      // When the validation was completed (optional)
}
