// cmd/report/main.go
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/eif-courses/hlabgen/internal/report"
)

func main() {
	// Command-line flags
	mode := flag.String("mode", "standard", "Report mode: standard|comparative|statistics|failures|latex|all")
	baseDir := flag.String("dir", "experiments", "Base experiments directory")
	outputDir := flag.String("out", "experiments/logs", "Output directory for reports")
	flag.Parse()

	fmt.Printf("📊 Generating reports (mode: %s)...\n", *mode)

	// Ensure output directory exists
	if err := os.MkdirAll(*outputDir, 0755); err != nil {
		log.Fatalf("❌ Failed to create output directory: %v", err)
	}

	// Collect all experiment results
	results, err := report.CollectAllExperiments(*baseDir)
	if err != nil {
		log.Fatalf("❌ Failed to collect experiments: %v", err)
	}

	if len(results) == 0 {
		log.Fatal("❌ No experiment results found")
	}

	// Load build metrics (kept for backward compatibility, but we now use result.Coverage)
	buildMetrics := loadBuildMetrics(*baseDir)

	// Generate reports based on mode
	switch *mode {
	case "standard":
		generateStandardReport(results, *outputDir)
	case "comparative":
		generateComparativeReport(results, buildMetrics, *outputDir)
	case "statistics":
		generateStatisticsReport(results, buildMetrics, *outputDir)
	case "failures":
		generateFailuresReport(results, *outputDir)
	case "latex":
		generateLaTeXReport(results, buildMetrics, *outputDir)
	case "all":
		generateAllReports(results, buildMetrics, *outputDir)
	default:
		log.Fatalf("❌ Unknown mode: %s", *mode)
	}

	// Print summary
	printSummary(results)
}

