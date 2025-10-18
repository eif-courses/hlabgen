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

	fmt.Fprintf(&buf, `Generate Go REST API files as a JSON array of objects with "filename" and "code" fields.
Requirements: %s

CRITICAL IMPORT RULES - READ CAREFULLY:
- Module name is EXACTLY: %s
- ALL imports must use the LOCAL module path: "%s/internal/..."
- DO NOT use "github.com/yourusername/..." or any GitHub paths
- DO NOT use "github.com/eif-courses/..." paths
- Use ONLY gorilla/mux for routing (import "github.com/gorilla/mux" only in routes.go)
- Use standard net/http handlers: func(w http.ResponseWriter, r *http.Request)
- DO NOT use gin.Context or any gin-specific types

CORRECT import examples for this project:
✅ import "%s/internal/models"
✅ import "%s/internal/handlers"
❌ import "github.com/yourusername/%s/internal/models"
❌ import "yourapp/internal/models"

Include:
- models with JSON tags in internal/models/
- handlers with standard http.ResponseWriter signature in internal/handlers/
- routes using mux.Router in internal/routes/routes.go (append to existing Register function)
- tests in internal/handlers/ using package handlers_test
- tasks.md with 3 lab tasks

Example handler structure:
func CreateBook(w http.ResponseWriter, r *http.Request) {
    var book models.Book
    if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(book)
}

CRITICAL: Return ONLY a valid JSON array. Do NOT wrap in markdown. Start with [ and end with ].`,
		string(b), s.AppName, s.AppName, s.AppName, s.AppName, s.AppName)
	return buf.String()
}
