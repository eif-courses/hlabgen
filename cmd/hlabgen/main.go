package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/eif-courses/hlabgen/internal/assemble"
	"github.com/eif-courses/hlabgen/internal/input"
	"github.com/eif-courses/hlabgen/internal/metrics"
	mlinternal "github.com/eif-courses/hlabgen/internal/ml"
	"github.com/eif-courses/hlabgen/internal/rules"
	"github.com/eif-courses/hlabgen/internal/validate"
	"github.com/joho/godotenv"
)

func main() {
	// --- Load environment variables ---
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  Warning: .env not found (using system environment)")
	}

	in := flag.String("input", "experiments/input/library.json", "path to schema.json")
	mode := flag.String("mode", "hybrid", "rules|ml|hybrid")
	out := flag.String("out", "experiments/out/LibraryAPI", "output directory")
	flag.Parse()

	schema, err := input.Load(*in)
	if err != nil {
		log.Fatal(err)
	}

	// --- 1) Rule-based scaffold ---
	if _, err := rules.Scaffold(*out, schema.AppName); err != nil {
		log.Fatal(err)
	}

	// --- 2) ML Layer (optional) ---
	var genMetrics mlinternal.GenerationMetrics
	if *mode == "ml" || *mode == "hybrid" {
		files, metrics, err := callML(schema)
		if err != nil {
			log.Fatal(err)
		}
		genMetrics = metrics

		if err := assemble.WriteMany(*out, files); err != nil {
			log.Fatal(err)
		}

		// Save ML generation metrics next to output
		if err := saveGenMetrics(filepath.Join(*out, "gen_metrics.json"), metrics); err != nil {
			log.Printf("⚠️ Failed to write generation metrics: %v\n", err)
		}
	}

	// --- 3) Validation & Runtime metrics ---
	m, err := validate.Run(*out)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n✅ Summary:\n")
	fmt.Printf("  • BuildSuccess = %v\n", m.BuildSuccess)
	fmt.Printf("  • LintWarnings = %d\n", m.LintWarnings)
	fmt.Printf("  • TestsPass    = %v\n", m.TestsPass)
	fmt.Printf("  • Coverage     = %.1f%%\n", m.CoveragePct)
	fmt.Printf("  • ML Duration  = %v (repair %d)\n", genMetrics.Duration, genMetrics.RepairAttempts)

	// Save validation + runtime metrics
	if err := saveMetrics(filepath.Join(*out, "metrics.json"), m); err != nil {
		log.Printf("⚠️ Failed to save validation metrics: %v\n", err)
	}

	// 4) Aggregate all metrics across experiments for summary
	summaryPath := "experiments/logs/summary.csv"
	if err := metrics.AggregateToCSV("experiments/out", summaryPath); err != nil {
		log.Printf("⚠️ Failed to aggregate metrics: %v\n", err)
	}

}

// callML invokes the ML generation and returns generated files + timing metrics.
func callML(s input.Schema) ([]assemble.File, mlinternal.GenerationMetrics, error) {
	files, metrics, err := mlinternal.Generate(mlinternal.Schema{
		AppName: s.AppName, Database: s.Database, APIPattern: s.APIPattern,
		Difficulty: s.Difficulty, Entities: s.Entities, Features: s.Features, Objectives: s.Objectives,
	})
	if err != nil {
		return nil, metrics, err
	}

	out := make([]assemble.File, 0, len(files))
	for _, f := range files {
		out = append(out, assemble.File{Filename: f.Filename, Content: f.Code})
	}
	return out, metrics, nil
}

// saveMetrics saves validation metrics after build/test.
func saveMetrics(path string, m metrics.Result) error {
	b, _ := json.MarshalIndent(m, "", "  ")
	return os.WriteFile(path, b, 0o644)
}

// saveGenMetrics saves ML generation timing & reliability metrics.
func saveGenMetrics(path string, m mlinternal.GenerationMetrics) error {
	b, _ := json.MarshalIndent(m, "", "  ")
	return os.WriteFile(path, b, 0o644)
}