func generateStandardReport(results []report.ExperimentResult, outputDir string) {
	outputPath := filepath.Join(outputDir, "results.md")

	// ✅ UPDATED TABLE WITH LINT/VET WARNINGS
	header := `# Experimental Evaluation Results

| App | Mode | Build | Tests | Coverage | Lint Warnings | Vet Warnings | Primary Success | Repairs | Fixes | Duration (s) |
|-----|------|-------|-------|----------|---------------|--------------|-----------------|---------|-------|--------------|
`

	var rows []string
	for _, r := range results {
		rows = append(rows, fmt.Sprintf(
			"| %s | %s | %v | %v | %.1f%% | %d | %d | %v | %d | %d | %.2f |",
			r.AppName,
			r.Mode,
			r.BuildSuccess,
			r.TestsPass,
			r.Coverage,
			r.LintWarnings, // ✅ NEW
			r.VetWarnings,  // ✅ NEW
			r.PrimarySuccess,
			r.RepairAttempts,
			r.RuleFixes,
			r.DurationSeconds,
		))
	}

	// ✅ CALCULATE SUMMARY STATISTICS
	content := header + strings.Join(rows, "\n") + "\n\n"

	content += "## Summary Statistics\n\n"

	// Count totals
	totalLint := 0
	totalVet := 0
	totalBuildSuccess := 0
	totalTestsPass := 0
	for _, r := range results {
		totalLint += r.LintWarnings
		totalVet += r.VetWarnings
		if r.BuildSuccess {
			totalBuildSuccess++
		}
		if r.TestsPass {
			totalTestsPass++
		}
	}

	content += fmt.Sprintf("- **Total Experiments**: %d\n", len(results))
	content += fmt.Sprintf("- **Build Success Rate**: %.1f%% (%d/%d)\n",
		float64(totalBuildSuccess)/float64(len(results))*100, totalBuildSuccess, len(results))
	content += fmt.Sprintf("- **Test Success Rate**: %.1f%% (%d/%d)\n",
		float64(totalTestsPass)/float64(len(results))*100, totalTestsPass, len(results))
	content += fmt.Sprintf("- **Average Coverage**: %.1f%%\n", avgCoverageFromResults(results, nil))
	content += fmt.Sprintf("- **Total Lint Warnings**: %d (avg: %.1f per app)\n",
		totalLint, float64(totalLint)/float64(len(results)))
	content += fmt.Sprintf("- **Total Vet Warnings**: %d (avg: %.1f per app)\n",
		totalVet, float64(totalVet)/float64(len(results)))
	content += fmt.Sprintf("- **Average Generation Duration**: %.2fs\n", avgDuration(results))
	content += fmt.Sprintf("- **Avg Repairs Needed**: %.1f per app\n", avgRepairs(results))

	// ✅ CODE QUALITY SECTION
	content += "\n## Code Quality Analysis\n\n"

	// Count apps with warnings
	appsWithLint := 0
	appsWithVet := 0
	maxLint := 0
	maxVet := 0
	var maxLintApp, maxVetApp string

	for _, r := range results {
		if r.LintWarnings > 0 {
			appsWithLint++
		}
		if r.VetWarnings > 0 {
			appsWithVet++
		}
		if r.LintWarnings > maxLint {
			maxLint = r.LintWarnings
			maxLintApp = r.AppName
		}
		if r.VetWarnings > maxVet {
			maxVet = r.VetWarnings
			maxVetApp = r.AppName
		}
	}

	content += fmt.Sprintf("- **Apps with Lint Warnings**: %d/%d\n", appsWithLint, len(results))
	content += fmt.Sprintf("- **Apps with Vet Warnings**: %d/%d\n", appsWithVet, len(results))
	if maxLint > 0 {
		content += fmt.Sprintf("- **Highest Lint Warnings**: %s (%d warnings)\n", maxLintApp, maxLint)
	}
	if maxVet > 0 {
		content += fmt.Sprintf("- **Highest Vet Warnings**: %s (%d warnings)\n", maxVetApp, maxVet)
	}

	// ✅ MODE COMPARISON SECTION
	content += "\n## Performance by Mode\n\n"

	modeStats := make(map[string]struct {
		count       int
		builds      int
		tests       int
		avgCov      float64
		avgLint     float64
		avgVet      float64
		avgDuration float64
	})

	for _, r := range results {
		stats := modeStats[r.Mode]
		stats.count++
		if r.BuildSuccess {
			stats.builds++
		}
		if r.TestsPass {
			stats.tests++
		}
		stats.avgCov += r.Coverage
		stats.avgLint += float64(r.LintWarnings)
		stats.avgVet += float64(r.VetWarnings)
		stats.avgDuration += r.DurationSeconds
		modeStats[r.Mode] = stats
	}

	content += "| Mode | Apps | Build Success | Test Success | Avg Coverage | Avg Lint | Avg Vet | Avg Duration |\n"
	content += "|------|------|---------------|--------------|--------------|----------|---------|---------------|\n"

	for mode, stats := range modeStats {
		if stats.count > 0 {
			content += fmt.Sprintf("| %s | %d | %.1f%% | %.1f%% | %.1f%% | %.1f | %.1f | %.2fs |\n",
				mode,
				stats.count,
				float64(stats.builds)/float64(stats.count)*100,
				float64(stats.tests)/float64(stats.count)*100,
				stats.avgCov/float64(stats.count),
				stats.avgLint/float64(stats.count),
				stats.avgVet/float64(stats.count),
				stats.avgDuration/float64(stats.count),
			)
		}
	}

	os.MkdirAll(filepath.Dir(outputPath), 0o755)
	if err := os.WriteFile(outputPath, []byte(content), 0o644); err != nil {
		log.Fatalf("❌ Failed to write standard report: %v", err)
	}
	fmt.Printf("✅ Standard report with lint/vet warnings: %s\n", outputPath)
}

