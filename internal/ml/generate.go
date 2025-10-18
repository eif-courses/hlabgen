package ml

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	openai "github.com/sashabaranov/go-openai"
)

// GenFile represents one generated code file.

// GenerationMetrics stores run information for analysis and paper results.
type GenerationMetrics struct {
	StartTime      time.Time
	EndTime        time.Time
	Duration       time.Duration
	PrimarySuccess bool
	RepairAttempts int
	FinalSuccess   bool
	ErrorMessage   string
	RuleFixes      int // NEW: counts fixes applied in WriteMany
}

// Generate creates Go code scaffolds using ML and repairs malformed output automatically.
func Generate(s Schema) ([]GenFile, GenerationMetrics, error) {
	metrics := GenerationMetrics{StartTime: time.Now()}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		metrics.ErrorMessage = "OPENAI_API_KEY not set"
		return nil, metrics, errors.New("OPENAI_API_KEY environment variable not set")
	}

	client := openai.NewClient(apiKey)
	prompt := BuildPrompt(s)

	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	startQuery := time.Now()
	resp, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT4oMini,
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleUser, Content: prompt},
		},
		Temperature: 0.2,
	})
	if err != nil {
		metrics.ErrorMessage = fmt.Sprintf("initial query failed: %v", err)
		return nil, metrics, err
	}
	fmt.Printf("⏱️  Initial model call took %v\n", time.Since(startQuery))

	if len(resp.Choices) == 0 {
		metrics.ErrorMessage = "no response choices from model"
		return nil, metrics, errors.New(metrics.ErrorMessage)
	}

	content := strings.TrimSpace(resp.Choices[0].Message.Content)

	files, err := tryParseModelOutput(content)
	if err != nil {
		metrics.PrimarySuccess = false
		metrics.RepairAttempts++

		fmt.Println("⚠️ Primary parse failed, retrying with JSON repair prompt...")

		startRepair := time.Now()
		fixedContent, repairErr := repairModelOutput(client, ctx, content)
		if repairErr != nil {
			metrics.ErrorMessage = fmt.Sprintf("repair model error: %v", repairErr)
			metrics.EndTime = time.Now()
			metrics.Duration = metrics.EndTime.Sub(metrics.StartTime)
			saveMetrics(metrics, "experiments/logs/metrics.csv")
			return nil, metrics, fmt.Errorf("failed to repair model output: %w", repairErr)
		}
		fmt.Printf("🔁 Repair call took %v\n", time.Since(startRepair))

		files, err = tryParseModelOutput(fixedContent)
		if err != nil {
			fmt.Println("--- CLEANED JSON ---")
			fmt.Println(fixedContent)
			fmt.Println("--- END ---")
			metrics.FinalSuccess = false
			metrics.ErrorMessage = fmt.Sprintf("failed to parse repaired output: %v", err)
			metrics.EndTime = time.Now()
			metrics.Duration = metrics.EndTime.Sub(metrics.StartTime)
			saveMetrics(metrics, "experiments/logs/metrics.csv")
			return nil, metrics, fmt.Errorf(metrics.ErrorMessage)
		}

		metrics.FinalSuccess = true
	} else {
		metrics.PrimarySuccess = true
		metrics.FinalSuccess = true
	}

	metrics.EndTime = time.Now()
	metrics.Duration = metrics.EndTime.Sub(metrics.StartTime)

	fmt.Printf("\n📊 Generation Summary:\n")
	fmt.Printf("  • Duration: %v\n", metrics.Duration)
	fmt.Printf("  • Primary Success: %v\n", metrics.PrimarySuccess)
	fmt.Printf("  • Repair Attempts: %d\n", metrics.RepairAttempts)
	fmt.Printf("  • Final Success: %v\n", metrics.FinalSuccess)
	fmt.Printf("  • Error: %s\n", metrics.ErrorMessage)

	saveMetrics(metrics, "experiments/logs/metrics.csv")

	// 🧩 Merge duplicate structs
	files = deduplicateAndMergeStructs(files)

	// 🔄 Merge duplicate handlers
	files = deduplicateHandlers(files)

	// 🧹 Remove duplicate helper functions
	files = deduplicateHelperFuncs(files)

	return files, metrics, nil
}

// --- Parsing and Repair ---

func tryParseModelOutput(content string) ([]GenFile, error) {
	originalContent := content

	// Remove markdown code blocks
	content = strings.TrimSpace(content)
	content = strings.TrimPrefix(content, "```json")
	content = strings.TrimPrefix(content, "```")
	content = strings.TrimSuffix(content, "```")
	content = strings.TrimSpace(content)

	// Find JSON array boundaries
	startIdx := strings.Index(content, "[")
	endIdx := strings.LastIndex(content, "]")

	if startIdx == -1 || endIdx == -1 || startIdx >= endIdx {
		fmt.Println("❌ Could not find JSON array")
		fmt.Println("First 300 chars:", originalContent[:min(300, len(originalContent))])
		return nil, errors.New("no valid JSON array found")
	}

	// Extract JSON
	jsonStr := content[startIdx : endIdx+1]

	var files []GenFile
	if err := json.Unmarshal([]byte(jsonStr), &files); err != nil {
		fmt.Printf("❌ JSON error: %v\n", err)
		return nil, fmt.Errorf("json parse error: %w", err)
	}

	fmt.Printf("✅ Parsed %d files\n", len(files))
	return files, nil
}

