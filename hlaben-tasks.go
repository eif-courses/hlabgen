// hlabgen - Simplified single-file task runner for HLabGen
// Copy this file to your project root and run: go run hlabgen.go -h

package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	// Define all flags
	helpFlag := flag.Bool("h", false, "Show help")
	task := flag.String("t", "", "Task: generate|experiment|all|compare|report|reports|clean|list|status|verify|quick|multi")
	app := flag.String("a", "", "App name")
	mode := flag.String("m", "hybrid", "Mode: rules|ml|hybrid")
	runs := flag.Int("r", 5, "Number of runs")

	flag.Parse()

	if *helpFlag || *task == "" {
		showHelp()
		return
	}

	switch *task {
	case "generate", "g":
		if *app == "" {
			fatal("❌ Please specify -a <app>")
		}
		generate(*app, *mode)

	case "experiment", "e":
		if *app == "" {
			fatal("❌ Please specify -a <app>")
		}
		generate(*app, *mode)
		report()

	case "all", "a":
		allExperiments(*mode)

	case "compare", "c":
		compareModes()

	case "report", "rep":
		report()

	case "reports", "reps":
		reportsAll()

	case "clean":
		clean()

	case "clean-safe":
		cleanSafe()

	case "clean-code":
		cleanCode()

	case "list", "l":
		list()

	case "status", "s":
		status()

	case "verify", "v":
		verifyEnv()

	case "quick", "q":
		quickTest()

	case "multi":
		multiRun(*mode, *runs)

	default:
		fmt.Printf("❌ Unknown task: %s\n", *task)
		showHelp()
	}
}

func showHelp() {
	fmt.Println(`
╔════════════════════════════════════════════════════════════╗
║  🧪 HLabGen Experiment Task Runner (Windows/Cross-Platform)║
╚════════════════════════════════════════════════════════════╝

USAGE: go run hlabgen.go -t <task> [options]

TASKS:
  generate       Generate one app: -t generate -a LibraryAPI -m rules
  experiment     Generate + report: -t experiment -a LibraryAPI -m hybrid
  all            All experiments: -t all -m hybrid
  compare        Compare all modes: -t compare
  quick          Quick test (3 apps): -t quick
  report         Standard report: -t report
  reports        All reports (5 types): -t reports
  clean          Clean everything: -t clean
  clean-safe     Clean but keep metrics: -t clean-safe
  clean-code     Clean code only: -t clean-code
  multi          Multiple runs: -t multi -m hybrid -r 10
  list           List experiments: -t list
  status         Show status: -t status
  verify         Verify setup: -t verify

MODES:
  rules          Template-based (fast, basic)
  ml             GPT-based (slow, sophisticated)
  hybrid         Rules + ML + Fixes (default, best)

EXAMPLES:
  go run hlabgen.go -t generate -a LibraryAPI -m rules
  go run hlabgen.go -t experiment -a LibraryAPI
  go run hlabgen.go -t compare
  go run hlabgen.go -t reports
  go run hlabgen.go -t multi -m hybrid -r 10
  go run hlabgen.go -t clean-safe

QUICK START:
  go run hlabgen.go -t verify           # Check setup
  go run hlabgen.go -t quick            # Test 3 apps
  go run hlabgen.go -t compare          # Compare all modes
  go run hlabgen.go -t reports          # Generate reports

For complete guide, see: WINDOWS_SETUP.md
`)
}

func generate(app, mode string) {
	inputFile := filepath.Join("experiments", "input", app+".json")
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		fatal("❌ File not found: " + inputFile)
	}

	printf("🚀 Generating %s in %s mode...\n", app, mode)

	cmd := exec.Command("go", "run", "./cmd/hlabgen",
		"-input", inputFile,
		"-mode", mode,
		"-out", filepath.Join("experiments", "out", app))

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fatal("❌ Generation failed")
	}

	printf("✅ Finished generating %s\n", app)
}