func generateComparativeReport(results []report.ExperimentResult, buildMetrics map[string]BuildMetrics, outputDir string) {
	outputPath := filepath.Join(outputDir, "comparative.md")

	// ✅ FIXED: Use the Mode field directly, NOT the RuleFixes/RepairAttempts logic
	var hybrid, mlOnly, rulesOnly []report.ExperimentResult
	for _, r := range results {
		switch r.Mode {
		case "hybrid":
			hybrid = append(hybrid, r)
		case "ml":
			mlOnly = append(mlOnly, r)
		case "rules":
			rulesOnly = append(rulesOnly, r)
		default:
			// Unknown mode, skip or categorize as rules
			rulesOnly = append(rulesOnly, r)
		}
	}

	var sb strings.Builder
	sb.WriteString("# Comparative Analysis: Generation Modes\n\n")
	sb.WriteString("## Overview\n\n")
	sb.WriteString(fmt.Sprintf("- **Total Experiments**: %d\n", len(results)))
	sb.WriteString(fmt.Sprintf("- **Hybrid Mode**: %d experiments\n", len(hybrid)))
	sb.WriteString(fmt.Sprintf("- **ML Only**: %d experiments\n", len(mlOnly)))
	sb.WriteString(fmt.Sprintf("- **Rules Only**: %d experiments\n\n", len(rulesOnly)))

	sb.WriteString("## Success Rates by Mode\n\n")
	sb.WriteString("| Mode | Experiments | Success Rate | Avg Coverage | Avg Duration (s) | Avg Repairs |\n")
	sb.WriteString("|------|-------------|--------------|--------------|------------------|-------------|\n")

	if len(hybrid) > 0 {
		sb.WriteString(fmt.Sprintf("| Hybrid | %d | %.1f%% | %.1f%% | %.2f | %.1f |\n",
			len(hybrid),
			successRate(hybrid)*100,
			avgCoverageFromResults(hybrid, buildMetrics),
			avgDuration(hybrid),
			avgRepairs(hybrid)))
	}

	if len(mlOnly) > 0 {
		sb.WriteString(fmt.Sprintf("| ML Only | %d | %.1f%% | %.1f%% | %.2f | %.1f |\n",
			len(mlOnly),
			successRate(mlOnly)*100,
			avgCoverageFromResults(mlOnly, buildMetrics),
			avgDuration(mlOnly),
			avgRepairs(mlOnly)))
	}

	if len(rulesOnly) > 0 {
		sb.WriteString(fmt.Sprintf("| Rules Only | %d | %.1f%% | %.1f%% | %.2f | %.1f |\n",
			len(rulesOnly),
			successRate(rulesOnly)*100,
			avgCoverageFromResults(rulesOnly, buildMetrics),
			avgDuration(rulesOnly),
			avgRepairs(rulesOnly)))
	}

	sb.WriteString("\n## Key Findings\n\n")

	if len(hybrid) > 0 && len(mlOnly) > 0 {
		improvement := (successRate(hybrid) - successRate(mlOnly)) * 100
		sb.WriteString(fmt.Sprintf("- Hybrid mode shows **%.1f percentage point** improvement over ML-only\n", improvement))
	}

	sb.WriteString(fmt.Sprintf("- Average rule fixes applied: %.1f per experiment\n", avgRuleFixes(results)))
	sb.WriteString(fmt.Sprintf("- Overall success rate: **%.1f%%**\n", successRate(results)*100))

	if err := os.WriteFile(outputPath, []byte(sb.String()), 0644); err != nil {
		log.Fatalf("❌ Failed to write comparative report: %v", err)
	}
	fmt.Printf("✅ Comparative report: %s\n", outputPath)
}

