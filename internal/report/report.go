package report

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type ExperimentResult struct {
	AppName         string  `json:"app_name"`
	Mode            string  `json:"mode"`
	PrimarySuccess  bool    `json:"primary_success"`
	RepairAttempts  int     `json:"repair_attempts"`
	FinalSuccess    bool    `json:"final_success"`
	ErrorMessage    string  `json:"error_message"`
	DurationSeconds float64 `json:"duration_seconds"`
	RuleFixes       int     `json:"rule_fixes"`
	Coverage        float64 `json:"coverage"`
}

// Around line 27-60, UPDATE the LoadMetricsFromJSON function:

func LoadMetricsFromJSON(path string) (ExperimentResult, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return ExperimentResult{}, err
	}

	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		return ExperimentResult{}, err
	}

	app := filepath.Base(filepath.Dir(path))

	// --- Extract fields from JSON ---
	mode := getString(m, "mode")
	if mode == "" {
		mode = readModeFromExperimentInfo(filepath.Dir(path), app)
	}

	// ✅ FIX: Extract duration correctly from nanoseconds
	duration := getFloat(m, "duration_sec")
	if duration == 0 {
		if d := getFloat(m, "Duration"); d > 0 {
			// If Duration is in nanoseconds, convert to seconds
			if d > 1e9 { // Likely nanoseconds
				duration = d / 1e9
			} else {
				duration = d
			}
		}
	}

	repairs := int(getFloat(m, "repair_attempts"))
	if repairs == 0 {
		repairs = int(getFloat(m, "RepairAttempts"))
	}

	fixes := int(getFloat(m, "rule_fixes"))
	if fixes == 0 {
		fixes = int(getFloat(m, "RuleFixes"))
	}

	primarySuccess := getBool(m, "primary_success")
	if !primarySuccess {
		primarySuccess = getBool(m, "PrimarySuccess")
	}

	finalSuccess := getBool(m, "final_success")
	if !finalSuccess {
		finalSuccess = getBool(m, "FinalSuccess")
	}

	errorMsg := getString(m, "error_message")
	if errorMsg == "" {
		errorMsg = getString(m, "ErrorMessage")
	}

	// ✅ FIX: Read coverage from directory
	appDir := filepath.Dir(path)
	coverage := readCoverage(appDir)

	return ExperimentResult{
		AppName:         app,
		Mode:            mode,
		PrimarySuccess:  primarySuccess,
		RepairAttempts:  repairs,
		FinalSuccess:    finalSuccess,
		ErrorMessage:    errorMsg,
		DurationSeconds: duration,
		RuleFixes:       fixes,
		Coverage:        coverage,
	}, nil
}

// OLD readCoverage function (lines ~80-110) - DELETE and REPLACE with:

// readCoverage extracts coverage percentage from gen_metrics JSON or coverage.json
func readCoverage(appDir string) float64 {
	// Try gen_metrics_*.json first (has coverage saved during validation)
	files, _ := filepath.Glob(filepath.Join(appDir, "gen_metrics_*.json"))
	for _, f := range files {
		if data, err := os.ReadFile(f); err == nil {
			var m map[string]interface{}
			if json.Unmarshal(data, &m) == nil {
				// Try all possible coverage field names
				if cov, ok := m["CoveragePct"].(float64); ok && cov > 0 {
					return cov
				}
				if cov, ok := m["coverage_pct"].(float64); ok && cov > 0 {
					return cov
				}
				if cov, ok := m["Coverage"].(float64); ok && cov > 0 {
					return cov
				}
				if cov, ok := m["coverage"].(float64); ok && cov > 0 {
					return cov
				}
			}
		}
	}

	// Fallback: Try coverage.json in the app directory
	if data, err := os.ReadFile(filepath.Join(appDir, "coverage.json")); err == nil {
		var m map[string]interface{}
		if json.Unmarshal(data, &m) == nil {
			if cov, ok := m["CoveragePct"].(float64); ok {
				return cov
			}
			if cov, ok := m["coverage_pct"].(float64); ok {
				return cov
			}
		}
	}

	// Last resort: Check parent directory for coverage.json
	parentDir := filepath.Dir(appDir)
	if data, err := os.ReadFile(filepath.Join(parentDir, "coverage.json")); err == nil {
		var m map[string]interface{}
		if json.Unmarshal(data, &m) == nil {
			if total, ok := m["total_coverage"].(float64); ok && total > 0 {
				return total
			}
		}
	}

	return 0
}

// --- Small JSON helper getters

