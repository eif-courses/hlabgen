package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/eif-courses/hlabgen/internal/assemble"
	"github.com/eif-courses/hlabgen/internal/input"
	"github.com/eif-courses/hlabgen/internal/metrics"
	mlinternal "github.com/eif-courses/hlabgen/internal/ml"
	"github.com/eif-courses/hlabgen/internal/rules"
	"github.com/eif-courses/hlabgen/internal/validate"
	"github.com/joho/godotenv"
)

func main() {
	// --- 0) Load .env (optional) ---
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  .env not found (using system environment)")
	}

	// --- 1) CLI Flags ---
	in := flag.String("input", "experiments/input/LibraryAPI.json", "path to schema.json")
	mode := flag.String("mode", "hybrid", "rules|ml|hybrid")
	out := flag.String("out", "experiments/out/LibraryAPI", "output directory")
	flag.Parse()

	// --- 2) Load schema ---
	schema, err := input.Load(*in)
	if err != nil {
		log.Fatalf("❌ Failed to load schema: %v", err)
	}

	fmt.Printf("\n🚀 Starting generation for app: %s (mode=%s)\n", schema.AppName, *mode)

	// --- 3) Rule-based scaffold ---
	if _, err := rules.Scaffold(*out, schema.AppName); err != nil {
		log.Fatalf("❌ Scaffold failed: %v", err)
	}

	// Always create placeholder tests in case ML fails
	if err := rules.GenerateFallbackTests(*out, schema.AppName); err != nil {
		log.Printf("⚠️  Failed to create fallback tests: %v\n", err)
	}

	fmt.Println("✅ Rule-based scaffold created")

	// --- 4) ML Layer (with retry and relaxed mode) ---
	var genMetrics mlinternal.GenerationMetrics
	var files []assemble.File

	if *mode == "ml" || *mode == "hybrid" {
		log.Println("🧠 Starting ML-based code generation...")

		files, genMetrics, err = mlinternal.Generate(mlinternal.Schema{
			AppName:    schema.AppName,
			Database:   schema.Database,
			APIPattern: schema.APIPattern,
			Difficulty: schema.Difficulty,
			Entities:   schema.Entities,
			Features:   schema.Features,
			Objectives: schema.Objectives,
		})

		if err != nil {
			log.Printf("⚠️  ML generation failed once: %v", err)
			log.Println("🔁 Retrying with relaxed mode...")
			files, genMetrics, err = mlinternal.GenerateRelaxed(schema)
		}

		if err != nil {
			log.Printf("❌ ML generation failed completely — falling back to rule-based only: %v", err)
		} else {
			if err := assemble.WriteMany(*out, files); err != nil {
				log.Fatalf("❌ Failed to write generated files: %v", err)
			}
			fmt.Printf("✅ ML generation completed (%.2fs)\n", genMetrics.Duration.Seconds())
		}
	} else {
		fmt.Println("⚙️  Skipping ML layer (rules-only mode)")
	}

	// --- 5) Validate & Collect Build Metrics ---
	m, err := validate.Run(*out)
	if err != nil {
		log.Fatalf("❌ Validation failed: %v", err)
	}

	fmt.Printf("\n📊 Summary for %s:\n", schema.AppName)
	fmt.Printf("  • BuildSuccess = %v\n", m.BuildSuccess)
	fmt.Printf("  • LintWarnings = %d\n", m.LintWarnings)
	fmt.Printf("  • TestsPass    = %v\n", m.TestsPass)
	fmt.Printf("  • Coverage     = %.1f%%\n", m.CoveragePct)
	fmt.Printf("  • ML Duration  = %v (repair %d)\n", genMetrics.Duration, genMetrics.RepairAttempts)

	// --- 6) Save metrics ---
	_ = metrics.SaveResult(*out, m)
	_ = metrics.SaveMLMetrics(*out, genMetrics)
	_ = metrics.SaveCombinedMetrics(*out, m, genMetrics)

	// --- 6.5) Save experiment repeatability metadata ---
	metaPath := filepath.Join(*out, "experiment_info.txt")
	meta := fmt.Sprintf(
		"App: %s\nMode: %s\nTimestamp: %s\nOpenAI Model: %s\nBuildSuccess: %v\nTestsPass: %v\nCoverage: %.1f%%\nMLDuration: %v\nRepairAttempts: %d\n",
		schema.AppName,
		*mode,
		time.Now().Format(time.RFC3339),
		getModelName(),
		m.BuildSuccess,
		m.TestsPass,
		m.CoveragePct,
		genMetrics.Duration,
		genMetrics.RepairAttempts,
	)
	if err := os.WriteFile(metaPath, []byte(meta), 0o644); err != nil {
		log.Printf("⚠️  Failed to write experiment metadata: %v\n", err)
	}

	// --- 7) Aggregate all results across experiments ---
	summaryPath := "experiments/logs/summary.csv"
	_ = os.MkdirAll(filepath.Dir(summaryPath), 0o755)
	if err := metrics.AggregateToCSV("experiments/out", summaryPath); err != nil {
		log.Printf("⚠️  Failed to aggregate metrics: %v\n", err)
	}

	fmt.Println("\n✅ Experiment complete.")
}

// getModelName safely reads the model name (for metadata logs)
func getModelName() string {
	model := os.Getenv("OPENAI_MODEL")
	if model == "" {
		return "gpt-4o-mini"
	}
	return model
}