func generateStatisticsReport(results []report.ExperimentResult, buildMetrics map[string]BuildMetrics, outputDir string) {
	outputPath := filepath.Join(outputDir, "statistics.md")

	// Calculate statistics
	durations := make([]float64, 0, len(results))
	coverages := make([]float64, 0, len(results))
	repairs := make([]float64, 0, len(results))
	ruleFixes := make([]float64, 0, len(results))

	for _, r := range results {
		durations = append(durations, r.DurationSeconds)
		repairs = append(repairs, float64(r.RepairAttempts))
		ruleFixes = append(ruleFixes, float64(r.RuleFixes))
		// ✅ FIX: Use r.Coverage directly instead of buildMetrics lookup
		coverages = append(coverages, r.Coverage)
	}

	var sb strings.Builder
	sb.WriteString("# Statistical Analysis\n\n")

	// Duration statistics
	sb.WriteString("## Generation Duration Statistics\n\n")
	writeStatistics(&sb, durations, "seconds")

	// Coverage statistics
	sb.WriteString("\n## Code Coverage Statistics\n\n")
	writeStatistics(&sb, coverages, "%")

	// Repair attempts
	sb.WriteString("\n## Repair Attempts Statistics\n\n")
	writeStatistics(&sb, repairs, "attempts")

	// Rule fixes
	sb.WriteString("\n## Rule Fixes Statistics\n\n")
	writeStatistics(&sb, ruleFixes, "fixes")

	// Correlation analysis
	sb.WriteString("\n## Correlation Analysis\n\n")
	sb.WriteString("| Metric Pair | Correlation (r) | Interpretation |\n")
	sb.WriteString("|-------------|-----------------|----------------|\n")

	if len(durations) == len(coverages) && len(durations) > 0 {
		r := pearsonCorrelation(durations, coverages)
		sb.WriteString(fmt.Sprintf("| Duration vs Coverage | %.3f | %s |\n", r, interpretCorrelation(r)))
	}

	if len(repairs) == len(coverages) && len(repairs) > 0 {
		r := pearsonCorrelation(repairs, coverages)
		sb.WriteString(fmt.Sprintf("| Repairs vs Coverage | %.3f | %s |\n", r, interpretCorrelation(r)))
	}

	if len(ruleFixes) == len(coverages) && len(ruleFixes) > 0 {
		r := pearsonCorrelation(ruleFixes, coverages)
		sb.WriteString(fmt.Sprintf("| Rule Fixes vs Coverage | %.3f | %s |\n", r, interpretCorrelation(r)))
	}

	// Success rate analysis
	sb.WriteString("\n## Success Rate Analysis\n\n")
	totalSuccess := 0
	for _, r := range results {
		if r.FinalSuccess {
			totalSuccess++
		}
	}
	successPct := float64(totalSuccess) / float64(len(results)) * 100
	sb.WriteString(fmt.Sprintf("- **Overall Success Rate**: %.1f%% (%d/%d)\n", successPct, totalSuccess, len(results)))
	sb.WriteString(fmt.Sprintf("- **Primary Success Rate** (no repairs): %.1f%%\n", primarySuccessRate(results)*100))
	sb.WriteString(fmt.Sprintf("- **Recovery Rate** (success after repair): %.1f%%\n", recoveryRate(results)*100))

	if err := os.WriteFile(outputPath, []byte(sb.String()), 0644); err != nil {
		log.Fatalf("❌ Failed to write statistics report: %v", err)
	}
	fmt.Printf("✅ Statistics report: %s\n", outputPath)
}

func generateFailuresReport(results []report.ExperimentResult, outputDir string) {
	outputPath := filepath.Join(outputDir, "failures.md")

	failures := []report.ExperimentResult{}
	for _, r := range results {
		if !r.FinalSuccess {
			failures = append(failures, r)
		}
	}

	var sb strings.Builder
	sb.WriteString("# Failure Analysis Report\n\n")
	sb.WriteString(fmt.Sprintf("## Overview\n\n"))
	sb.WriteString(fmt.Sprintf("- **Total Experiments**: %d\n", len(results)))
	sb.WriteString(fmt.Sprintf("- **Failures**: %d\n", len(failures)))
	sb.WriteString(fmt.Sprintf("- **Success Rate**: %.1f%%\n\n", (1.0-float64(len(failures))/float64(len(results)))*100))

	if len(failures) == 0 {
		sb.WriteString("🎉 **No failures detected!** All experiments succeeded.\n")
	} else {
		// Categorize failures
		jsonErrors := 0
		syntaxErrors := 0
		buildErrors := 0
		otherErrors := 0

		for _, f := range failures {
			msg := strings.ToLower(f.ErrorMessage)
			if strings.Contains(msg, "json") || strings.Contains(msg, "parse") {
				jsonErrors++
			} else if strings.Contains(msg, "syntax") {
				syntaxErrors++
			} else if strings.Contains(msg, "build") {
				buildErrors++
			} else {
				otherErrors++
			}
		}

		sb.WriteString("## Failure Categories\n\n")
		sb.WriteString("| Category | Count | Percentage |\n")
		sb.WriteString("|----------|-------|------------|\n")
		sb.WriteString(fmt.Sprintf("| JSON Parse Errors | %d | %.1f%% |\n", jsonErrors, float64(jsonErrors)/float64(len(failures))*100))
		sb.WriteString(fmt.Sprintf("| Syntax Errors | %d | %.1f%% |\n", syntaxErrors, float64(syntaxErrors)/float64(len(failures))*100))
		sb.WriteString(fmt.Sprintf("| Build Errors | %d | %.1f%% |\n", buildErrors, float64(buildErrors)/float64(len(failures))*100))
		sb.WriteString(fmt.Sprintf("| Other | %d | %.1f%% |\n\n", otherErrors, float64(otherErrors)/float64(len(failures))*100))

		sb.WriteString("## Individual Failures\n\n")
		for i, f := range failures {
			sb.WriteString(fmt.Sprintf("### %d. %s\n\n", i+1, f.AppName))
			sb.WriteString(fmt.Sprintf("- **Error**: %s\n", truncate(f.ErrorMessage, 150)))
			sb.WriteString(fmt.Sprintf("- **Repair Attempts**: %d\n", f.RepairAttempts))
			sb.WriteString(fmt.Sprintf("- **Rule Fixes Applied**: %d\n", f.RuleFixes))
			sb.WriteString(fmt.Sprintf("- **Duration**: %.2fs\n\n", f.DurationSeconds))
		}
	}

	if err := os.WriteFile(outputPath, []byte(sb.String()), 0644); err != nil {
		log.Fatalf("❌ Failed to write failures report: %v", err)
	}
	fmt.Printf("✅ Failures report: %s\n", outputPath)
}

