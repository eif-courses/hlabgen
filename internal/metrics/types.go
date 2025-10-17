package metrics

type Result struct {
	BuildSuccess  bool    `json:"build_success"`
	VetWarnings   int     `json:"vet_warnings"`
	LintWarnings  int     `json:"lint_warnings"`
	TestsPass     bool    `json:"tests_pass"`
	CoveragePct   float64 `json:"coverage_pct"`
	CyclomaticAvg float64 `json:"cyclomatic_avg"`
	GenTimeSec    float64 `json:"gen_time_sec"`
}