//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}

func repairModelOutput(client *openai.Client, ctx context.Context, broken string) (string, error) {
	repairPrompt := fmt.Sprintf(`The following JSON output is invalid. 
Fix it so that it becomes valid JSON array of objects.
Return ONLY the corrected JSON (no explanations, no markdown).

Input:
%s`, broken)

	resp, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT4oMini,
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleUser, Content: repairPrompt},
		},
		Temperature: 0.0,
	})
	if err != nil {
		return "", err
	}
	if len(resp.Choices) == 0 {
		return "", errors.New("no repair response from model")
	}

	content := strings.TrimSpace(resp.Choices[0].Message.Content)
	content = strings.TrimPrefix(content, "```json")
	content = strings.TrimPrefix(content, "```")
	content = strings.TrimSuffix(content, "```")

	return strings.TrimSpace(content), nil
}

// --- Metrics ---

func saveMetrics(m GenerationMetrics, path string) {
	os.MkdirAll(filepath.Dir(path), 0o755)
	line := fmt.Sprintf("%s,%t,%d,%t,%s,%v\n",
		m.StartTime.Format(time.RFC3339),
		m.PrimarySuccess,
		m.RepairAttempts,
		m.FinalSuccess,
		m.ErrorMessage,
		m.Duration,
		m.RuleFixes,
	)
	f, _ := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	defer f.Close()
	f.WriteString(line)
}

// --- Deduplication ---

func deduplicateAndMergeStructs(files []GenFile) []GenFile {
	structPattern := regexp.MustCompile(`(?ms)^type\s+(\w+)\s+struct\s*{(.*?)}\s*$`)
	structs := make(map[string]string)
	for _, f := range files {
		matches := structPattern.FindAllStringSubmatch(f.Code, -1)
		for _, m := range matches {
			name, body := m[1], strings.TrimSpace(m[2])
			if existing, ok := structs[name]; ok {
				existingLines := strings.Split(existing, "\n")
				bodyLines := strings.Split(body, "\n")
				lineSet := make(map[string]bool)
				for _, l := range existingLines {
					lineSet[strings.TrimSpace(l)] = true
				}
				for _, l := range bodyLines {
					l = strings.TrimSpace(l)
					if l != "" && !lineSet[l] {
						existingLines = append(existingLines, l)
						lineSet[l] = true
					}
				}
				structs[name] = strings.Join(existingLines, "\n")
				fmt.Printf("🧩 Merged duplicate struct: %s\n", name)
			} else {
				structs[name] = body
			}
		}
	}

	for i := range files {
		files[i].Code = structPattern.ReplaceAllStringFunc(files[i].Code, func(s string) string {
			m := structPattern.FindStringSubmatch(s)
			if m == nil {
				return s
			}
			name := m[1]
			if merged, ok := structs[name]; ok {
				return fmt.Sprintf("type %s struct {\n%s\n}", name, merged)
			}
			return s
		})
	}
	return files
}

func deduplicateHandlers(files []GenFile) []GenFile {
	funcPattern := regexp.MustCompile(`(?ms)^func\s+(\w+)\s*\([^)]*\)\s*{(.*?)}\s*$`)
	handlers := make(map[string]string)

	for i := range files {
		matches := funcPattern.FindAllStringSubmatch(files[i].Code, -1)
		if matches == nil {
			continue
		}
		for _, m := range matches {
			name, body := m[1], strings.TrimSpace(m[2])
			if prev, ok := handlers[name]; ok {
				prevLines := strings.Split(prev, "\n")
				bodyLines := strings.Split(body, "\n")
				lineSet := make(map[string]bool)
				for _, l := range prevLines {
					lineSet[strings.TrimSpace(l)] = true
				}
				for _, l := range bodyLines {
					l = strings.TrimSpace(l)
					if l != "" && !lineSet[l] {
						prevLines = append(prevLines, l)
						lineSet[l] = true
					}
				}
				handlers[name] = strings.Join(prevLines, "\n")
				fmt.Printf("🔄 Merged duplicate handler: %s\n", name)
			} else {
				handlers[name] = body
			}
		}
	}

	for i := range files {
		files[i].Code = funcPattern.ReplaceAllStringFunc(files[i].Code, func(s string) string {
			m := funcPattern.FindStringSubmatch(s)
			if m == nil {
				return s
			}
			name := m[1]
			if merged, ok := handlers[name]; ok {
				return fmt.Sprintf("func %s() {\n%s\n}", name, merged)
			}
			return s
		})
	}
	return files
}

func deduplicateHelperFuncs(files []GenFile) []GenFile {
	funcPattern := regexp.MustCompile(`(?m)^func\s+(FromJSON|ToJSON|ToString)\s*\(.*\)\s*{`)
	seen := make(map[string]bool)

	for i := range files {
		lines := strings.Split(files[i].Code, "\n")
		newLines := []string{}
		skip := false

		for _, line := range lines {
			if m := funcPattern.FindStringSubmatch(line); m != nil {
				name := m[1]
				if seen[name] {
					fmt.Printf("🧹 Removed duplicate helper: %s\n", name)
					skip = true
					continue
				}
				seen[name] = true
			}
			if skip && strings.TrimSpace(line) == "}" {
				skip = false
				continue
			}
			if !skip {
				newLines = append(newLines, line)
			}
		}
		files[i].Code = strings.Join(newLines, "\n")
	}
	return files
}
