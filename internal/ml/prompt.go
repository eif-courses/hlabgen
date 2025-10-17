package ml

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Schema struct {
	AppName    string
	Database   string
	APIPattern string
	Difficulty string
	Entities   []string
	Features   []string
	Objectives []string
}

func BuildPrompt(s Schema) string {

	var buf bytes.Buffer
	b, _ := json.Marshal(s)

	fmt.Fprintf(&buf, `Generate Go REST API files as JSON array of {"filename","code"}.
Requirements: %s
Include:
- models with JSON tags
- handlers CRUD for each entity
- routes glue (append to internal/routes/routes.go Register)
- tests in /tests
- tasks.md with 3 lab tasks (statement, hints, expected output)
Return ONLY valid JSON array.`, string(b))
	return buf.String()
}
