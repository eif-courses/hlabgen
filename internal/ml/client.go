package ml

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	openai "github.com/sashabaranov/go-openai"
)

func Generate(s Schema) ([]GenFile, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return nil, errors.New("OPENAI_API_KEY environment variable not set")
	}

	client := openai.NewClient(apiKey)
	prompt := BuildPrompt(s)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	resp, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT4oMini, // or openai.GPT4o for better quality
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
		Temperature: 0.2,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create chat completion: %w", err)
	}

	if len(resp.Choices) == 0 {
		return nil, errors.New("no response choices returned from OpenAI")
	}

	content := resp.Choices[0].Message.Content

	// Parse the JSON array from the response
	var files []GenFile
	if err := json.Unmarshal([]byte(content), &files); err != nil {
		return nil, fmt.Errorf("failed to parse generated files from response: %w (content: %s)", err, content)
	}

	return files, nil
}