func getString(m map[string]interface{}, key string) string {
	if v, ok := m[key].(string); ok {
		return v
	}
	return ""
}

func getFloat(m map[string]interface{}, key string) float64 {
	if v, ok := m[key].(float64); ok {
		return v
	}
	return 0
}

func getBool(m map[string]interface{}, key string) bool {
	if v, ok := m[key].(bool); ok {
		return v
	}
	return false
}

var csvModeCache map[string]string

// readModeFromExperimentInfo extracts mode (rules/ml/hybrid) from text or summary.
func readModeFromExperimentInfo(metricsDir string, appName string) string {
	possiblePaths := []string{
		filepath.Join(metricsDir, "experiment_info.txt"),
		filepath.Join("experiments", "out", appName, "experiment_info.txt"),
		filepath.Join("experiments", appName, "experiment_info.txt"),
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

	return readModeFromSummaryCSV(appName)
}

func readModeFromSummaryCSV(appName string) string {
	if csvModeCache == nil {
		csvModeCache = loadSummaryCSVCache()
	}

	if mode, ok := csvModeCache[appName]; ok {
		return mode
	}

	return "unknown"
}

func loadSummaryCSVCache() map[string]string {
	cache := make(map[string]string)

	csvPath := filepath.Join("experiments", "logs", "summary.csv")
	file, err := os.Open(csvPath)
	if err != nil {
		return cache
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, _ = reader.Read() // Skip header

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
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

// CollectAllExperiments gathers all experiment results under experiments/out/.
func CollectAllExperiments(baseDir string) ([]ExperimentResult, error) {
	var results []ExperimentResult

	outDir := filepath.Join(baseDir, "out")
	if _, err := os.Stat(outDir); os.IsNotExist(err) {
		return nil, fmt.Errorf("output directory not found: %s", outDir)
	}

	entries, err := os.ReadDir(outDir)
	if err != nil {
		return nil, err
	}

	for _, e := range entries {
		if !e.IsDir() {
			continue
		}

		appDir := filepath.Join(outDir, e.Name())
		files, err := os.ReadDir(appDir)
		if err != nil {
			continue
		}

		for _, f := range files {
			if strings.HasPrefix(f.Name(), "gen_metrics") && strings.HasSuffix(f.Name(), ".json") {
				filePath := filepath.Join(appDir, f.Name())
				res, err := LoadMetricsFromJSON(filePath)
				if err == nil {
					results = append(results, res)
					break
				}
			}
		}
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("no gen_metrics files found")
	}

	return results, nil
}

// GenerateMarkdownReport writes Markdown table + saves error logs.
func GenerateMarkdownReport(results []ExperimentResult, output string) error {
	header := `# Experimental Evaluation Results

| App | Mode | Primary Success | Repair Attempts | Rule Fixes | Final Success | Duration (s) | Coverage (%) | Error |
|-----|------|----------------|----------------|-------------|----------------|---------------|--------------|-------|
`

	var rows []string
	os.MkdirAll("experiments/logs/errors", 0o755)

	for _, r := range results {
		rows = append(rows, fmt.Sprintf(
			"| %s | %s | %v | %d | %d | %v | %.2f | %.1f | %s |",
			r.AppName,
			r.Mode,
			r.PrimarySuccess,
			r.RepairAttempts,
			r.RuleFixes,
			r.FinalSuccess,
			r.DurationSeconds,
			r.Coverage,
			shorten(r.ErrorMessage, 50),
		))

		// Save full error log for reproducibility
		if r.ErrorMessage != "" {
			errPath := filepath.Join("experiments/logs/errors", fmt.Sprintf("%s_%s.txt", r.AppName, r.Mode))
			_ = os.WriteFile(errPath, []byte(r.ErrorMessage), 0o644)
		}
	}

	content := header + strings.Join(rows, "\n") + "\n"

	os.MkdirAll(filepath.Dir(output), 0o755)
	if err := os.WriteFile(output, []byte(content), 0o644); err != nil {
		return err
	}

	fmt.Printf("📊 Markdown results saved → %s\n", output)
	return nil
}

func GenerateSummaryJSONReport() error {
	baseDir := "experiments"
	output := filepath.Join(baseDir, "logs", "results.md")

	results, err := CollectAllExperiments(baseDir)
	if err != nil {
		return fmt.Errorf("collect experiments: %w", err)
	}

	if len(results) == 0 {
		return fmt.Errorf("no gen_metrics.json files found")
	}

	return GenerateMarkdownReport(results, output)
}

func shorten(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max-3] + "..."
}
