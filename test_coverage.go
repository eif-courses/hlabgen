package main

import (
	"fmt"
	"github.com/eif-courses/hlabgen/internal/report"
	"path/filepath"
)

func main() {
	// Find the gen_metrics file
	files, _ := filepath.Glob("experiments/out/LibraryAPI_rules/gen_metrics*.json")
	if len(files) == 0 {
		fmt.Println("No gen_metrics files found")
		return
	}

	fmt.Printf("Testing file: %s\n", files[0])

	result, err := report.LoadMetricsFromJSON(files[0])
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("App: %s\n", result.AppName)
	fmt.Printf("Mode: %s\n", result.Mode)
	fmt.Printf("Coverage: %.1f%%\n", result.Coverage)
	fmt.Printf("Duration: %.2f seconds\n", result.DurationSeconds)
}
