package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

// Row represents a single experiment record.
type Row struct {
	AppName        string
	BuildSuccess   bool
	TestsPass      bool
	CoveragePct    float64
	MLDurationSec  float64
	RepairAttempts int
}

func main() {
	logDir := filepath.Join("experiments", "logs")
	filePath := filepath.Join(logDir, "coverage.csv")

	// --- Ensure file exists ---
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Printf("⚠️  No coverage.csv found — creating a new one at %s\n", filePath)
		_ = os.MkdirAll(logDir, 0o755)
		header := "AppName,BuildSuccess,TestsPass,CoveragePct,MLDurationSec,RepairAttempts\n"
		_ = os.WriteFile(filePath, []byte(header), 0o644)
		fmt.Println("✅ Empty CSV initialized. Run experiments first to generate data.")
		return
	}

	// --- Open and parse ---
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("❌ Could not open %s: %v\n", filePath, err)
		return
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println("❌ Failed to read CSV:", err)
		return
	}
	if len(records) <= 1 {
		fmt.Println("⚠️  No data rows found in coverage.csv")
		return
	}

	// --- Parse records ---
	var rows []Row
	for i, rec := range records {
		if i == 0 {
			continue // skip header
		}
		if len(rec) < 6 {
			fmt.Printf("⚠️  Skipping malformed row %d: %v\n", i, rec)
			continue
		}

		cov, _ := strconv.ParseFloat(rec[3], 64)
		dur, _ := strconv.ParseFloat(rec[4], 64)
		repair, _ := strconv.Atoi(rec[5])

		rows = append(rows, Row{
			AppName:        rec[0],
			BuildSuccess:   rec[1] == "true",
			TestsPass:      rec[2] == "true",
			CoveragePct:    cov,
			MLDurationSec:  dur,
			RepairAttempts: repair,
		})
	}

	if len(rows) == 0 {
		fmt.Println("⚠️  No valid data found in CSV.")
		return
	}

	// --- Compute aggregates ---
	var totalCov, totalDur float64
	var totalBuild, totalPass int
	for _, row := range rows {
		if row.BuildSuccess {
			totalBuild++
		}
		if row.TestsPass {
			totalPass++
		}
		totalCov += row.CoveragePct
		totalDur += row.MLDurationSec
	}

	n := float64(len(rows))
	avgCov := totalCov / n
	avgDur := totalDur / n
	buildRate := float64(totalBuild) / n * 100
	passRate := float64(totalPass) / n * 100

	// --- Print Summary ---
	fmt.Println("📊 Aggregate Experiment Summary")
	fmt.Println("--------------------------------")
	fmt.Printf("Total Projects: %d\n", len(rows))
	fmt.Printf("Average Coverage: %.2f%%\n", avgCov)
	fmt.Printf("Average Generation Time: %.2fs\n", avgDur)
	fmt.Printf("Build Success Rate: %.1f%%\n", buildRate)
	fmt.Printf("Test Pass Rate: %.1f%%\n", passRate)
	fmt.Println("--------------------------------")

	fmt.Println("\nDetailed Results:")
	for _, row := range rows {
		fmt.Printf("• %-15s | Cov: %5.1f%% | Time: %6.2fs | Build: %-5v | Test: %-5v | Repairs: %d\n",
			row.AppName, row.CoveragePct, row.MLDurationSec, row.BuildSuccess, row.TestsPass, row.RepairAttempts)
	}
}
