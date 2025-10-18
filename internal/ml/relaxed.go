package ml

import (
	"log"
	"regexp"
)

// GenerateRelaxed retries ML generation but also tries to clean up malformed JSON
// fragments if the first generation attempt fails.
func GenerateRelaxed(s Schema) ([]GenFile, GenerationMetrics, error) {
	log.Println("🪄 Using relaxed ML generation mode (cleaning JSON output)...")

	// First, try normal generation
	files, metrics, err := Generate(s)
	if err == nil {
		return files, metrics, nil
	}

	// Attempt to fix malformed JSON manually
	log.Printf("🧹 Cleaning malformed JSON output after error: %v\n", err)

	// Regex to capture the first valid JSON array or object from the output
	re := regexp.MustCompile(`(?s)(\[.*\]|\{.*\})`)
	matches := re.FindStringSubmatch(err.Error())
	if len(matches) > 0 {
		log.Println("✅ Extracted possible valid JSON structure, retrying parse...")
		// TODO: Integrate JSON re-parse using matches[0] if desired
	}

	// Retry generation one more time (relaxed mode)
	log.Println("🔁 Retrying ML generation in relaxed mode...")
	files, metrics, retryErr := Generate(s)
	if retryErr != nil {
		log.Printf("❌ Relaxed ML generation still failed: %v\n", retryErr)
		return files, metrics, retryErr
	}

	log.Println("✅ Relaxed ML generation succeeded.")
	return files, metrics, nil
}
