package ml

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/anthropics/anthropic-sdk-go"
)

package ml

import (
"context"
"encoding/json"
"fmt"
"os"

anthropic "github.com/anthropics/anthropic-sdk-go"
)

func GenerateWithClaude(s Schema) ([]GenFile, GenerationMetrics, error) {
	metrics := GenerationMetrics{StartTime: time.Now()}

	apiKey := os.Getenv("ANTHROPIC_API_KEY")
	if apiKey == "" {
		metrics.ErrorMessage = "ANTHROPIC_API_KEY not set"
		return nil, metrics, errors.New("ANTHROPIC_API_KEY not set")
	}

	client := anthropic.NewClient(apiKey)
	prompt := BuildPrompt(s)

	ctx := context.Background()

	resp, err := client.Messages.Create(ctx, anthropic.MessageCreateParams{
		Model: anthropic.String("claude-3-7-sonnet-20250219"),
		Messages: []anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock(prompt)),
		},
		MaxTokens: anthropic.Int(4096),
	})

	if err != nil {
		metrics.ErrorMessage = fmt.Sprintf("Claude API error: %v", err)
		return nil, metrics, err
	}

	content := resp.Content[0].Text
	files, err := tryParseModelOutput(content)

	metrics.EndTime = time.Now()
	metrics.Duration = metrics.EndTime.Sub(metrics.StartTime)

	// ... rest of metrics logic same as Generate()

	return files, metrics, nil
}