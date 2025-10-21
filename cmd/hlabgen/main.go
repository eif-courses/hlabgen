package main

import (
	"bufio"
	"flag"
	"fmt"
	"go/parser"
	"go/token"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/eif-courses/hlabgen/internal/assemble"
	"github.com/eif-courses/hlabgen/internal/input"
	"github.com/eif-courses/hlabgen/internal/metrics"
	mlinternal "github.com/eif-courses/hlabgen/internal/ml"
	"github.com/eif-courses/hlabgen/internal/report"
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

	// --- 3) Main mode dispatcher ---
	var genMetrics mlinternal.GenerationMetrics

	switch *mode {
	case "ml":
		genMetrics = generateMLOnly(schema, *out)
	case "hybrid":
		genMetrics = generateHybrid(schema, *out)
	case "rules":
		genMetrics = generateRulesOnly(schema, *out)
	default:
		log.Fatalf("❌ Unknown mode: %s (use: rules|ml|hybrid)", *mode)
	}

	// --- 4) Validate & Collect Build Metrics ---
	buildMetrics, err := validate.Run(*out)
	if err != nil {
		log.Fatalf("❌ Validation failed: %v", err)
	}

	fmt.Printf("\n📊 Summary for %s (mode: %s):\n", schema.AppName, *mode)
	fmt.Printf("  • BuildSuccess = %v\n", buildMetrics.BuildSuccess)
	fmt.Printf("  • LintWarnings = %d\n", buildMetrics.LintWarnings)
	fmt.Printf("  • TestsPass    = %v\n", buildMetrics.TestsPass)
	fmt.Printf("  • Coverage     = %.1f%%\n", buildMetrics.CoveragePct)
	fmt.Printf("  • Generation Duration = %v\n", genMetrics.Duration)
	fmt.Printf("  • Repair Attempts = %d\n", genMetrics.RepairAttempts)
	fmt.Printf("  • Rule Fixes   = %d\n", genMetrics.RuleFixes)

	// --- 5) Save metrics ---
	genMetrics.Mode = *mode // ✅ Ensure mode is set
	_ = metrics.SaveResult(*out, buildMetrics)
	_ = metrics.SaveMLMetrics(*out, genMetrics)
	_ = metrics.SaveCombinedMetrics(*out, buildMetrics, genMetrics)

	// --- 6) Save experiment repeatability metadata ---
	metaPath := filepath.Join(*out, "experiment_info.txt")
	meta := fmt.Sprintf(
		"App: %s\nMode: %s\nTimestamp: %s\nOpenAI Model: %s\nBuildSuccess: %v\nTestsPass: %v\nCoverage: %.1f%%\nDuration: %v\nRepairAttempts: %d\nRuleFixes: %d\n",
		schema.AppName,
		*mode,
		time.Now().Format(time.RFC3339),
		getModelName(),
		buildMetrics.BuildSuccess,
		buildMetrics.TestsPass,
		buildMetrics.CoveragePct,
		genMetrics.Duration,
		genMetrics.RepairAttempts,
		genMetrics.RuleFixes,
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

	fmt.Println("\n🧾 Generating Markdown summary from JSON metrics...")
	if err := report.GenerateSummaryJSONReport(); err != nil {
		fmt.Println("⚠️ Failed to generate JSON summary:", err)
	} else {
		fmt.Println("✅ Summary successfully written to experiments/logs/results.md")
	}

	fmt.Println("\n✅ Experiment complete..")
}

// ============================================================================
// 🤖 GENERATION MODE: ML-ONLY (Pure LLM, no rules)
// ============================================================================
func generateMLOnly(schema input.Schema, outDir string) mlinternal.GenerationMetrics {
	log.Println("🤖 Starting PURE ML-based generation (no rules)...")

	genMetrics := mlinternal.GenerationMetrics{
		StartTime: time.Now(),
		Mode:      "ml",
	}

	// ✅ KEY DIFFERENCE: SKIP rules.Scaffold() entirely
	// ML-only should generate everything from scratch

	// --- First attempt ---
	mlFiles, mlMetrics, err := mlinternal.Generate(mlinternal.Schema{
		AppName:    schema.AppName,
		Database:   schema.Database,
		APIPattern: schema.APIPattern,
		Difficulty: schema.Difficulty,
		Entities:   schema.Entities,
		Features:   schema.Features,
		Objectives: schema.Objectives,
	})

	genMetrics = mlMetrics
	genMetrics.Mode = "ml" // ✅ Preserve mode
	files := convertGenFiles(mlFiles)

	// --- Retry logic (ML-only still needs this for robustness) ---
	if err != nil {
		log.Printf("⚠️  ML generation failed once: %v", err)
		log.Println("🔁 Retrying with relaxed mode...")
		genMetrics.RepairAttempts++

		mlFiles, mlMetrics, err = mlinternal.GenerateRelaxed(mlinternal.Schema{
			AppName:    schema.AppName,
			Database:   schema.Database,
			APIPattern: schema.APIPattern,
			Difficulty: schema.Difficulty,
			Entities:   schema.Entities,
			Features:   schema.Features,
			Objectives: schema.Objectives,
		})

		genMetrics = mlMetrics
		genMetrics.Mode = "ml" // ✅ Preserve mode
		files = convertGenFiles(mlFiles)
	}

	if err != nil {
		log.Printf("❌ ML generation failed completely: %v", err)
		genMetrics.FinalSuccess = false
		genMetrics.EndTime = time.Now()
		genMetrics.Duration = genMetrics.EndTime.Sub(genMetrics.StartTime)
		return genMetrics
	}

	// ✅ Write ML output WITHOUT rule-based fixes (pure ML output)
	if err := assemble.WriteMany(outDir, files, &genMetrics); err != nil {
		log.Fatalf("❌ Failed to write ML files: %v", err)
	}

	fmt.Println("🔍 Validating Go syntax (ML-only)...")
	syntaxErrors := validateGoSyntax(outDir)
	if len(syntaxErrors) > 0 {
		fmt.Println("⚠️  Syntax errors found in ML-only output:")
		for _, e := range syntaxErrors {
			fmt.Printf("  - %s\n", e)
		}
	} else {
		fmt.Println("✅ All ML-generated files have valid Go syntax")
	}

	// Fix imports and tidy
	fixImportsToModule(outDir)
	tidyDependencies(outDir)

	genMetrics.FinalSuccess = true
	genMetrics.EndTime = time.Now()
	genMetrics.Duration = genMetrics.EndTime.Sub(genMetrics.StartTime)

	fmt.Printf("✅ ML-only generation completed (%.2fs)\n", genMetrics.Duration.Seconds())
	return genMetrics
}

// ============================================================================
// 🔀 GENERATION MODE: HYBRID (Rules + ML + Validation)
// ============================================================================
func generateHybrid(schema input.Schema, outDir string) mlinternal.GenerationMetrics {
	log.Println("🔀 Starting HYBRID generation (rules + ML + validation)...")

	genMetrics := mlinternal.GenerationMetrics{
		StartTime: time.Now(),
		Mode:      "hybrid",
	}

	// --- Step 1: Create rule-based scaffold (structure & templates) ---
	log.Println("📐 Step 1/3: Creating rule-based scaffold...")
	if _, err := rules.Scaffold(outDir, schema.AppName); err != nil {
		log.Fatalf("❌ Scaffold failed: %v", err)
	}
	fmt.Println("✅ Rule-based scaffold created (structure only)")

	// --- Step 2: Generate ML content to enhance scaffold ---
	log.Println("🧠 Step 2/3: Using ML to enhance scaffold logic...")
	mlFiles, mlMetrics, err := mlinternal.Generate(mlinternal.Schema{
		AppName:    schema.AppName,
		Database:   schema.Database,
		APIPattern: schema.APIPattern,
		Difficulty: schema.Difficulty,
		Entities:   schema.Entities,
		Features:   schema.Features,
		Objectives: schema.Objectives,
	})

	genMetrics = mlMetrics
	genMetrics.Mode = "hybrid" // ✅ Preserve hybrid mode
	files := convertGenFiles(mlFiles)

	// Retry if needed
	if err != nil {
		log.Printf("⚠️  ML enhancement failed: %v", err)
		log.Println("🔁 Retrying with relaxed mode...")
		genMetrics.RepairAttempts++

		mlFiles, mlMetrics, err = mlinternal.GenerateRelaxed(mlinternal.Schema{
			AppName:    schema.AppName,
			Database:   schema.Database,
			APIPattern: schema.APIPattern,
			Difficulty: schema.Difficulty,
			Entities:   schema.Entities,
			Features:   schema.Features,
			Objectives: schema.Objectives,
		})

		genMetrics = mlMetrics
		genMetrics.Mode = "hybrid" // ✅ Preserve hybrid mode
		files = convertGenFiles(mlFiles)
	}

	if err != nil {
		log.Printf("❌ ML generation failed, using rules-only fallback")
		genMetrics.Mode = "hybrid" // Still mark as hybrid attempt
		genMetrics.FinalSuccess = false
		genMetrics.EndTime = time.Now()
		genMetrics.Duration = genMetrics.EndTime.Sub(genMetrics.StartTime)
		return genMetrics
	}

	// --- Step 3: Apply rule-based validation & repair ---
	log.Println("🔧 Step 3/3: Applying rule-based validation & fixes...")
	if err := assemble.WriteMany(outDir, files, &genMetrics); err != nil {
		log.Fatalf("❌ Failed to write hybrid files: %v", err)
	}

	// Apply rule-based fixes to ML output
	fmt.Println("🔧 Running rule-based auto-fix on ML-generated files...")
	if err := assemble.FixAllGeneratedFiles(outDir); err != nil {
		log.Printf("⚠️  Some auto-fixes failed: %v", err)
	} else {
		fmt.Println("✅ Rule-based fixes applied successfully")
		genMetrics.RuleFixes++
	}

	// Validate syntax
	fmt.Println("🔍 Validating Go syntax...")
	syntaxErrors := validateGoSyntax(outDir)
	if len(syntaxErrors) > 0 {
		fmt.Println("⚠️  Syntax errors found:")
		for _, e := range syntaxErrors {
			fmt.Printf("  - %s\n", e)
		}
	} else {
		fmt.Println("✅ All generated files have valid Go syntax")
	}

	// Fix imports and tidy
	fixImportsToModule(outDir)
	tidyDependencies(outDir)

	genMetrics.FinalSuccess = true
	genMetrics.EndTime = time.Now()
	genMetrics.Duration = genMetrics.EndTime.Sub(genMetrics.StartTime)

	fmt.Printf("✅ Hybrid generation completed (%.2fs) - rules + ML synergy applied\n", genMetrics.Duration.Seconds())
	return genMetrics
}

// ============================================================================
// ⚙️  GENERATION MODE: RULES-ONLY (Deterministic, no ML)
// ============================================================================
func generateRulesOnly(schema input.Schema, outDir string) mlinternal.GenerationMetrics {
	log.Println("⚙️  Starting PURE RULE-based generation...")

	genMetrics := mlinternal.GenerationMetrics{
		StartTime: time.Now(),
		Mode:      "rules",
	}

	var files []assemble.File

	// Generate files for each entity (template-driven)
	for _, entity := range schema.Entities {
		// Model
		files = append(files, assemble.File{
			Filename: fmt.Sprintf("internal/models/%s.go", strings.ToLower(entity)),
			Content:  rules.GenerateModel(entity),
		})

		// Handler
		files = append(files, assemble.File{
			Filename: fmt.Sprintf("internal/handlers/%s.go", strings.ToLower(entity)),
			Content:  rules.GenerateHandler(entity, schema.AppName),
		})

		// Test
		files = append(files, assemble.File{
			Filename: fmt.Sprintf("internal/handlers/%s_test.go", strings.ToLower(entity)),
			Content:  rules.GenerateTest(entity, schema.AppName),
		})
	}

	// Routes (template-driven)
	files = append(files, assemble.File{
		Filename: "internal/routes/routes.go",
		Content:  rules.GenerateRoutes(schema.Entities, schema.AppName),
	})

	// Tasks markdown
	files = append(files, assemble.File{
		Filename: "tasks.md",
		Content:  rules.GenerateTasksMarkdown(schema.Entities),
	})

	// Write files directly (no ML, no fixes)
	for _, f := range files {
		fullPath := filepath.Join(outDir, f.Filename)
		if err := os.MkdirAll(filepath.Dir(fullPath), 0o755); err != nil {
			log.Fatalf("❌ Failed to create directory: %v", err)
		}

		if err := os.WriteFile(fullPath, []byte(f.Content), 0o644); err != nil {
			log.Fatalf("❌ Failed to write file %s: %v", fullPath, err)
		}

		fmt.Printf("✅ Written: %s\n", fullPath)
	}

	// Validate syntax
	fmt.Println("\n🔍 Validating Go syntax...")
	syntaxErrors := validateGoSyntax(outDir)
	if len(syntaxErrors) > 0 {
		fmt.Println("⚠️  Syntax errors found:")
		for _, e := range syntaxErrors {
			fmt.Printf("  - %s\n", e)
		}
	} else {
		fmt.Println("✅ All rule-generated files have valid Go syntax")
	}

	// Fix imports and tidy
	fixImportsToModule(outDir)
	tidyDependencies(outDir)

	genMetrics.PrimarySuccess = true
	genMetrics.FinalSuccess = true
	genMetrics.RuleFixes = len(files)
	genMetrics.EndTime = time.Now()
	genMetrics.Duration = genMetrics.EndTime.Sub(genMetrics.StartTime)

	fmt.Printf("✅ Rules-only generation completed (%.2fs)\n", genMetrics.Duration.Seconds())
	fmt.Printf("✅ Generated %d files using rule-based templates\n", len(files))

	return genMetrics
}

// ============================================================================
// 🛠️  HELPER FUNCTIONS
// ============================================================================

// convertGenFiles converts []mlinternal.GenFile → []assemble.File
func convertGenFiles(in []mlinternal.GenFile) []assemble.File {
	out := make([]assemble.File, len(in))
	for i, f := range in {
		out[i] = assemble.File{Filename: f.Filename, Content: f.Code}
	}
	return out
}

// tidyDependencies runs go mod tidy
func tidyDependencies(projectDir string) {
	fmt.Println("🔧 Running go mod tidy...")
	tidyCmd := exec.Command("go", "mod", "tidy")
	tidyCmd.Dir = projectDir
	tidyCmd.Stdout = os.Stdout
	tidyCmd.Stderr = os.Stderr
	if err := tidyCmd.Run(); err != nil {
		log.Printf("⚠️  go mod tidy failed: %v", err)
	} else {
		fmt.Println("✅ Dependencies tidied")
	}
}

// fixImportsToModule detects module name and fixes imports automatically
func fixImportsToModule(projectDir string) {
	goMod := filepath.Join(projectDir, "go.mod")
	f, err := os.Open(goMod)
	if err != nil {
		log.Printf("⚠️  No go.mod found in %s (skipping import fix)", projectDir)
		return
	}
	defer f.Close()

	// Detect module name
	scanner := bufio.NewScanner(f)
	moduleName := ""
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "module ") {
			moduleName = strings.TrimSpace(strings.TrimPrefix(line, "module "))
			break
		}
	}
	if moduleName == "" {
		log.Printf("⚠️  Could not detect module name in go.mod (skipping import fix)")
		return
	}

	log.Printf("🔧 Detected module name: %s — fixing imports...", moduleName)

	fixCount := 0
	filepath.WalkDir(projectDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return nil
		}

		original := string(content)
		newContent := original

		// Fix all possible wrong import patterns
		wrongPatterns := []string{
			`"github.com/eif-courses/hlabgen/internal/`,
			`"github.com/yourusername/` + moduleName + `/internal/`,
			`"github.com/yourusername/` + moduleName + `/`,
			`"yourapp/`,
			`"your_project/`,
		}

		for _, wrongPattern := range wrongPatterns {
			correctPattern := fmt.Sprintf(`"%s/internal/`, moduleName)
			if strings.Contains(wrongPattern, "/internal/") {
				newContent = strings.ReplaceAll(newContent, wrongPattern, correctPattern)
			} else {
				newContent = strings.ReplaceAll(newContent, wrongPattern, fmt.Sprintf(`"%s/`, moduleName))
			}
		}

		// Write back if changed
		if newContent != original {
			err = os.WriteFile(path, []byte(newContent), 0o644)
			if err == nil {
				fixCount++
				relPath, _ := filepath.Rel(projectDir, path)
				log.Printf("  ✅ Updated imports in: %s", relPath)
			}
		}
		return nil
	})

	if fixCount > 0 {
		log.Printf("✅ Fixed imports in %d file(s)", fixCount)
	} else {
		log.Println("✅ No import fixes needed")
	}
}

// getModelName returns the OpenAI model name from env or default
func getModelName() string {
	model := os.Getenv("OPENAI_MODEL")
	if model == "" {
		return "gpt-4o-mini"
	}
	return model
}

// validateGoSyntax validates Go syntax for all .go files
func validateGoSyntax(projectPath string) []string {
	var errors []string

	filepath.Walk(projectPath, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}

		fset := token.NewFileSet()
		_, parseErr := parser.ParseFile(fset, path, nil, parser.AllErrors)

		if parseErr != nil {
			relPath, _ := filepath.Rel(projectPath, path)
			errors = append(errors, fmt.Sprintf("%s: %v", relPath, parseErr))
		}

		return nil
	})

	return errors
}
