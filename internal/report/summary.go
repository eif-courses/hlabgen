package report

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ExperimentResult represents a single experiment row.
type ExperimentResult struct {
	AppName         string
	BuildSuccess    bool
	TestsPass       bool
	Coverage        float64
	PrimarySuccess  bool
	RepairAttempts  int
	FinalSuccess    bool
	DurationSeconds float64
	RuleFixes       int // NEW: number of rule-based fixes applied
}

// LoadSummaryCSV parses combined summary.csv and metrics.csv into results.
func LoadSummaryCSV(summaryPath, metricsPath string) ([]ExperimentResult, error) {
	summaryFile, err := os.Open(summaryPath)
	if err != nil {
		return nil, fmt.Errorf("open summary: %w", err)
	}
	defer summaryFile.Close()

	summaryReader := csv.NewReader(summaryFile)
	summaryReader.FieldsPerRecord = -1
	rows, _ := summaryReader.ReadAll()

	results := map[string]*ExperimentResult{}

	// --- Parse summary.csv ---
	for i, row := range rows {
		if i == 0 {
			continue // skip header
		}
		if len(row) < 6 {
			continue
		}
		app := strings.TrimSpace(row[0])
		coverage, _ := strconv.ParseFloat(strings.TrimSuffix(row[5], "%"), 64)

		results[app] = &ExperimentResult{
			AppName:      app,
			BuildSuccess: strings.ToLower(row[2]) == "true",
			TestsPass:    strings.ToLower(row[3]) == "true",
			Coverage:     coverage,
		}
	}

	// --- Parse metrics.csv (optional) ---
	if _, err := os.Stat(metricsPath); err == nil {
		data, _ := os.ReadFile(metricsPath)
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			if line == "" {
				continue
			}
			parts := strings.Split(line, ",")
			if len(parts) < 7 { // now expecting 7 columns including RuleFixes
				continue
			}

			// Extract metrics info
			pSuccess := parts[1] == "true"
			repairs, _ := strconv.Atoi(parts[2])
			fSuccess := parts[3] == "true"
			durStr := strings.TrimSuffix(parts[5], "s")
			d, _ := strconv.ParseFloat(durStr, 64)
			ruleFixes, _ := strconv.Atoi(strings.TrimSpace(parts[6]))

			// Apply metrics to all results (approximation)
			for _, r := range results {
				r.PrimarySuccess = pSuccess
				r.RepairAttempts = repairs
				r.FinalSuccess = fSuccess
				r.DurationSeconds = d
				r.RuleFixes = ruleFixes // NEW
			}
		}
	}

	out := []ExperimentResult{}
	for _, r := range results {
		out = append(out, *r)
	}

	return out, nil
}

// GenerateMarkdownReport creates a markdown table for your article.
func GenerateMarkdownReport(results []ExperimentResult, output string) error {
	header := `# Experimental Evaluation Results

| App | Primary Success | Repair Attempts | Rule Fixes | Final Success | Build Success | Tests Pass | Coverage (%) | Duration (s) |
`
	var rows []string
	for _, r := range results {
		rows = append(rows, fmt.Sprintf("| %s | %v | %d | %d | %v | %v | %v | %.1f | %.2f |",
			r.AppName,
			r.PrimarySuccess,
			r.RepairAttempts,
			r.RuleFixes, // NEW
			r.FinalSuccess,
			r.BuildSuccess,
			r.TestsPass,
			r.Coverage,
			r.DurationSeconds,
		))
	}

	content := header + strings.Join(rows, "\n") + "\n"

	if err := os.WriteFile(output, []byte(content), 0o644); err != nil {
		return err
	}
	fmt.Printf("📊 Markdown results saved → %s\n", output)
	return nil
}

// GenerateSummary runs the full pipeline to produce results.md.
func GenerateSummary() error {
	summary := "experiments/logs/summary.csv"
	metrics := "experiments/logs/metrics.csv"
	output := "experiments/logs/results.md"

	results, err := LoadSummaryCSV(summary, metrics)
	if err != nil {
		return err
	}
	return GenerateMarkdownReport(results, output)
}