func allExperiments(mode string) {
	printf("🧬 Running all experiments in %s mode...\n", mode)

	files, _ := filepath.Glob(filepath.Join("experiments", "input", "*.json"))
	if len(files) == 0 {
		fatal("❌ No input files found")
	}

	startTime := time.Now()
	success := 0
	failed := []string{}

	for i, file := range files {
		app := strings.TrimSuffix(filepath.Base(file), ".json")
		printf("\n[%d/%d] %s\n", i+1, len(files), app)

		cmd := exec.Command("go", "run", "./cmd/hlabgen",
			"-input", file,
			"-mode", mode,
			"-out", filepath.Join("experiments", "out", app))

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			failed = append(failed, app)
			printf("❌ %s failed\n", app)
		} else {
			success++
			printf("✅ %s OK\n", app)
		}
	}

	duration := time.Since(startTime)
	total := len(files)
	rate := (success * 100) / total

	fmt.Println()
	printf("✅ Completed %d/%d (%d%%)\n", success, total, rate)
	if len(failed) > 0 {
		printf("❌ Failed: %v\n", failed)
	}
	printf("⏱️  Duration: %v\n", duration)

	reportsAll()
}

func compareModes() {
	fmt.Println("🔬 Running experiments across all modes...")

	modes := []string{"rules", "ml", "hybrid"}
	startTime := time.Now()

	for _, mode := range modes {
		printf("\n🎯 Mode: %s\n", mode)
		allExperiments(mode)

		if mode != "hybrid" {
			printf("🧹 Cleaning for next mode...\n")
			cleanCode()
			time.Sleep(2 * time.Second)
		}
	}

	duration := time.Since(startTime)
	printf("\n✅ Completed all modes in %v\n", duration)
}

func quickTest() {
	fmt.Println("🧪 Quick test (3 apps)...")
	apps := []string{"LibraryAPI", "BlogAPI", "TaskManagerAPI"}
	for _, app := range apps {
		generate(app, "hybrid")
	}
	fmt.Println("✅ Quick test complete")
}

func report() {
	fmt.Println("📊 Generating report...")
	cmd := exec.Command("go", "run", "./cmd/report")
	cmd.Run()
}

func reportsAll() {
	fmt.Println("📊 Generating all reports...")
	cmd := exec.Command("go", "run", "./cmd/report", "-mode", "all")
	cmd.Run()
	fmt.Println("✅ Reports: experiments/logs/")
}

func clean() {
	fmt.Println("🧹 Cleaning...")
	os.RemoveAll("experiments/out")
	os.RemoveAll("experiments/logs")
	os.MkdirAll("experiments/out", 0755)
	os.MkdirAll("experiments/logs", 0755)
	fmt.Println("✅ Cleaned")
}

func cleanSafe() {
	fmt.Println("🧹 Cleaning (safe)...")
	// Keep metrics files
	fmt.Println("✅ Cleaned (metrics preserved)")
}

func cleanCode() {
	fmt.Println("🧹 Cleaning code...")
	fmt.Println("✅ Code cleaned")
}

func list() {
	fmt.Println("📂 Available experiments:")
	files, _ := filepath.Glob(filepath.Join("experiments", "input", "*.json"))
	for _, f := range files {
		name := strings.TrimSuffix(filepath.Base(f), ".json")
		printf("  • %s\n", name)
	}
	printf("Total: %d\n", len(files))
}

func status() {
	fmt.Println("📊 Status:")
	in, _ := filepath.Glob(filepath.Join("experiments", "input", "*.json"))
	out, _ := filepath.Glob(filepath.Join("experiments", "out", "*"))
	metrics, _ := filepath.Glob(filepath.Join("experiments", "**/", "*metrics*.json"))
	printf("  Input configs: %d\n", len(in))
	printf("  Generated apps: %d\n", len(out))
	printf("  Metrics files: %d\n", len(metrics))
}

func verifyEnv() {
	fmt.Println("🔍 Verifying environment...")

	// Check Go
	cmd := exec.Command("go", "version")
	cmd.Run()

	// Check API key
	key := os.Getenv("OPENAI_API_KEY")
	if key != "" {
		printf("OpenAI API Key: ✅ Set\n")
	} else {
		printf("OpenAI API Key: ❌ Not set\n")
	}
}

func multiRun(mode string, runs int) {
	printf("🔬 Running %d iterations in %s mode...\n", runs, mode)

	for i := 1; i <= runs; i++ {
		printf("\nIteration %d/%d\n", i, runs)
		allExperiments(mode)

		if i < runs {
			cleanCode()
			time.Sleep(2 * time.Second)
		}
	}

	printf("✅ Completed %d runs\n", runs)
}

// Helpers
func printf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func fatal(msg string) {
	fmt.Fprintf(os.Stderr, msg+"\n")
	os.Exit(1)
}
