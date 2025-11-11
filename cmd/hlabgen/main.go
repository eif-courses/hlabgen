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
	"regexp"
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
	genMetrics.Mode = *mode // Ensure mode is set
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
	genMetrics.Mode = "ml"
	files := convertGenFiles(mlFiles)

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
		genMetrics.Mode = "ml"
		files = convertGenFiles(mlFiles)
	}

	if err != nil {
		log.Printf("❌ ML generation failed completely: %v", err)
		genMetrics.FinalSuccess = false
		genMetrics.EndTime = time.Now()
		genMetrics.Duration = genMetrics.EndTime.Sub(genMetrics.StartTime)
		return genMetrics
	}

	if err := assemble.WriteMany(outDir, files, &genMetrics); err != nil {
		log.Fatalf("❌ Failed to write ML files: %v", err)
	}

	fixParseIDTypeMismatch(outDir)

	fmt.Println("🔧 Running rule-based auto-fix on ML-generated files...")
	if err := assemble.FixAllGeneratedFiles(outDir); err != nil {
		log.Printf("⚠️  Some auto-fixes failed: %v", err)
	} else {
		fmt.Println("✅ Rule-based fixes applied successfully")
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

	ensureGoMod(outDir, schema.AppName)
	fixImportsToModule(outDir)
	tidyDependencies(outDir)

	genMetrics.FinalSuccess = true
	genMetrics.EndTime = time.Now()
	genMetrics.Duration = genMetrics.EndTime.Sub(genMetrics.StartTime)

	fmt.Printf("✅ ML-only generation completed (%.2fs)\n", genMetrics.Duration.Seconds())
	return genMetrics
}

// ============================================================================
// 🔀 GENERATION MODE: HYBRID (Intelligent Strategy Selection)
// ============================================================================
func generateHybrid(schema input.Schema, outDir string) mlinternal.GenerationMetrics {
	log.Println("🔀 Starting INTELLIGENT HYBRID generation...")

	genMetrics := mlinternal.GenerationMetrics{
		StartTime: time.Now(),
		Mode:      "hybrid",
	}

	// 🆕 STEP 0: Analyze complexity and select strategy
	log.Println("📊 Step 0/4: Analyzing API complexity...")
	complexity := rules.AnalyzeComplexity(schema.Features)
	fmt.Println(complexity.DebugInfo())

	strategy := complexity.GetStrategy()
	log.Printf("🎯 Selected Strategy: %s\n", strategy)

	// Execute appropriate strategy
	switch strategy {
	case "ML_PRIMARY":
		genMetrics = generateMLPrimaryHybrid(schema, outDir, genMetrics)

	case "HYBRID_BALANCED":
		genMetrics = generateHybridBalancedStrategy(schema, outDir, genMetrics)

	case "RULES_PRIMARY":
		genMetrics = generateRulesPrimaryStrategy(schema, outDir, genMetrics)

	default:
		genMetrics.FinalSuccess = false
		genMetrics.ErrorMessage = "unknown strategy: " + strategy
	}

	genMetrics.EndTime = time.Now()
	genMetrics.Duration = genMetrics.EndTime.Sub(genMetrics.StartTime)
	fmt.Printf("✅ Hybrid generation completed (%.2fs) - Strategy: %s\n", genMetrics.Duration.Seconds(), strategy)

	return genMetrics
}

// 🆕 ML-PRIMARY Strategy Implementation
func generateMLPrimaryHybrid(schema input.Schema, outDir string, genMetrics mlinternal.GenerationMetrics) mlinternal.GenerationMetrics {
	log.Println("\n🤖 Executing ML-PRIMARY Strategy")
	log.Println("  • ML generates business logic")
	log.Println("  • Rules create structure & validation")

	// Step 1: Create rule-based scaffold
	log.Println("📐 Step 1/4: Creating rule-based scaffold...")
	if _, err := rules.Scaffold(outDir, schema.AppName); err != nil {
		log.Fatalf("❌ Scaffold failed: %v", err)
	}

	// Generate models and handlers with business logic placeholders
	var files []assemble.File

	for _, entity := range schema.Entities {
		// Model
		files = append(files, assemble.File{
			Filename: fmt.Sprintf("internal/models/%s.go", strings.ToLower(entity)),
			Content:  rules.GenerateMLPrimaryModel(entity, schema.Features),
		})

		// Handler with business logic placeholders
		files = append(files, assemble.File{
			Filename: fmt.Sprintf("internal/handlers/%s.go", strings.ToLower(entity)),
			Content:  rules.GenerateMLPrimaryHandler(entity, schema.Features),
		})
	}

	// Step 2: Generate ML business logic layer
	log.Println("🧠 Step 2/4: Using ML to generate business logic...")
	mlFiles, mlMetrics, err := mlinternal.Generate(mlinternal.Schema{
		AppName:    schema.AppName,
		Database:   schema.Database,
		APIPattern: schema.APIPattern,
		Difficulty: schema.Difficulty,
		Entities:   schema.Entities,
		Features:   schema.Features,
		Objectives: schema.Objectives,
	})

	if err != nil {
		log.Printf("⚠️  ML generation failed: %v", err)
		genMetrics.RepairAttempts++
		genMetrics.FinalSuccess = false
	} else {
		files = append(files, convertGenFiles(mlFiles)...)
		genMetrics = mlMetrics
		genMetrics.FinalSuccess = true
	}

	// Step 3: Apply rule-based validation & repair
	log.Println("🔧 Step 3/4: Applying rule-based validation...")
	if err := assemble.WriteMany(outDir, files, &genMetrics); err != nil {
		log.Fatalf("❌ Failed to write files: %v", err)
	}

	if err := assemble.FixAllGeneratedFiles(outDir); err != nil {
		log.Printf("⚠️  Some auto-fixes failed: %v", err)
	}

	// Step 4: Finalization
	log.Println("✅ Step 4/4: Finalizing...")
	fixParseIDTypeMismatch(outDir)
	ensureGoMod(outDir, schema.AppName)
	fixImportsToModule(outDir)
	tidyDependencies(outDir)

	genMetrics.FinalSuccess = true
	return genMetrics
}

