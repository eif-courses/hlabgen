package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/eif-courses/hlabgen/internal/assemble"
	"github.com/eif-courses/hlabgen/internal/input"
	"github.com/eif-courses/hlabgen/internal/metrics"
	mlinternal "github.com/eif-courses/hlabgen/internal/ml"
	"github.com/eif-courses/hlabgen/internal/rules"
	"github.com/eif-courses/hlabgen/internal/validate"
	"github.com/joho/godotenv"

	"log"
	"path/filepath"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	in := flag.String("input", "experiments/input/library.json", "path to schema.json")
	mode := flag.String("mode", "hybrid", "rules|ml|hybrid")
	out := flag.String("out", "experiments/out/LibraryAPI", "output directory")
	flag.Parse()

	schema, err := input.Load(*in)
	if err != nil {
		log.Fatal(err)
	}
	// 1) rule-based scaffold (always)
	if _, err := rules.Scaffold(*out, schema.AppName); err != nil {
		log.Fatal(err)
	}

	// 2) ML layer (if ml-only or hybrid)
	if *mode == "ml" || *mode == "hybrid" {
		files, err := callML(schema) // implement in M1
		if err != nil {
			log.Fatal(err)
		}
		if err := assemble.WriteMany(*out, files); err != nil {
			log.Fatal(err)
		}
	}

	// 3) validate & metrics
	m, err := validate.Run(*out)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("BuildSuccess=%v Lint=%d TestsPass=%v Coverage=%.1f%%\n", m.BuildSuccess, m.LintWarnings, m.TestsPass, m.CoveragePct)
	_ = saveMetrics(filepath.Join(*out, "metrics.json"), m) // implement shortly
}

func callML(s input.Schema) ([]assemble.File, error) {
	files, err := mlinternal.Generate(mlinternal.Schema{
		AppName: s.AppName, Database: s.Database, APIPattern: s.APIPattern,
		Difficulty: s.Difficulty, Entities: s.Entities, Features: s.Features, Objectives: s.Objectives,
	})
	if err != nil {
		return nil, err
	}
	out := make([]assemble.File, 0, len(files))
	for _, f := range files {
		out = append(out, assemble.File{Filename: f.Filename, Content: f.Code})
	}
	return out, nil
}

func saveMetrics(path string, m metrics.Result) error {
	b, _ := json.MarshalIndent(m, "", "  ")
	return os.WriteFile(path, b, 0o644)
}
