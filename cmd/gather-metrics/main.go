package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Metrics struct {
	Build struct {
		BuildSuccess  bool    `json:"build_success"`
		VetWarnings   int     `json:"vet_warnings"`
		LintWarnings  int     `json:"lint_warnings"`
		TestsPass     bool    `json:"tests_pass"`
		CoveragePct   float64 `json:"coverage_pct"`
		CyclomaticAvg float64 `json:"cyclomatic_avg"`
		GenTimeSec    float64 `json:"gen_time_sec"`
	} `json:"build"`
	Generation struct {
		Duration       int64  `json:"Duration"`
		RuleFixes      int    `json:"RuleFixes"`
		RepairAttempts int    `json:"RepairAttempts"`
		Mode           string `json:"Mode"`
	} `json:"generation"`
}

type Run struct {
	App            string
	Mode           string
	Build          bool
	Tests          bool
	Coverage       float64
	Duration       float64
	LintWarnings   int
	VetWarnings    int
	CyclomaticAvg  float64
	Fixes          int
	RepairAttempts int
}

type Summary struct {
	Mode            string
	Total           int
	BuildSuccessPct float64
	TestsPassPct    float64
	AvgCoverage     float64
	AvgDuration     float64
	AvgLint         float64
	AvgVet          float64
	AvgCyclomatic   float64
	AvgFixes        float64
	AvgRepairs      float64
}

func main() {
	outDir := flag.String("out", "experiments/out", "Output directory")
	reportDir := flag.String("report", "experiments/reports", "Report output directory")
	flag.Parse()

	// Create report directory if it doesn't exist
	os.MkdirAll(*reportDir, 0755)

	data := make(map[string][]Run)
	entries, err := os.ReadDir(*outDir)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		os.Exit(1)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		appName := entry.Name()
		appPath := filepath.Join(*outDir, appName)

		files, err := filepath.Glob(filepath.Join(appPath, "combined_metrics_*.json"))
		if err != nil {
			continue
		}

		for _, file := range files {
			content, err := os.ReadFile(file)
			if err != nil {
				continue
			}

			var m Metrics
			if err := json.Unmarshal(content, &m); err != nil {
				continue
			}

			mode := m.Generation.Mode
			duration := float64(m.Generation.Duration) / 1e9

			run := Run{
				App:            appName,
				Mode:           mode,
				Build:          m.Build.BuildSuccess,
				Tests:          m.Build.TestsPass,
				Coverage:       m.Build.CoveragePct,
				Duration:       duration,
				LintWarnings:   m.Build.LintWarnings,
				VetWarnings:    m.Build.VetWarnings,
				CyclomaticAvg:  m.Build.CyclomaticAvg,
				Fixes:          m.Generation.RuleFixes,
				RepairAttempts: m.Generation.RepairAttempts,
			}

			data[mode] = append(data[mode], run)
		}
	}

	// Calculate summaries
	summaries := calculateSummaries(data)

	// Generate outputs
	generateConsoleOutput(summaries)
	generateConsoleOutput(summaries)
	generateCSV(data, summaries, *reportDir)
	generateMarkdown(data, summaries, *reportDir)
	generateLaTeX(data, summaries, *reportDir)

	fmt.Printf("\n✅ Reports generated in %s\n", *reportDir)
}

func calculateSummaries(data map[string][]Run) []Summary {
	var summaries []Summary
	modes := []string{"rules", "ml", "hybrid"}

	for _, mode := range modes {
		runs, exists := data[mode]
		if !exists || len(runs) == 0 {
			continue
		}

		total := len(runs)
		buildSuccess := 0
		testsPass := 0
		totalCoverage := 0.0
		totalDuration := 0.0
		totalFixes := 0
		totalLint := 0
		totalVet := 0
		totalCyclomatic := 0.0
		totalRepairs := 0

		for _, r := range runs {
			if r.Build {
				buildSuccess++
			}
			if r.Tests {
				testsPass++
			}
			totalCoverage += r.Coverage
			totalDuration += r.Duration
			totalFixes += r.Fixes
			totalLint += r.LintWarnings
			totalVet += r.VetWarnings
			totalCyclomatic += r.CyclomaticAvg
			totalRepairs += r.RepairAttempts
		}

		summaries = append(summaries, Summary{
			Mode:            mode,
			Total:           total,
			BuildSuccessPct: float64(buildSuccess) / float64(total) * 100,
			TestsPassPct:    float64(testsPass) / float64(total) * 100,
			AvgCoverage:     totalCoverage / float64(total),
			AvgDuration:     totalDuration / float64(total),
			AvgLint:         float64(totalLint) / float64(total),
			AvgVet:          float64(totalVet) / float64(total),
			AvgCyclomatic:   totalCyclomatic / float64(total),
			AvgFixes:        float64(totalFixes) / float64(total),
			AvgRepairs:      float64(totalRepairs) / float64(total),
		})
	}

	return summaries
}

