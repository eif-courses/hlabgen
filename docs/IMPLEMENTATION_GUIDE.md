# HLabGen Gather-Metrics Implementation Guide

## Quick Fix Summary

Your logs aren't "broken" — they're showing **all 450 raw experimental runs** (30 apps × 5 runs × 3 modes).

**Problem**: Looks like duplicates because same App-Mode pairs appear multiple times.  
**Solution**: Show aggregated means ± std dev per App-Mode combination instead of raw runs.

---

## Understanding Current Output

### What You Have Now

**summary.csv** (3 rows - by mode only):
```
Mode,Total,Build%,Tests%,Coverage%,Duration(s),Lint,Vet,Cyclo,Fixes,Repairs
rules,150,100.0,100.0,72.2,0.0,0.0,0.0,3.13,11.5,0.0
ml,150,89.3,0.0,0.0,90.0,12.3,1.5,2.94,24.3,0.0
hybrid,150,92.0,91.3,64.4,77.0,3.8,0.3,2.94,25.1,0.1
```

**results.md** (450 rows - all raw runs):
```
AuctionAPI | hybrid | 70.2% | 123.29s ← Run 1
AuctionAPI | hybrid | 70.2% | 63.74s  ← Run 2
AuctionAPI | hybrid | 70.2% | 79.93s  ← Run 3
AuctionAPI | hybrid | 70.2% | 65.20s  ← Run 4
AuctionAPI | hybrid | 70.2% | 74.81s  ← Run 5
```

↑ **This looks like duplicates but it's 5 runs of the same configuration!**

### What You Should Generate Instead

**results_aggregated.md** (90 rows - per app-mode with means):
```
AuctionAPI | hybrid | 5 | 70.2±0.4% | 81.4±23.6s ← All 5 runs aggregated
AuctionAPI | ml     | 5 | 0.0±0.0%  | 76.9±9.2s
AuctionAPI | rules  | 5 | 72.2±0.1% | 0.03±0.01s
...
```

Much clearer! Shows variance explicitly.

---

## Implementation Steps

### Step 1: Define Aggregated Data Structure

Add to `cmd/gather-metrics/main.go`:

```go
type AggregatedRun struct {
    App              string
    Mode             string
    RunCount         int
    BuildSuccessRate float64  // percentage
    TestsPassRate    float64  // percentage
    AvgCoverage      float64
    StdDevCoverage   float64
    AvgDuration      float64
    StdDevDuration   float64
    AvgLintWarnings  float64
    StdDevLint       float64
    AvgVetWarnings   float64
    AvgCyclomatic    float64
    AvgFixes         float64
    AvgRepairs       float64
}
```

### Step 2: Create Aggregation Function

```go
func aggregateByAppMode(allRuns []Run) []AggregatedRun {
    // Group runs by (App, Mode)
    groups := make(map[string]map[string][]Run)
    
    for _, run := range allRuns {
        if groups[run.App] == nil {
            groups[run.App] = make(map[string][]Run)
        }
        groups[run.App][run.Mode] = append(groups[run.App][run.Mode], run)
    }
    
    // Calculate aggregates
    var aggregated []AggregatedRun
    
    for app := range groups {
        for mode := range groups[app] {
            runs := groups[app][mode]
            agg := aggregateRuns(app, mode, runs)
            aggregated = append(aggregated, agg)
        }
    }
    
    // Sort for consistent output
    sort.Slice(aggregated, func(i, j int) bool {
        if aggregated[i].App != aggregated[j].App {
            return aggregated[i].App < aggregated[j].App
        }
        return aggregated[i].Mode < aggregated[j].Mode
    })
    
    return aggregated
}

func aggregateRuns(app, mode string, runs []Run) AggregatedRun {
    n := float64(len(runs))
    agg := AggregatedRun{
        App:      app,
        Mode:     mode,
        RunCount: len(runs),
    }
    
    // Collect metrics for aggregation
    var (
        buildSuccesses float64
        testSuccesses  float64
        coverages      []float64
        durations      []float64
        lints          []float64
        vets           []float64
        cyclos         []float64
        fixes          []float64
        repairs        []float64
    )
    
    for _, r := range runs {
        if r.Build {
            buildSuccesses += 1.0
        }
        if r.Tests {
            testSuccesses += 1.0
        }
        coverages = append(coverages, r.Coverage)
        durations = append(durations, r.Duration)
        lints = append(lints, float64(r.LintWarnings))
        vets = append(vets, float64(r.VetWarnings))
        cyclos = append(cyclos, r.CyclomaticAvg)
        fixes = append(fixes, float64(r.Fixes))
        repairs = append(repairs, float64(r.RepairAttempts))
    }
    
    // Calculate means and std devs
    agg.BuildSuccessRate = (buildSuccesses / n) * 100.0
    agg.TestsPassRate = (testSuccesses / n) * 100.0
    agg.AvgCoverage, agg.StdDevCoverage = meanStdDev(coverages)
    agg.AvgDuration, agg.StdDevDuration = meanStdDev(durations)
    agg.AvgLintWarnings, agg.StdDevLint = meanStdDev(lints)
    agg.AvgVetWarnings, _ = meanStdDev(vets)
    agg.AvgCyclomatic, _ = meanStdDev(cyclos)
    agg.AvgFixes, _ = meanStdDev(fixes)
    agg.AvgRepairs, _ = meanStdDev(repairs)
    
    return agg
}

func meanStdDev(values []float64) (float64, float64) {
    if len(values) == 0 {
        return 0, 0
    }
    
    sum := 0.0
    for _, v := range values {
        sum += v
    }
    mean := sum / float64(len(values))
    
    variance := 0.0
    for _, v := range values {
        variance += (v - mean) * (v - mean)
    }
    stdDev := math.Sqrt(variance / float64(len(values)))
    
    return mean, stdDev
}
```