// ✅ ALSO UPDATE generateLaTeXReport FUNCTION:

func generateLaTeXReport(results []report.ExperimentResult, buildMetrics map[string]BuildMetrics, outputDir string) {
	outputPath := filepath.Join(outputDir, "tables.tex")

	var sb strings.Builder

	// Main results table with lint/vet
	sb.WriteString("% Experimental Results with Code Quality Metrics\n")
	sb.WriteString("\\begin{table}[htbp]\n")
	sb.WriteString("\\centering\n")
	sb.WriteString("\\caption{Experimental Results with Lint and Vet Warnings}\n")
	sb.WriteString("\\label{tab:results-quality}\n")
	sb.WriteString("\\begin{tabular}{lccccccc}\n")
	sb.WriteString("\\toprule\n")
	sb.WriteString("\\textbf{App} & \\textbf{Mode} & \\textbf{Build} & \\textbf{Tests} & \\textbf{Cov} & \\textbf{Lint} & \\textbf{Vet} & \\textbf{Dur} \\\\\n")
	sb.WriteString("\\midrule\n")

	sort.Slice(results, func(i, j int) bool {
		return results[i].AppName < results[j].AppName
	})

	for _, r := range results {
		sb.WriteString(fmt.Sprintf("%s & %s & %s & %s & %.1f & %d & %d & %.1f \\\\\n",
			escapeLaTeX(r.AppName),
			r.Mode,
			boolToLaTeX(r.BuildSuccess),
			boolToLaTeX(r.TestsPass),
			r.Coverage,
			r.LintWarnings,
			r.VetWarnings,
			r.DurationSeconds,
		))
	}

	sb.WriteString("\\bottomrule\n")
	sb.WriteString("\\end{tabular}\n")
	sb.WriteString("\\end{table}\n\n")

	// Summary statistics table
	sb.WriteString("% Summary Statistics\n")
	sb.WriteString("\\begin{table}[htbp]\n")
	sb.WriteString("\\centering\n")
	sb.WriteString("\\caption{Code Quality Summary}\n")
	sb.WriteString("\\label{tab:quality-summary}\n")
	sb.WriteString("\\begin{tabular}{lcc}\n")
	sb.WriteString("\\toprule\n")
	sb.WriteString("\\textbf{Metric} & \\textbf{Total} & \\textbf{Average} \\\\\n")
	sb.WriteString("\\midrule\n")

	totalLint := 0
	totalVet := 0
	for _, r := range results {
		totalLint += r.LintWarnings
		totalVet += r.VetWarnings
	}

	sb.WriteString(fmt.Sprintf("Lint Warnings & %d & %.1f \\\\\n", totalLint, float64(totalLint)/float64(len(results))))
	sb.WriteString(fmt.Sprintf("Vet Warnings & %d & %.1f \\\\\n", totalVet, float64(totalVet)/float64(len(results))))
	sb.WriteString(fmt.Sprintf("Avg Coverage & -- & %.1f\\%% \\\\\n", avgCoverageFromResults(results, nil)))
	sb.WriteString(fmt.Sprintf("Avg Duration & -- & %.2fs \\\\\n", avgDuration(results)))

	sb.WriteString("\\bottomrule\n")
	sb.WriteString("\\end{tabular}\n")
	sb.WriteString("\\end{table}\n")

	os.MkdirAll(filepath.Dir(outputPath), 0o755)
	if err := os.WriteFile(outputPath, []byte(sb.String()), 0o644); err != nil {
		log.Fatalf("❌ Failed to write LaTeX report: %v", err)
	}
	fmt.Printf("✅ LaTeX tables with lint/vet warnings: %s\n", outputPath)
}
func generateAllReports(results []report.ExperimentResult, buildMetrics map[string]BuildMetrics, outputDir string) {
	fmt.Println("📊 Generating all report types...")
	generateStandardReport(results, outputDir)
	generateComparativeReport(results, buildMetrics, outputDir)
	generateStatisticsReport(results, buildMetrics, outputDir)
	generateFailuresReport(results, outputDir)
	generateLaTeXReport(results, buildMetrics, outputDir)
	fmt.Println("✅ All reports generated successfully!")
}

