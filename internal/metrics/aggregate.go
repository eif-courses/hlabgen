package metrics

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// AggregateToCSV scans all experiment result folders under baseDir
// and appends reproducibility results to a single summary CSV.
// If the file already exists, new results are appended instead of overwriting.
func AggregateToCSV(baseDir, outputPath string) error {
	var newRows [][]string

	entries, err := os.ReadDir(baseDir)
	if err != nil {
		return fmt.Errorf("failed to read baseDir: %w", err)
	}

	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		appDir := filepath.Join(baseDir, e.Name())

		// --- Load metrics.json (build/test metrics) ---
		var res Result
		if data, err := os.ReadFile(filepath.Join(appDir, "metrics.json")); err == nil {
			_ = json.Unmarshal(data, &res)
		}

		// ✅ FIX: Load gen_metrics_*.json for ML-specific data (duration, repairs)
		var mlDuration string = ""
		var repairAttempts string = ""

		genFiles, _ := filepath.Glob(filepath.Join(appDir, "gen_metrics_*.json"))
		if len(genFiles) > 0 {
			// Use the latest gen_metrics file
			latestGenFile := genFiles[len(genFiles)-1]
			if data, err := os.ReadFile(latestGenFile); err == nil {
				var genMetrics map[string]interface{}
				if json.Unmarshal(data, &genMetrics) == nil {
					// Extract duration - might be in nanoseconds or seconds
					if d, ok := genMetrics["Duration"].(float64); ok && d > 0 {
						// If Duration > 1e9, it's probably in nanoseconds
						if d > 1e9 {
							mlDuration = fmt.Sprintf("%.2f", d/1e9)
						} else {
							mlDuration = fmt.Sprintf("%.2f", d)
						}
					}
					// Extract repair attempts
					if r, ok := genMetrics["RepairAttempts"].(float64); ok {
						repairAttempts = fmt.Sprintf("%.0f", r)
					}
				}
			}
		}

		// --- Load experiment_info.txt (metadata) ---
		meta := parseMeta(filepath.Join(appDir, "experiment_info.txt"))

		// ✅ FIX: Override with actual ML metrics from gen_metrics JSON if available
		if mlDuration != "" {
			meta["MLDuration"] = mlDuration
		}
		if repairAttempts != "" {
			meta["RepairAttempts"] = repairAttempts
		}

		// --- Add timestamp in case missing ---
		if meta["Timestamp"] == "" {
			meta["Timestamp"] = time.Now().Format(time.RFC3339)
		}

		// --- Build row with all metrics populated ---
		row := []string{
			e.Name(),
			meta["Mode"],
			meta["Timestamp"],
			meta["Model"],
			fmt.Sprintf("%v", res.BuildSuccess),
			fmt.Sprintf("%v", res.TestsPass),
			fmt.Sprintf("%.2f", res.CoveragePct),
			fmt.Sprintf("%d", res.LintWarnings),
			meta["MLDuration"],     // ✅ NOW POPULATED from gen_metrics
			meta["RepairAttempts"], // ✅ NOW POPULATED from gen_metrics
		}
		newRows = append(newRows, row)
	}

	// --- Ensure output dir exists ---
	if err := os.MkdirAll(filepath.Dir(outputPath), 0o755); err != nil {
		return fmt.Errorf("failed to create output dir: %w", err)
	}

	// --- Detect if CSV exists (append or create) ---
	fileExists := false
	if _, err := os.Stat(outputPath); err == nil {
		fileExists = true
	}

	f, err := os.OpenFile(outputPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("failed to open CSV: %w", err)
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()

	// --- Write header only once ---
	if !fileExists {
		header := []string{
			"AppName",
			"Mode",
			"Timestamp",
			"Model",
			"BuildSuccess",
			"TestsPass",
			"CoveragePct",
			"LintWarnings",
			"MLDuration",
			"RepairAttempts",
		}
		if err := writer.Write(header); err != nil {
			return err
		}
	}

	// --- Write new rows ---
	for _, row := range newRows {
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	fmt.Printf("🧾 Appended %d experiments into %s\n", len(newRows), outputPath)
	return nil
}

// parseMeta reads key:value pairs from experiment_info.txt
func parseMeta(path string) map[string]string {
	result := map[string]string{
		"Mode":           "",
		"Timestamp":      "",
		"Model":          "",
		"MLDuration":     "",
		"RepairAttempts": "",
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return result
	}

	lines := strings.Split(string(data), "\n")
	for _, l := range lines {
		if parts := strings.SplitN(l, ":", 2); len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			val := strings.TrimSpace(parts[1])
			switch key {
			case "Mode":
				result["Mode"] = val
			case "Timestamp":
				result["Timestamp"] = val
			case "OpenAI Model":
				result["Model"] = val
			case "MLDuration":
				result["MLDuration"] = val
			case "RepairAttempts":
				result["RepairAttempts"] = val
			}
		}
	}
	return result
}
