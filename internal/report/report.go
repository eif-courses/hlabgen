package report

import (
	"encoding/csv"
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
	Mode            string  `json:"mode"`
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

// Global cache for CSV data
var csvModeCache map[string]string

// LoadMetricsFromJSON loads a single experiment's metrics JSON.
func LoadMetricsFromJSON(path string) (ExperimentResult, error) {
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

	// ✅ Try to read mode from experiment_info.txt first, then fall back to CSV
	mode := readModeFromExperimentInfo(filepath.Dir(path), app)

	return ExperimentResult{
		AppName:         app,
		Mode:            mode,
		PrimarySuccess:  m.PrimarySuccess,
		RepairAttempts:  m.RepairAttempts,
		FinalSuccess:    m.FinalSuccess,
		ErrorMessage:    m.ErrorMessage,
		DurationSeconds: duration.Seconds(),
		RuleFixes:       m.RuleFixes,
		RelaxedMode:     false,
	}, nil
}

// ✅ Helper: Read Mode from experiment_info.txt or fall back to CSV
func readModeFromExperimentInfo(metricsDir string, appName string) string {
	// Try locations where experiment_info.txt might be:
	possiblePaths := []string{
		filepath.Join(metricsDir, "experiment_info.txt"),                    // Same directory as gen_metrics.json
		filepath.Join("experiments", "out", appName, "experiment_info.txt"), // experiments/out/<AppName>/
		filepath.Join("experiments", appName, "experiment_info.txt"),        // experiments/<AppName>/
	}

	for _, infoPath := range possiblePaths {
		data, err := os.ReadFile(infoPath)
		if err != nil {
			continue
		}

		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			if strings.HasPrefix(line, "Mode:") {
				parts := strings.Split(line, ":")
				if len(parts) >= 2 {
					mode := strings.TrimSpace(parts[1])
					if mode == "rules" || mode == "ml" || mode == "hybrid" {
						return mode
					}
				}
			}
		}
	}

	// ✅ FALLBACK: Read from summary.csv if experiment_info.txt not found
	return readModeFromSummaryCSV(appName)
}

// ✅ NEW: Read mode from summary.csv
func readModeFromSummaryCSV(appName string) string {
	// Initialize cache if not already done
	if csvModeCache == nil {
		csvModeCache = loadSummaryCSVCache()
	}

	if mode, ok := csvModeCache[appName]; ok {
		return mode
	}

	return "unknown"
}

// ✅ NEW: Load summary.csv into memory cache
func loadSummaryCSVCache() map[string]string {
	cache := make(map[string]string)

	csvPath := filepath.Join("experiments", "logs", "summary.csv")
	file, err := os.Open(csvPath)
	if err != nil {
		return cache
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// Skip header
	reader.Read()

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		// CSV format: AppName, Mode, Timestamp, Model, ...
		if len(record) >= 2 {
			appName := strings.TrimSpace(record[0])
			mode := strings.TrimSpace(record[1])
			if mode == "rules" || mode == "ml" || mode == "hybrid" {
				cache[appName] = mode
			}
		}
	}

	return cache
}

// CollectAllExperiments scans experiments/* for metrics files.
func CollectAllExperiments(baseDir string) ([]ExperimentResult, error) {
	var results []ExperimentResult

	// Look in experiments/out/ directory
	outDir := filepath.Join(baseDir, "out")
	entries, err := os.ReadDir(outDir)
	if err != nil {
		// Fallback to baseDir itself
		entries, err = os.ReadDir(baseDir)
		if err != nil {
			return nil, err
		}
	}

	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		appDir := filepath.Join(outDir, e.Name())

		mainPath := filepath.Join(appDir, "gen_metrics.json")
		if _, err := os.Stat(mainPath); err == nil {
			res, err := LoadMetricsFromJSON(mainPath)
			if err == nil {
				results = append(results, res)
			}
		}

		relaxedPath := filepath.Join(appDir, "gen_metrics_relaxed.json")
		if _, err := os.Stat(relaxedPath); err == nil {
			res, err := LoadMetricsFromJSON(relaxedPath)
			if err == nil {
				results = append(results, res)
			}
		}
	}

	// If no results from out/ directory, try root experiments/ directory
	if len(results) == 0 {
		entries, err := os.ReadDir(baseDir)
		if err == nil {
			for _, e := range entries {
				if !e.IsDir() {
					continue
				}
				appDir := filepath.Join(baseDir, e.Name())

				mainPath := filepath.Join(appDir, "gen_metrics.json")
				if _, err := os.Stat(mainPath); err == nil {
					res, err := LoadMetricsFromJSON(mainPath)
					if err == nil {
						results = append(results, res)
					}
				}

				relaxedPath := filepath.Join(appDir, "gen_metrics_relaxed.json")
				if _, err := os.Stat(relaxedPath); err == nil {
					res, err := LoadMetricsFromJSON(relaxedPath)
					if err == nil {
						results = append(results, res)
					}
				}
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
		rows = append(rows, fmt.Sprintf(
			"| %s | %s | %v | %d | %d | %v | %.2f | %s |",
			r.AppName,
			r.Mode,
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