### Step 3: Generate Aggregated Report

```go
func generateAggregatedResultsMarkdown(aggregated []AggregatedRun, reportDir string) {
    resultsPath := filepath.Join(reportDir, "results_aggregated.md")
    f, err := os.Create(resultsPath)
    if err != nil {
        fmt.Printf("Error creating aggregated results: %v\n", err)
        return
    }
    defer f.Close()
    
    fmt.Fprintf(f, "# Aggregated Experimental Results\n\n")
    fmt.Fprintf(f, "Each metric shows mean ± standard deviation across %d runs per configuration.\n\n", 5)
    
    fmt.Fprintf(f, "## Summary Table\n\n")
    fmt.Fprintf(f, "| App | Mode | N | Build%% | Tests%% | Coverage%% | Duration(s) | Lint | Vet | Cyclo | Fixes |\n")
    fmt.Fprintf(f, "|-----|------|---|--------|--------|-----------|-------------|------|-----|-------|-------|\n")
    
    for _, agg := range aggregated {
        fmt.Fprintf(f, "| %s | %s | %d | %.1f | %.1f | %.1f±%.1f | %.1f±%.1f | %.1f±%.1f | %.1f | %.2f | %.1f |\n",
            agg.App,
            agg.Mode,
            agg.RunCount,
            agg.BuildSuccessRate,
            agg.TestsPassRate,
            agg.AvgCoverage,
            agg.StdDevCoverage,
            agg.AvgDuration,
            agg.StdDevDuration,
            agg.AvgLintWarnings,
            agg.StdDevLint,
            agg.AvgVetWarnings,
            agg.AvgCyclomatic,
            agg.AvgFixes)
    }
    
    // Add interpretation section
    fmt.Fprintf(f, "\n## Interpretation Guide\n\n")
    fmt.Fprintf(f, "- **Build%%**: Success rate of `go build`\n")
    fmt.Fprintf(f, "- **Tests%%**: Success rate of all unit tests\n")
    fmt.Fprintf(f, "- **Coverage%%**: Code coverage percentage with variance\n")
    fmt.Fprintf(f, "- **Duration(s)**: Generation time with variance (mean±std dev)\n")
    fmt.Fprintf(f, "- **Lint/Vet**: Average linting/vetting warnings\n")
    fmt.Fprintf(f, "- **Cyclo**: Cyclomatic complexity (lower is better)\n")
    fmt.Fprintf(f, "- **Fixes**: Average rule-based fixes applied\n\n")
    
    fmt.Printf("✅ Aggregated Markdown: %s\n", resultsPath)
}
```

### Step 4: Update Main Function

Replace the section that calls report generation functions:

```go
// In main():

// Collect all runs
allRuns := []Run{}
for _, runs := range data {
    allRuns = append(allRuns, runs...)
}

// Generate outputs
generateConsoleOutput(summaries)
generateCSV(data, summaries, *reportDir)
generateMarkdown(data, summaries, *reportDir)
generateLaTeX(data, summaries, *reportDir)
generateResultsMarkdown(data, *reportDir)
generateStatisticsMarkdown(data, summaries, *reportDir)

// NEW: Generate aggregated results
aggregated := aggregateByAppMode(allRuns)
generateAggregatedResultsMarkdown(aggregated, *reportDir)
generateAggregatedCSV(aggregated, *reportDir)  // Optional

fmt.Printf("\n✅ Reports generated in %s\n", *reportDir)
```

### Step 5: Optional - Aggregated CSV

