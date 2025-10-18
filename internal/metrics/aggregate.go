package metrics

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	mlinternal "github.com/eif-courses/hlabgen/internal/ml"
)

// CombinedMetrics represents both validation and ML generation metrics merged.
type CombinedMetrics struct {
	AppName        string
	OutPath        string
	BuildSuccess   bool
	TestsPass      bool
	CoveragePct    float64
	LintWarnings   int
	MLDuration     time.Duration
	RepairAttempts int
	PrimarySuccess bool
	FinalSuccess   bool
	ErrorMessage   string
	Timestamp      string
}

// AggregateToCSV scans all experiment output folders and builds a summary CSV.
func AggregateToCSV(baseDir string, csvPath string) error {
	rows := [][]string{
		{"AppName", "BuildSuccess", "TestsPass", "CoveragePct", "LintWarnings", "MLDurationSec", "RepairAttempts", "PrimarySuccess", "FinalSuccess", "ErrorMessage", "Timestamp"},
	}

	// Walk through all experiment subfolders
	err := filepath.WalkDir(baseDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			return nil
		}

		metricsPath := filepath.Join(path, "metrics.json")
		genMetricsPath := filepath.Join(path, "gen_metrics.json")

		if _, err := os.Stat(metricsPath); err != nil {
			return nil // skip folders without metrics
		}

		var val Result
		data, err := os.ReadFile(metricsPath)
		if err != nil {
			return nil
		}
		_ = json.Unmarshal(data, &val)

		var gen mlinternal.GenerationMetrics
		if b, err := os.ReadFile(genMetricsPath); err == nil {
			_ = json.Unmarshal(b, &gen)
		}

		appName := filepath.Base(path)
		row := []string{
			appName,
			fmt.Sprintf("%v", val.BuildSuccess),
			fmt.Sprintf("%v", val.TestsPass),
			fmt.Sprintf("%.1f", val.CoveragePct),
			fmt.Sprintf("%d", val.LintWarnings),
			fmt.Sprintf("%.2f", gen.Duration.Seconds()),
			fmt.Sprintf("%d", gen.RepairAttempts),
			fmt.Sprintf("%v", gen.PrimarySuccess),
			fmt.Sprintf("%v", gen.FinalSuccess),
			fmt.Sprintf("\"%s\"", gen.ErrorMessage),
			time.Now().Format(time.RFC3339),
		}

		rows = append(rows, row)
		return nil
	})
	if err != nil {
		return err
	}

	// Ensure logs directory exists
	if err := os.MkdirAll(filepath.Dir(csvPath), 0o755); err != nil {
		return err
	}

	f, err := os.Create(csvPath)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, r := range rows {
		fmt.Fprintln(f, joinCSV(r))
	}

	fmt.Printf("📊 Aggregated %d experiments → %s\n", len(rows)-1, csvPath)
	return nil
}

// joinCSV formats CSV-safe lines.
func joinCSV(fields []string) string {
	return fmt.Sprintf("%s", joinWithComma(fields))
}

func joinWithComma(items []string) string {
	if len(items) == 0 {
		return ""
	}
	out := items[0]
	for _, v := range items[1:] {
		out += "," + v
	}
	return out
}
