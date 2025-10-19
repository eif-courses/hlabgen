package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/eif-courses/hlabgen/internal/report"
)

func main() {
	fmt.Println("📊 Generating experiment summary...")

	if err := report.GenerateSummaryJSONReport(); err != nil {
		log.Fatalf("❌ Report generation failed: %v", err)
	}

	results, _ := report.CollectAllExperiments("experiments")

	// Count statistics
	success := 0
	total := len(results)
	for _, r := range results {
		if r.FinalSuccess {
			success++
		}
	}

	fmt.Printf("✅ %d/%d experiments succeeded (%.1f%%)\n", success, total, float64(success)/float64(total)*100)
	fmt.Printf("📄 Markdown summary saved at %s\n", filepath.Join("experiments", "logs", "results.md"))
}
