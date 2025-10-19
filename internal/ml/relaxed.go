package ml

import (
	"log"
	"path/filepath"
	"regexp"
)

// GenerateRelaxed retries ML generation but also tries to clean up malformed JSON
// fragments if the first generation attempt fails.
func GenerateRelaxed(s Schema) ([]GenFile, GenerationMetrics, error) {
	log.Println("🪄 Using relaxed ML generation mode (cleaning JSON output)...")

	// --- 1️⃣ Try normal generation first ---
	files, metrics, err := Generate(s)
	if err == nil {
		log.Println("✅ Normal ML generation succeeded — no relaxed mode needed.")
		return files, metrics, nil
	}

	log.Printf("🧹 Cleaning malformed JSON output after error: %v\n", err)

	// --- 2️⃣ Try to extract possible valid JSON structure from the error ---
	re := regexp.MustCompile(`(?s)(\[.*\]|\{.*\})`)
	matches := re.FindStringSubmatch(err.Error())
	if len(matches) > 0 {
		log.Println("✅ Extracted possible valid JSON structure, attempting re-parse...")
		jsonCandidate := matches[0]

		// Try parsing recovered JSON directly
		recoveredFiles, parseErr := tryParseModelOutput(jsonCandidate)
		if parseErr == nil {
			log.Println("✅ Successfully recovered valid JSON output after cleanup.")
			metrics.FinalSuccess = true
			saveMetrics(s.AppName, metrics, filepath.Join("experiments", s.AppName, "gen_metrics_relaxed.json"))
			return recoveredFiles, metrics, nil
		}
		log.Printf("⚠️  JSON re-parse failed: %v\n", parseErr)
	}

	// --- 3️⃣ Retry full generation in relaxed mode ---
	log.Println("🔁 Retrying ML generation in relaxed mode (second API call)...")
	files, metrics, retryErr := Generate(s)
	if retryErr != nil {
		log.Printf("❌ Relaxed ML generation still failed: %v\n", retryErr)
		metrics.FinalSuccess = false
		metrics.ErrorMessage = retryErr.Error()

		// Save metrics specifically for relaxed run
		saveMetrics(s.AppName, metrics, filepath.Join("experiments", s.AppName, "gen_metrics_relaxed.json"))
		return files, metrics, retryErr
	}

	log.Println("✅ Relaxed ML generation succeeded.")
	saveMetrics(s.AppName, metrics, filepath.Join("experiments", s.AppName, "gen_metrics_relaxed.json"))
	return files, metrics, nil
}
