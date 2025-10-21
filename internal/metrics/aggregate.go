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
// and REPLACES the summary CSV with current results (no duplicates).
func AggregateToCSV(baseDir, outputPath string) error {
	// ✅ FIX 1: Collect results in memory first (prevents duplicates)
	resultMap := make(map[string][]string) // key: appName, value: row data

	entries, err := os.ReadDir(baseDir)
	if err != nil {
		return fmt.Errorf("failed to read baseDir: %w", err)
	}

	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		appDir := filepath.Join(baseDir, e.Name())
		appName := e.Name()

		// --- Load metrics.json (build/test metrics) ---
		var res Result
		if data, err := os.ReadFile(filepath.Join(appDir, "metrics.json")); err == nil {
			_ = json.Unmarshal(data, &res)
		}

		// ✅ FIX 2: Load gen_metrics_*.json with CORRECT duration parsing
		var mlDuration string = "0.00"
		var repairAttempts string = "0"

		genFiles, _ := filepath.Glob(filepath.Join(appDir, "gen_metrics_*.json"))
		if len(genFiles) > 0 {
			latestGenFile := genFiles[len(genFiles)-1]
			if data, err := os.ReadFile(latestGenFile); err == nil {
				var genMetrics map[string]interface{}
				if json.Unmarshal(data, &genMetrics) == nil {
					// ✅ Extract duration with smart unit detection
					if d := parseMLDuration(genMetrics); d > 0 {
						mlDuration = fmt.Sprintf("%.2f", d)
					}

					// Extract repair attempts
					if r, ok := genMetrics["repair_attempts"].(float64); ok {
						repairAttempts = fmt.Sprintf("%.0f", r)
					} else if r, ok := genMetrics["RepairAttempts"].(float64); ok {
						repairAttempts = fmt.Sprintf("%.0f", r)
					}
				}
			}
		}

		// --- Load experiment_info.txt (metadata) ---
		meta := parseMeta(filepath.Join(appDir, "experiment_info.txt"))

		// Override with actual ML metrics
		if mlDuration != "0.00" {
			meta["MLDuration"] = mlDuration
		}
		if repairAttempts != "0" {
			meta["RepairAttempts"] = repairAttempts
		}

		// Add timestamp if missing
		if meta["Timestamp"] == "" {
			meta["Timestamp"] = time.Now().Format(time.RFC3339)
		}

		// Build row
		row := []string{
			appName,
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

		// ✅ Store in map (overwrites duplicates automatically)
		resultMap[appName] = row
	}

	// ✅ FIX 3: Write FRESH file (no appending - prevents duplicates)
	if err := os.MkdirAll(filepath.Dir(outputPath), 0o755); err != nil {
		return fmt.Errorf("failed to create output dir: %w", err)
	}

	f, err := os.Create(outputPath) // ← Use Create instead of OpenFile
	if err != nil {
		return fmt.Errorf("failed to create CSV: %w", err)
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()

	// Write header
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

	// Write rows (deduplicated)
	for _, row := range resultMap {
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	fmt.Printf("🧾 Wrote %d unique experiments to %s\n", len(resultMap), outputPath)
	return nil
}

// ✅ NEW: Smart ML duration parser (handles nanoseconds, microseconds, seconds)
func parseMLDuration(m map[string]interface{}) float64 {
	// Try duration_sec first
	if d, ok := m["duration_sec"].(float64); ok && d > 0 {
		return d
	}

	// Try Duration field (could be in different units)
	if d, ok := m["Duration"].(float64); ok && d > 0 {
		// Smart detection based on realistic ranges:
		// - 0.001 to 1000 → already seconds
		// - 10 to 10,000,000 → microseconds (10µs to 10s)
		// - > 10,000,000 → nanoseconds

		if d < 10 {
			// Already in seconds or very small
			return d
		} else if d >= 10 && d <= 10000000 {
			// Microseconds range (10µs to 10 seconds)
			return d / 1e6
		} else {
			// Nanoseconds (> 10 seconds in ns)
			return d / 1e9
		}
	}

	return 0
}

// parseMeta reads key:value pairs from experiment_info.txt
func parseMeta(path string) map[string]string {
	result := map[string]string{
		"Mode":           "",
		"Timestamp":      "",
		"Model":          "",
		"MLDuration":     "0.00",
		"RepairAttempts": "0",
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
			}
		}
	}
	return result
}