// ============================================================================
// Helper functions
// ============================================================================

type BuildMetrics struct {
	BuildSuccess bool    `json:"build_success"`
	TestsPass    bool    `json:"tests_pass"`
	CoveragePct  float64 `json:"coverage_pct"`
}

func loadBuildMetrics(baseDir string) map[string]BuildMetrics {
	metrics := make(map[string]BuildMetrics)

	// First try: look in experiments/out/
	entries, err := os.ReadDir(filepath.Join(baseDir, "out"))
	if err != nil {
		// Second try: look directly in baseDir
		entries, _ = os.ReadDir(baseDir)
	}

	for _, e := range entries {
		if !e.IsDir() {
			continue
		}

		appName := e.Name()

		// Try multiple possible locations
		possiblePaths := []string{
			filepath.Join(baseDir, "out", appName, "metrics.json"),
			filepath.Join(baseDir, "out", appName, "gen_metrics.json"),
			filepath.Join(baseDir, appName, "metrics.json"),
			filepath.Join(baseDir, appName, "gen_metrics.json"),
		}

		for _, metricsPath := range possiblePaths {
			data, err := os.ReadFile(metricsPath)
			if err != nil {
				continue
			}

			var bm BuildMetrics
			if err := json.Unmarshal(data, &bm); err == nil {
				metrics[appName] = bm
				break
			}
		}
	}

	return metrics
}

func successRate(results []report.ExperimentResult) float64 {
	if len(results) == 0 {
		return 0
	}
	success := 0
	for _, r := range results {
		if r.FinalSuccess {
			success++
		}
	}
	return float64(success) / float64(len(results))
}

func avgDuration(results []report.ExperimentResult) float64 {
	if len(results) == 0 {
		return 0
	}
	sum := 0.0
	count := 0
	for _, r := range results {
		// ✅ Only count non-zero durations in average
		if r.DurationSeconds > 0 {
			sum += r.DurationSeconds
			count++
		}
	}
	if count == 0 {
		return 0
	}
	return sum / float64(count)
}

// Add this function if missing, or update existing:

func avgRepairs(results []report.ExperimentResult) float64 {
	if len(results) == 0 {
		return 0
	}
	sum := 0
	for _, r := range results {
		sum += r.RepairAttempts
	}
	return float64(sum) / float64(len(results))
}

func avgRuleFixes(results []report.ExperimentResult) float64 {
	if len(results) == 0 {
		return 0
	}
	sum := 0
	for _, r := range results {
		sum += r.RuleFixes
	}
	return float64(sum) / float64(len(results))
}

func avgCoverageFromResults(results []report.ExperimentResult, buildMetrics map[string]BuildMetrics) float64 {
	if len(results) == 0 {
		return 0
	}
	sum := 0.0
	count := 0
	for _, r := range results {
		// ✅ FIX: Use Coverage field directly from ExperimentResult
		if r.Coverage > 0 {
			sum += r.Coverage
			count++
		}
	}
	if count == 0 {
		return 0
	}
	return sum / float64(count)
}