func generateConsoleOutput(summaries []Summary) {
	fmt.Println("\n" + strings.Repeat("=", 120))
	fmt.Println("EXPERIMENT SUMMARY - ALL MODES")
	fmt.Println(strings.Repeat("=", 120) + "\n")

	for _, s := range summaries {
		fmt.Printf("Mode: %s\n", strings.ToUpper(s.Mode))
		fmt.Printf("  Total runs: %d\n", s.Total)
		fmt.Printf("  Build success: %.1f%%\n", s.BuildSuccessPct)
		fmt.Printf("  Tests pass: %.1f%%\n", s.TestsPassPct)
		fmt.Printf("  Avg coverage: %.1f%%\n", s.AvgCoverage)
		fmt.Printf("  Avg duration: %.1f s\n", s.AvgDuration)
		fmt.Printf("  Avg lint warnings: %.1f\n", s.AvgLint)
		fmt.Printf("  Avg vet warnings: %.1f\n", s.AvgVet)
		fmt.Printf("  Avg cyclomatic complexity: %.2f\n", s.AvgCyclomatic)
		fmt.Printf("  Avg fixes: %.1f\n", s.AvgFixes)
		fmt.Printf("  Avg repairs: %.1f\n\n", s.AvgRepairs)
	}
}

func generateCSV(data map[string][]Run, summaries []Summary, logDir string) {
	// Summary CSV
	summaryPath := filepath.Join(logDir, "summary.csv")
	f, _ := os.Create(summaryPath)
	defer f.Close()

	w := csv.NewWriter(f)
	w.Write([]string{"Mode", "Total", "Build%", "Tests%", "Coverage%", "Duration(s)", "Lint", "Vet", "Cyclo", "Fixes", "Repairs"})

	for _, s := range summaries {
		w.Write([]string{
			s.Mode,
			fmt.Sprintf("%d", s.Total),
			fmt.Sprintf("%.1f", s.BuildSuccessPct),
			fmt.Sprintf("%.1f", s.TestsPassPct),
			fmt.Sprintf("%.1f", s.AvgCoverage),
			fmt.Sprintf("%.1f", s.AvgDuration),
			fmt.Sprintf("%.1f", s.AvgLint),
			fmt.Sprintf("%.1f", s.AvgVet),
			fmt.Sprintf("%.2f", s.AvgCyclomatic),
			fmt.Sprintf("%.1f", s.AvgFixes),
			fmt.Sprintf("%.1f", s.AvgRepairs),
		})
	}
	w.Flush()
	fmt.Printf("✅ CSV: %s\n", summaryPath)

	// Detailed CSV
	detailPath := filepath.Join(logDir, "detailed.csv")
	f2, _ := os.Create(detailPath)
	defer f2.Close()

	w2 := csv.NewWriter(f2)
	w2.Write([]string{"App", "Mode", "Build", "Tests", "Coverage%", "Duration(s)", "Lint", "Vet", "Cyclo", "Fixes", "Repairs"})

	allRuns := []Run{}
	for _, runs := range data {
		allRuns = append(allRuns, runs...)
	}
	sort.Slice(allRuns, func(i, j int) bool {
		if allRuns[i].App != allRuns[j].App {
			return allRuns[i].App < allRuns[j].App
		}
		return allRuns[i].Mode < allRuns[j].Mode
	})

	for _, r := range allRuns {
		build := "No"
		if r.Build {
			build = "Yes"
		}
		tests := "No"
		if r.Tests {
			tests = "Yes"
		}

		w2.Write([]string{
			r.App,
			r.Mode,
			build,
			tests,
			fmt.Sprintf("%.1f", r.Coverage),
			fmt.Sprintf("%.1f", r.Duration),
			fmt.Sprintf("%d", r.LintWarnings),
			fmt.Sprintf("%d", r.VetWarnings),
			fmt.Sprintf("%.2f", r.CyclomaticAvg),
			fmt.Sprintf("%d", r.Fixes),
			fmt.Sprintf("%d", r.RepairAttempts),
		})
	}
	w2.Flush()
	fmt.Printf("✅ CSV: %s\n", detailPath)
}

