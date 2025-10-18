package metrics

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// AggregateToCSV scans all experiment result folders under baseDir
// and writes a reproducibility summary CSV with metrics + metadata.
func AggregateToCSV(baseDir, outputPath string) error {
	rows := [][]string{
		{
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
		},
	}

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

		// --- Load gen_metrics.json ---
		var gen map[string]interface{}
		if data, err := os.ReadFile(filepath.Join(appDir, "gen_metrics.json")); err == nil {
			_ = json.Unmarshal(data, &gen)
		}

		// --- Load experiment_info.txt (metadata) ---
		meta := parseMeta(filepath.Join(appDir, "experiment_info.txt"))

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
		rows = append(rows, row)
	}

	// --- Ensure output dir exists ---
	if err := os.MkdirAll(filepath.Dir(outputPath), 0o755); err != nil {
		return fmt.Errorf("failed to create output dir: %w", err)
	}

	// --- Write CSV ---
	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create CSV: %w", err)
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()

	for _, row := range rows {
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	fmt.Printf("🧾 Aggregated %d experiments into %s\n", len(rows)-1, outputPath)
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