// 🆕 HYBRID-BALANCED Strategy Implementation
func generateHybridBalancedStrategy(schema input.Schema, outDir string, genMetrics mlinternal.GenerationMetrics) mlinternal.GenerationMetrics {
	log.Println("\n⚖️  Executing HYBRID-BALANCED Strategy")
	log.Println("  • Rules scaffold structure")
	log.Println("  • ML fills business logic holes")

	// Step 1: Create rule-based scaffold with business logic hooks
	log.Println("📐 Step 1/4: Creating rule-based scaffold with hooks...")
	if _, err := rules.Scaffold(outDir, schema.AppName); err != nil {
		log.Fatalf("❌ Scaffold failed: %v", err)
	}

	var files []assemble.File

	for _, entity := range schema.Entities {
		// Model with business fields
		files = append(files, assemble.File{
			Filename: fmt.Sprintf("internal/models/%s.go", strings.ToLower(entity)),
			Content:  rules.GenerateHybridBalancedModel(entity, schema.Features),
		})

		// Handler with ML customization hooks
		files = append(files, assemble.File{
			Filename: fmt.Sprintf("internal/handlers/%s.go", strings.ToLower(entity)),
			Content:  rules.GenerateHybridBalancedHandler(entity, schema.Features),
		})

		// ✅ FIXED: Line 351 - GenerateRulesPrimaryTest → GenerateSimpleTest
		// Test
		files = append(files, assemble.File{
			Filename: fmt.Sprintf("internal/handlers/%s_test.go", strings.ToLower(entity)),
			Content:  rules.GenerateSimpleTest(entity, schema.AppName),
		})
	}

	// Step 2: ML fills in the business logic
	log.Println("🧠 Step 2/4: Using ML to fill business logic...")
	mlFiles, mlMetrics, err := mlinternal.Generate(mlinternal.Schema{
		AppName:    schema.AppName,
		Database:   schema.Database,
		APIPattern: schema.APIPattern,
		Difficulty: schema.Difficulty,
		Entities:   schema.Entities,
		Features:   schema.Features,
		Objectives: schema.Objectives,
	})

	if err != nil {
		log.Printf("⚠️  ML generation failed: %v", err)
		genMetrics.RepairAttempts++
		genMetrics.FinalSuccess = false
	} else {
		files = append(files, convertGenFiles(mlFiles)...)
		genMetrics = mlMetrics
		genMetrics.FinalSuccess = true
	}

	// Step 3: Apply rule-based validation
	log.Println("🔧 Step 3/4: Applying rule-based validation...")
	if err := assemble.WriteMany(outDir, files, &genMetrics); err != nil {
		log.Fatalf("❌ Failed to write files: %v", err)
	}

	if err := assemble.FixAllGeneratedFiles(outDir); err != nil {
		log.Printf("⚠️  Some auto-fixes failed: %v", err)
	}

	// Step 4: Finalization
	log.Println("✅ Step 4/4: Finalizing...")
	fixParseIDTypeMismatch(outDir)
	ensureGoMod(outDir, schema.AppName)
	fixImportsToModule(outDir)
	tidyDependencies(outDir)

	genMetrics.FinalSuccess = true
	return genMetrics
}

