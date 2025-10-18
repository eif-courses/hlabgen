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

// GenerationMetrics stores run information for analysis and paper results.
type GenerationMetrics struct {
	StartTime      time.Time
	EndTime        time.Time
	Duration       time.Duration
	PrimarySuccess bool
	RepairAttempts int
	FinalSuccess   bool
	ErrorMessage   string
}

// Generate creates Go code scaffolds using ML and repairs malformed output automatically.
// It also tracks performance metrics for research evaluation.
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

	// 🔹 Save metrics for later research analysis
	saveMetrics(metrics, "experiments/logs/metrics.csv")

	return files, metrics, nil
}

// tryParseModelOutput cleans and parses JSON content from LLM response.
func tryParseModelOutput(content string) ([]GenFile, error) {
	content = strings.TrimPrefix(content, "```json")
	content = strings.TrimPrefix(content, "```")
	content = strings.TrimSuffix(content, "```")
	content = strings.TrimSpace(content)

	re := regexp.MustCompile(`(?s)\[\s*{.*}\s*\]`)
	matches := re.FindString(content)
	if matches == "" {
		return nil, errors.New("no valid JSON array found")
	}

	var files []GenFile
	if err := json.Unmarshal([]byte(matches), &files); err != nil {
		return nil, fmt.Errorf("json parse error: %w", err)
	}

	return files, nil
}

// repairModelOutput asks OpenAI to fix malformed JSON from a previous response.
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

// saveMetrics appends metrics to a CSV file for later research analysis.
func saveMetrics(m GenerationMetrics, path string) {
	os.MkdirAll(filepath.Dir(path), 0o755)

	line := fmt.Sprintf("%s,%t,%d,%t,%s,%v\n",
		m.StartTime.Format(time.RFC3339),
		m.PrimarySuccess,
		m.RepairAttempts,
		m.FinalSuccess,
		m.ErrorMessage,
		m.Duration,
	)

	f, _ := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	defer f.Close()
	f.WriteString(line)
}