```go
func generateAggregatedCSV(aggregated []AggregatedRun, reportDir string) {
    aggPath := filepath.Join(reportDir, "results_aggregated.csv")
    f, _ := os.Create(aggPath)
    defer f.Close()
    
    w := csv.NewWriter(f)
    w.Write([]string{
        "App", "Mode", "N", "Build%", "Tests%", 
        "AvgCoverage%", "StdDevCoverage%",
        "AvgDuration(s)", "StdDevDuration(s)",
        "AvgLint", "StdDevLint", "AvgVet", "AvgCyclo", "AvgFixes",
    })
    
    for _, agg := range aggregated {
        w.Write([]string{
            agg.App,
            agg.Mode,
            fmt.Sprintf("%d", agg.RunCount),
            fmt.Sprintf("%.1f", agg.BuildSuccessRate),
            fmt.Sprintf("%.1f", agg.TestsPassRate),
            fmt.Sprintf("%.1f", agg.AvgCoverage),
            fmt.Sprintf("%.1f", agg.StdDevCoverage),
            fmt.Sprintf("%.1f", agg.AvgDuration),
            fmt.Sprintf("%.1f", agg.StdDevDuration),
            fmt.Sprintf("%.1f", agg.AvgLintWarnings),
            fmt.Sprintf("%.1f", agg.StdDevLint),
            fmt.Sprintf("%.1f", agg.AvgVetWarnings),
            fmt.Sprintf("%.2f", agg.AvgCyclomatic),
            fmt.Sprintf("%.1f", agg.AvgFixes),
        })
    }
    w.Flush()
    
    fmt.Printf("✅ Aggregated CSV: %s\n", aggPath)
}
```

---

## Updated Output Files

After implementing these changes, you'll have:

```
experiments/reports/
├── summary.csv                 # Mode aggregates (3 rows)
├── summary_by_mode_stats.md    # Statistical analysis (existing)
├── results_aggregated.csv      # App-Mode aggregates (NEW)
├── results_aggregated.md       # App-Mode aggregates Markdown (NEW) ⭐
├── detailed.csv                # All 450 runs raw
├── results.md                  # All 450 runs Markdown (keep as archive)
├── statistics.md               # Distribution analysis
└── report.md                   # Summary tables (existing)
```

**Use `results_aggregated.md` for presentations/papers!** ← Shows 90 rows with means ± std dev

---

## Example Output Comparison

### Before (Raw - 450 rows):
```
| AuctionAPI | hybrid | 70.2% | 123.29 |
| AuctionAPI | hybrid | 70.2% | 63.74  |  ← Looks like duplicates!
| AuctionAPI | hybrid | 70.2% | 79.93  |
| AuctionAPI | ml | 0.0% | 69.73 |
| AuctionAPI | rules | 72.2% | 0.03 |
```

### After (Aggregated - 90 rows):
```
| AuctionAPI | hybrid | 5 | 70.2±0.4% | 81.4±23.6 |  ← Clear variance!
| AuctionAPI | ml     | 5 | 0.0±0.0%  | 76.9±9.2  |
| AuctionAPI | rules  | 5 | 72.2±0.1% | 0.03±0.01 |
```

Much clearer what's happening!

---

## Testing the Implementation

```bash
# Build and run gather-metrics
cd cmd/gather-metrics
go build -o gather-metrics .

# Run it
./gather-metrics -out experiments/out -report experiments/reports

# Check outputs
ls -la experiments/reports/
cat experiments/reports/results_aggregated.md
cat experiments/reports/results_aggregated.csv
```

Expected console output:
```
✅ Reports generated in experiments/reports
✅ Aggregated Markdown: experiments/reports/results_aggregated.md
✅ Aggregated CSV: experiments/reports/results_aggregated.csv
```

---

## Integration with Makefile

Update your `Makefile`:

```makefile
.PHONY: report
report:
	cd cmd/gather-metrics && go run main.go -out experiments/out -report experiments/reports
	@echo "📊 Reports generated:"
	@echo "   - summary.csv (by mode)"
	@echo "   - results_aggregated.md (by app-mode, with means±std dev) ⭐"
	@echo "   - results.md (all raw runs)"
	@echo "   - statistics.md (distribution analysis)"

.PHONY: view-results
view-results:
	cat experiments/reports/results_aggregated.md
```

Then use:
```bash
make all-experiments && make report && make view-results
```

---

## Key Differences Explained

| Aspect | Before | After |
|--------|--------|-------|
| **Rows in results** | 450 (all runs) | 90 (aggregated) |
| **Appearance** | Duplicates look "broken" | Clear means with variance |
| **For presentations** | Confusing | Professional |
| **Raw data preservation** | ✓ (detailed.csv) | ✓ (detailed.csv) |
| **Reproducibility** | ✓ (all runs kept) | ✓ (all runs kept + aggregated) |

---

## Summary

✅ **Your logs aren't broken** — they show all 450 raw runs as intended  
✅ **This is good for reproducibility** — complete data is available  
⚠️ **But it looks confusing** — many duplicate (App, Mode) pairs visible  

**Solution**: Generate aggregated view that shows means ± std dev per App-Mode combination

**Result**: Professional 90-row table perfect for papers, presentations, and analysis

This follows best practices for experimental reporting: **show both raw data AND aggregated statistics**