// 🆕 RULES-PRIMARY Strategy Implementation
func generateRulesPrimaryStrategy(schema input.Schema, outDir string, genMetrics mlinternal.GenerationMetrics) mlinternal.GenerationMetrics {
	log.Println("\n⚙️  Executing RULES-PRIMARY Strategy")
	log.Println("  • 100% rule-based generation")
	log.Println("  • No ML needed for simple CRUD")

	// Step 1: Create complete rule-based scaffold
	log.Println("📐 Step 1/3: Creating complete rule-based scaffold...")
	if _, err := rules.Scaffold(outDir, schema.AppName); err != nil {
		log.Fatalf("❌ Scaffold failed: %v", err)
	}

	var files []assemble.File

	for _, entity := range schema.Entities {
		// ✅ FIXED: Line 416 - GenerateRulesPrimaryModel → GenerateSimpleModel
		// Model (simple, no business logic)
		files = append(files, assemble.File{
			Filename: fmt.Sprintf("internal/models/%s.go", strings.ToLower(entity)),
			Content:  rules.GenerateSimpleModel(entity),
		})

		// ✅ FIXED: Line 422 - GenerateRulesPrimaryHandler → GenerateSimpleHandler
		// Handler (pure CRUD, no business logic)
		files = append(files, assemble.File{
			Filename: fmt.Sprintf("internal/handlers/%s.go", strings.ToLower(entity)),
			Content:  rules.GenerateSimpleHandler(entity, schema.AppName),
		})

		// ✅ FIXED: Line 428 - GenerateRulesPrimaryTest → GenerateSimpleTest
		// Test
		files = append(files, assemble.File{
			Filename: fmt.Sprintf("internal/handlers/%s_test.go", strings.ToLower(entity)),
			Content:  rules.GenerateSimpleTest(entity, schema.AppName),
		})
	}

	// Routes
	files = append(files, assemble.File{
		Filename: "internal/routes/routes.go",
		Content:  rules.GenerateRoutes(schema.Entities, schema.AppName),
	})

	// Tasks
	files = append(files, assemble.File{
		Filename: "tasks.md",
		Content:  rules.GenerateTasksMarkdown(schema.Entities),
	})

	// Step 2: Write files with rule-based fixes
	log.Println("🔧 Step 2/3: Writing files with rule-based validation...")
	if err := assemble.WriteMany(outDir, files, &genMetrics); err != nil {
		log.Fatalf("❌ Failed to write files: %v", err)
	}

	if err := assemble.FixAllGeneratedFiles(outDir); err != nil {
		log.Printf("⚠️  Some auto-fixes failed: %v", err)
	}

	// Step 3: Finalization
	log.Println("✅ Step 3/3: Finalizing...")
	ensureGoMod(outDir, schema.AppName)
	fixImportsToModule(outDir)
	tidyDependencies(outDir)

	genMetrics.PrimarySuccess = true
	genMetrics.FinalSuccess = true
	genMetrics.RuleFixes = len(files)

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

	// Detect if this is a complex API (has features)
	isComplexAPI := len(schema.Features) > 0

	if isComplexAPI {
		log.Println("🔍 Detected complex API with business logic features")
		log.Printf("   Features: %v", schema.Features)
	}

	// Generate files for each entity
	for _, entity := range schema.Entities {
		if isComplexAPI {
			files = append(files, assemble.File{
				Filename: fmt.Sprintf("internal/models/%s.go", strings.ToLower(entity)),
				Content:  rules.GenerateComplexModel(entity, schema.Features),
			})

			generator := rules.NewComplexHandler(entity, schema.AppName, schema.Features)
			files = append(files, assemble.File{
				Filename: fmt.Sprintf("internal/handlers/%s.go", strings.ToLower(entity)),
				Content:  generator.GenerateComplexHandler(),
			})

			log.Printf("   ✅ Generated complex handler for %s (discount=%v, tax=%v, state=%v)",
				entity, generator.HasDiscount, generator.HasTax, generator.HasState)
		} else {
			// ✅ FIXED: Line 507 - GenerateModel → GenerateSimpleModel
			files = append(files, assemble.File{
				Filename: fmt.Sprintf("internal/models/%s.go", strings.ToLower(entity)),
				Content:  rules.GenerateSimpleModel(entity),
			})

			// ✅ FIXED: Line 512 - GenerateHandler → GenerateSimpleHandler
			files = append(files, assemble.File{
				Filename: fmt.Sprintf("internal/handlers/%s.go", strings.ToLower(entity)),
				Content:  rules.GenerateSimpleHandler(entity, schema.AppName),
			})
		}

		// ✅ FIXED: Line 518 - GenerateTest → GenerateSimpleTest
		files = append(files, assemble.File{
			Filename: fmt.Sprintf("internal/handlers/%s_test.go", strings.ToLower(entity)),
			Content:  rules.GenerateSimpleTest(entity, schema.AppName),
		})
	}

	// Routes
	files = append(files, assemble.File{
		Filename: "internal/routes/routes.go",
		Content:  rules.GenerateRoutes(schema.Entities, schema.AppName),
	})

	// Tasks markdown
	files = append(files, assemble.File{
		Filename: "tasks.md",
		Content:  rules.GenerateTasksMarkdown(schema.Entities),
	})

	// Write files directly
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

	ensureGoMod(outDir, schema.AppName)
	fixImportsToModule(outDir)
	tidyDependencies(outDir)

	genMetrics.PrimarySuccess = true
	genMetrics.FinalSuccess = true
	genMetrics.RuleFixes = len(files)
	genMetrics.EndTime = time.Now()
	genMetrics.Duration = genMetrics.EndTime.Sub(genMetrics.StartTime)

	if isComplexAPI {
		fmt.Printf("✅ Rules-only generation with business logic completed (%.2fs)\n", genMetrics.Duration.Seconds())
	} else {
		fmt.Printf("✅ Rules-only generation completed (%.2fs)\n", genMetrics.Duration.Seconds())
	}
	fmt.Printf("✅ Generated %d files using rule-based templates\n", len(files))

	return genMetrics
}

