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

		// --- Load metrics.json ---
		var res Result
		if data, err := os.ReadFile(filepath.Join(appDir, "metrics.json")); err == nil {
			_ = json.Unmarshal(data, &res)
		}

		// --- Load experiment_info.txt (metadata) ---
		meta := parseMeta(filepath.Join(appDir, "experiment_info.txt"))

		// --- Add timestamp in case missing ---
		if meta["Timestamp"] == "" {
			meta["Timestamp"] = time.Now().Format(time.RFC3339)
		}

		// --- Build row ---
		row := []string{
			e.Name(),
			meta["Mode"],
			meta["Timestamp"],
			meta["Model"],
			fmt.Sprintf("%v", res.BuildSuccess),
			fmt.Sprintf("%v", res.TestsPass),
			fmt.Sprintf("%.2f", res.CoveragePct),
			fmt.Sprintf("%d", res.LintWarnings),
			meta["MLDuration"],
			meta["RepairAttempts"],
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

// parseMeta reads key:value pairs from experiment_info.txt.
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
