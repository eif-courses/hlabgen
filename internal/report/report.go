package report

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// ExperimentResult represents combined metrics for one experiment.
type ExperimentResult struct {
	AppName         string  `json:"app_name"`
	PrimarySuccess  bool    `json:"primary_success"`
	RepairAttempts  int     `json:"repair_attempts"`
	FinalSuccess    bool    `json:"final_success"`
	ErrorMessage    string  `json:"error_message"`
	DurationSeconds float64 `json:"duration_seconds"`
	RuleFixes       int     `json:"rule_fixes"`
	RelaxedMode     bool    `json:"relaxed_mode"`
}

// GenerationMetrics mirrors your ml.GenerationMetrics struct.
type GenerationMetrics struct {
	StartTime      time.Time `json:"start_time"`
	EndTime        time.Time `json:"end_time"`
	Duration       string    `json:"duration"`
	PrimarySuccess bool      `json:"primary_success"`
	RepairAttempts int       `json:"repair_attempts"`
	FinalSuccess   bool      `json:"final_success"`
	ErrorMessage   string    `json:"error_message"`
	RuleFixes      int       `json:"rule_fixes"`
}

// LoadMetricsFromJSON loads a single experiment’s metrics JSON.
func LoadMetricsFromJSON(path string, relaxed bool) (ExperimentResult, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return ExperimentResult{}, err
	}

	var m GenerationMetrics
	if err := json.Unmarshal(data, &m); err != nil {
		return ExperimentResult{}, err
	}

	// Convert duration to seconds
	duration, _ := time.ParseDuration(m.Duration)

	app := filepath.Base(filepath.Dir(path))
	return ExperimentResult{
		AppName:         app,
		PrimarySuccess:  m.PrimarySuccess,
		RepairAttempts:  m.RepairAttempts,
		FinalSuccess:    m.FinalSuccess,
		ErrorMessage:    m.ErrorMessage,
		DurationSeconds: duration.Seconds(),
		RuleFixes:       m.RuleFixes,
		RelaxedMode:     relaxed,
	}, nil
}

// CollectAllExperiments scans experiments/* for metrics files.
func CollectAllExperiments(baseDir string) ([]ExperimentResult, error) {
	var results []ExperimentResult

	entries, err := os.ReadDir(baseDir)
	if err != nil {
		return nil, err
	}

	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		appDir := filepath.Join(baseDir, e.Name())

		mainPath := filepath.Join(appDir, "gen_metrics.json")
		if _, err := os.Stat(mainPath); err == nil {
			res, err := LoadMetricsFromJSON(mainPath, false)
			if err == nil {
				results = append(results, res)
			}
		}

		relaxedPath := filepath.Join(appDir, "gen_metrics_relaxed.json")
		if _, err := os.Stat(relaxedPath); err == nil {
			res, err := LoadMetricsFromJSON(relaxedPath, true)
			if err == nil {
				results = append(results, res)
			}
		}
	}

	return results, nil
}

// GenerateMarkdownReport writes a summary Markdown table.
func GenerateMarkdownReport(results []ExperimentResult, output string) error {
	header := `# Experimental Evaluation Results

| App | Mode | Primary Success | Repair Attempts | Rule Fixes | Final Success | Duration (s) | Error |
|-----|------|----------------|----------------|-------------|----------------|---------------|-------|
`

	var rows []string
	for _, r := range results {
		mode := "Normal"
		if r.RelaxedMode {
			mode = "Relaxed"
		}
		rows = append(rows, fmt.Sprintf(
			"| %s | %s | %v | %d | %d | %v | %.2f | %s |",
			r.AppName,
			mode,
			r.PrimarySuccess,
			r.RepairAttempts,
			r.RuleFixes,
			r.FinalSuccess,
			r.DurationSeconds,
			shorten(r.ErrorMessage, 50),
		))
	}

	content := header + strings.Join(rows, "\n") + "\n"

	os.MkdirAll(filepath.Dir(output), 0o755)
	if err := os.WriteFile(output, []byte(content), 0o644); err != nil {
		return err
	}

	fmt.Printf("📊 Markdown results saved → %s\n", output)
	return nil
}

// GenerateSummaryJSONReport scans all experiments and produces results.md.
func GenerateSummaryJSONReport() error {
	baseDir := "experiments"
	output := filepath.Join(baseDir, "logs", "results.md")

	results, err := CollectAllExperiments(baseDir)
	if err != nil {
		return fmt.Errorf("collect experiments: %w", err)
	}

	if len(results) == 0 {
		return fmt.Errorf("no gen_metrics.json files found under %s", baseDir)
	}

	return GenerateMarkdownReport(results, output)
}

// shorten truncates error messages for cleaner tables.
func shorten(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max-3] + "..."
}
