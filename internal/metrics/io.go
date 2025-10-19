package metrics

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	mlinternal "github.com/eif-courses/hlabgen/internal/ml"
)

// SaveResult saves validation/build metrics.
func SaveResult(outDir string, m Result) error {
	filename := fmt.Sprintf("metrics_%s.json", timestamp())
	path := filepath.Join(outDir, filename)
	return writeJSON(path, m)
}

// SaveMLMetrics saves machine learning generation metrics.
func SaveMLMetrics(outDir string, m mlinternal.GenerationMetrics) error {
	filename := fmt.Sprintf("gen_metrics_%s.json", timestamp())
	path := filepath.Join(outDir, filename)
	return writeJSON(path, m)
}

// SaveCombinedMetrics saves both ML and build metrics in one merged file.
func SaveCombinedMetrics(outDir string, build Result, gen mlinternal.GenerationMetrics) error {
	data := struct {
		Timestamp  string                       `json:"timestamp"`
		Build      Result                       `json:"build"`
		Generation mlinternal.GenerationMetrics `json:"generation"`
	}{
		Timestamp:  time.Now().Format(time.RFC3339),
		Build:      build,
		Generation: gen,
	}

	filename := fmt.Sprintf("combined_metrics_%s.json", timestamp())
	path := filepath.Join(outDir, filename)
	return writeJSON(path, data)
}

// writeJSON handles safe directory creation and writing.
func writeJSON(path string, v any) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}
	return os.WriteFile(path, b, 0o644)
}

// timestamp returns a compact timestamp for filenames (no colons).
func timestamp() string {
	return time.Now().Format("2006-01-02_15-04-05")
}