func generateMarkdown(data map[string][]Run, summaries []Summary, logDir string) {
	mdPath := filepath.Join(logDir, "report.md")
	f, _ := os.Create(mdPath)
	defer f.Close()

	fmt.Fprintf(f, "# Experimental Results Report\n\n")
	fmt.Fprintf(f, "## Summary by Mode\n\n")
	fmt.Fprintf(f, "| Mode | Total | Build%% | Tests%% | Coverage%% | Duration(s) | Lint | Vet | Cyclo | Fixes | Repairs |\n")
	fmt.Fprintf(f, "|------|-------|--------|--------|-----------|-------------|------|-----|-------|-------|----------|\n")

	for _, s := range summaries {
		fmt.Fprintf(f, "| %s | %d | %.1f | %.1f | %.1f | %.1f | %.1f | %.1f | %.2f | %.1f | %.1f |\n",
			s.Mode, s.Total, s.BuildSuccessPct, s.TestsPassPct, s.AvgCoverage, s.AvgDuration,
			s.AvgLint, s.AvgVet, s.AvgCyclomatic, s.AvgFixes, s.AvgRepairs)
	}

	fmt.Fprintf(f, "\n## Detailed Results\n\n")
	fmt.Fprintf(f, "| App | Mode | Build | Tests | Coverage%% | Duration(s) | Lint | Vet | Cyclo | Fixes | Repairs |\n")
	fmt.Fprintf(f, "|-----|------|-------|-------|-----------|-------------|------|-----|-------|-------|----------|\n")

	allRuns := []Run{}
	for _, runs := range data {
		allRuns = append(allRuns, runs...)
	}
	sort.Slice(allRuns, func(i, j int) bool {
		if allRuns[i].App != allRuns[j].App {
			return allRuns[i].App < allRuns[j].App
		}
		return allRuns[i].Mode < allRuns[j].Mode
	})

	for _, r := range allRuns {
		build := "❌"
		if r.Build {
			build = "✅"
		}
		tests := "❌"
		if r.Tests {
			tests = "✅"
		}

		fmt.Fprintf(f, "| %s | %s | %s | %s | %.1f | %.1f | %d | %d | %.2f | %d | %d |\n",
			r.App, r.Mode, build, tests, r.Coverage, r.Duration, r.LintWarnings, r.VetWarnings,
			r.CyclomaticAvg, r.Fixes, r.RepairAttempts)
	}

	fmt.Printf("✅ Markdown: %s\n", mdPath)
}

func generateLaTeX(data map[string][]Run, summaries []Summary, logDir string) {
	texPath := filepath.Join(logDir, "report.tex")
	f, _ := os.Create(texPath)
	defer f.Close()

	fmt.Fprintf(f, "\\documentclass{article}\n")
	fmt.Fprintf(f, "\\usepackage{booktabs}\n")
	fmt.Fprintf(f, "\\begin{document}\n\n")

	fmt.Fprintf(f, "\\section*{Experimental Results}\n\n")

	fmt.Fprintf(f, "\\begin{table}[h]\n")
	fmt.Fprintf(f, "\\centering\n")
	fmt.Fprintf(f, "\\caption{Summary by Mode}\n")
	fmt.Fprintf(f, "\\begin{tabular}{lrrrrrrrrrr}\n")
	fmt.Fprintf(f, "\\toprule\n")
	fmt.Fprintf(f, "Mode & Total & Build\\%% & Tests\\%% & Cov\\%% & Dur(s) & Lint & Vet & Cyclo & Fixes & Repairs \\\\\n")
	fmt.Fprintf(f, "\\midrule\n")

	for _, s := range summaries {
		fmt.Fprintf(f, "%s & %d & %.1f & %.1f & %.1f & %.1f & %.1f & %.1f & %.2f & %.1f & %.1f \\\\\n",
			s.Mode, s.Total, s.BuildSuccessPct, s.TestsPassPct, s.AvgCoverage, s.AvgDuration,
			s.AvgLint, s.AvgVet, s.AvgCyclomatic, s.AvgFixes, s.AvgRepairs)
	}

	fmt.Fprintf(f, "\\bottomrule\n")
	fmt.Fprintf(f, "\\end{tabular}\n")
	fmt.Fprintf(f, "\\end{table}\n\n")

	fmt.Fprintf(f, "\\end{document}\n")

	fmt.Printf("✅ LaTeX: %s\n", texPath)
}