// ============================================================================
// 🛠️  HELPER FUNCTIONS
// ============================================================================

func convertGenFiles(in []mlinternal.GenFile) []assemble.File {
	out := make([]assemble.File, len(in))
	for i, f := range in {
		out[i] = assemble.File{Filename: f.Filename, Content: f.Code}
	}
	return out
}

func ensureGoMod(projectDir string, appName string) {
	goModPath := filepath.Join(projectDir, "go.mod")

	if _, err := os.Stat(goModPath); err == nil {
		return
	}

	goMod := []byte(fmt.Sprintf(`module github.com/eif-courses/hlabgen/experiments/out/%s

go 1.25

require github.com/gorilla/mux v1.8.1
`, appName))

	if err := os.WriteFile(goModPath, goMod, 0o644); err != nil {
		log.Printf("⚠️  Failed to create go.mod: %v", err)
		return
	}
	fmt.Println("📄 Created go.mod")
}

func fixParseIDTypeMismatch(projectDir string) {
	fmt.Println("🔧 Fixing parseID type mismatches...")

	fixCount := 0
	filepath.WalkDir(projectDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return nil
		}

		original := string(content)
		fixed := original

		fixed = strings.ReplaceAll(fixed, "parseID(id)", "id")

		parseIDFunc := `

func parseID(s string) int {
	id, _ := strconv.Atoi(s)
	return id
}`
		fixed = strings.ReplaceAll(fixed, parseIDFunc, "")
		fixed = strings.ReplaceAll(fixed, "func parseID(s string) int { id, _ := strconv.Atoi(s); return id }", "")

		if fixed != original {
			os.WriteFile(path, []byte(fixed), 0o644)
			fixCount++
			relPath, _ := filepath.Rel(projectDir, path)
			fmt.Printf("  ✅ Fixed parseID in %s\n", filepath.Base(relPath))
		}

		return nil
	})

	if fixCount > 0 {
		fmt.Printf("✅ Fixed parseID mismatches in %d file(s)\n", fixCount)
	}
}

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

// ✅ FIX 3: Enhanced fixImportsToModule with smart module detection
func fixImportsToModule(projectDir string) {
	goMod := filepath.Join(projectDir, "go.mod")
	f, err := os.Open(goMod)
	if err != nil {
		log.Printf("⚠️  No go.mod found in %s (skipping import fix)", projectDir)
		return
	}
	defer f.Close()

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

		// ✅ FIX: Replace MODULENAME placeholder with actual module name
		newContent = strings.ReplaceAll(newContent, "MODULENAME", moduleName)

		// Get app name from module (last part after /)
		parts := strings.Split(moduleName, "/")
		appName := parts[len(parts)-1]

		// Fix various import patterns
		newContent = strings.ReplaceAll(newContent,
			fmt.Sprintf(`"%s/internal/`, appName),
			fmt.Sprintf(`"%s/internal/`, moduleName))

		wrongPatterns := []string{
			`"github.com/eif-courses/hlabgen/internal/`,
			`"github.com/yourusername/` + appName + `/internal/`,
			`"github.com/yourusername/` + appName + `/`,
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

		// ✅ Fix remaining placeholder imports using regex
		re := regexp.MustCompile(`"[a-zA-Z_]+/internal/`)
		if re.MatchString(newContent) && !strings.Contains(newContent, moduleName+"/internal/") {
			newContent = re.ReplaceAllString(newContent, fmt.Sprintf(`"%s/internal/`, moduleName))
		}

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

func getModelName() string {
	model := os.Getenv("OPENAI_MODEL")
	if model == "" {
		return "gpt-4o-mini"
	}
	return model
}

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