func primarySuccessRate(results []report.ExperimentResult) float64 {
	if len(results) == 0 {
		return 0
	}
	success := 0
	for _, r := range results {
		if r.PrimarySuccess {
			success++
		}
	}
	return float64(success) / float64(len(results))
}

func recoveryRate(results []report.ExperimentResult) float64 {
	failed := 0
	recovered := 0
	for _, r := range results {
		if !r.PrimarySuccess {
			failed++
			if r.FinalSuccess {
				recovered++
			}
		}
	}
	if failed == 0 {
		return 0
	}
	return float64(recovered) / float64(failed)
}

type Stats struct {
	Mean   float64
	StdDev float64
	Min    float64
	Max    float64
	CI95   [2]float64
}

func calculateStats(values []float64) Stats {
	if len(values) == 0 {
		return Stats{}
	}

	// Mean
	sum := 0.0
	for _, v := range values {
		sum += v
	}
	mean := sum / float64(len(values))

	// Std Dev
	variance := 0.0
	for _, v := range values {
		variance += math.Pow(v-mean, 2)
	}
	stdDev := math.Sqrt(variance / float64(len(values)))

	// Min/Max
	min, max := values[0], values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	// 95% CI
	marginError := 1.96 * (stdDev / math.Sqrt(float64(len(values))))

	return Stats{
		Mean:   mean,
		StdDev: stdDev,
		Min:    min,
		Max:    max,
		CI95:   [2]float64{mean - marginError, mean + marginError},
	}
}

func writeStatistics(sb *strings.Builder, values []float64, unit string) {
	stats := calculateStats(values)
	sb.WriteString(fmt.Sprintf("- **Mean**: %.2f %s\n", stats.Mean, unit))
	sb.WriteString(fmt.Sprintf("- **Std Dev**: %.2f %s\n", stats.StdDev, unit))
	sb.WriteString(fmt.Sprintf("- **Min**: %.2f %s\n", stats.Min, unit))
	sb.WriteString(fmt.Sprintf("- **Max**: %.2f %s\n", stats.Max, unit))
	sb.WriteString(fmt.Sprintf("- **95%% CI**: [%.2f, %.2f] %s\n", stats.CI95[0], stats.CI95[1], unit))
}

func pearsonCorrelation(x, y []float64) float64 {
	if len(x) != len(y) || len(x) == 0 {
		return 0
	}

	meanX := 0.0
	meanY := 0.0
	for i := range x {
		meanX += x[i]
		meanY += y[i]
	}
	meanX /= float64(len(x))
	meanY /= float64(len(y))

	var numerator, denomX, denomY float64
	for i := range x {
		dx := x[i] - meanX
		dy := y[i] - meanY
		numerator += dx * dy
		denomX += dx * dx
		denomY += dy * dy
	}

	if denomX == 0 || denomY == 0 {
		return 0
	}

	return numerator / math.Sqrt(denomX*denomY)
}

func interpretCorrelation(r float64) string {
	absR := math.Abs(r)
	var strength string
	if absR < 0.3 {
		strength = "weak"
	} else if absR < 0.7 {
		strength = "moderate"
	} else {
		strength = "strong"
	}

	direction := "positive"
	if r < 0 {
		direction = "negative"
	}

	return fmt.Sprintf("%s %s", strength, direction)
}

func boolToLaTeX(b bool) string {
	if b {
		return "\\checkmark"
	}
	return "\\times"
}

func escapeLaTeX(s string) string {
	s = strings.ReplaceAll(s, "_", "\\_")
	s = strings.ReplaceAll(s, "&", "\\&")
	s = strings.ReplaceAll(s, "%", "\\%")
	s = strings.ReplaceAll(s, "#", "\\#")
	return s
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max-3] + "..."
}

func printSummary(results []report.ExperimentResult) {
	success := 0
	total := len(results)
	for _, r := range results {
		if r.FinalSuccess {
			success++
		}
	}

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("📊 REPORT GENERATION SUMMARY")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("✅ %d/%d experiments succeeded (%.1f%%)\n", success, total, float64(success)/float64(total)*100)
	fmt.Printf("📁 Reports saved in: experiments/logs/\n")
	fmt.Println(strings.Repeat("=", 60))
}
