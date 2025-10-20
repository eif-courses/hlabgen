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

// LoadMetricsFromJSON loads a single experiment's metrics JSON.
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

	// ✅ FIX 4: Read mode directly from JSON first (priority)
	mode := ""
	if modeVal, ok := m["mode"].(string); ok {
		mode = modeVal
	}

	// Fallback to reading from experiment_info.txt if not in JSON
	if mode == "" {
		mode = readModeFromExperimentInfo(filepath.Dir(path), app)
	}

	// Extract values from gen_metrics
	duration := 0.0
	if d, ok := m["Duration"].(float64); ok {
		duration = d / 1e9 // Convert nanoseconds to seconds
	}
	// Also try duration_sec field (lowercase, in seconds already)
	if d, ok := m["duration_sec"].(float64); ok {
		duration = d
	}

	repairs := 0
	if r, ok := m["RepairAttempts"].(float64); ok {
		repairs = int(r)
	}
	// Also try lowercase variant
	if r, ok := m["repair_attempts"].(float64); ok {
		repairs = int(r)
	}

	fixes := 0
	if f, ok := m["RuleFixes"].(float64); ok {
		fixes = int(f)
	}
	// Also try lowercase variant
	if f, ok := m["rule_fixes"].(float64); ok {
		fixes = int(f)
	}

	primarySuccess := false
	if p, ok := m["PrimarySuccess"].(bool); ok {
		primarySuccess = p
	}
	// Also try lowercase variant
	if p, ok := m["primary_success"].(bool); ok {
		primarySuccess = p
	}

	finalSuccess := false
	if f, ok := m["FinalSuccess"].(bool); ok {
		finalSuccess = f
	}
	// Also try lowercase variant
	if f, ok := m["final_success"].(bool); ok {
		finalSuccess = f
	}

	errorMsg := ""
	if e, ok := m["ErrorMessage"].(string); ok {
		errorMsg = e
	}
	// Also try lowercase variant
	if e, ok := m["error_message"].(string); ok {
		errorMsg = e
	}

	// Try to load metrics file for coverage
	appDir := filepath.Dir(path)
	coverage := 0.0
	metricsPath := filepath.Join(appDir, "metrics.json")
	if data, err := os.ReadFile(metricsPath); err == nil {
		var metrics map[string]interface{}
		if err := json.Unmarshal(data, &metrics); err == nil {
			if cov, ok := metrics["coverage_pct"].(float64); ok {
				coverage = cov
			}
		}
	}

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

var csvModeCache map[string]string

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
	reader.Read() // Skip header

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

func CollectAllExperiments(baseDir string) ([]ExperimentResult, error) {
	var results []ExperimentResult

	outDir := filepath.Join(baseDir, "out")
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
