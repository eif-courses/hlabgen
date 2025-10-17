package input

type Schema struct {
	AppName    string   `json:"app_name"`
	Database   string   `json:"database"`   // "PostgreSQL" | "SQLite"
	Difficulty string   `json:"difficulty"` // "Beginner" | "Intermediate" | "Advanced"
	Entities   []string `json:"entities"`
	Features   []string `json:"features"`
	Objectives []string `json:"objectives"`
	APIPattern string   `json:"api_pattern"` // "REST"
}
