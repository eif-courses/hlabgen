package rules

import (
	"fmt"
	"strings"
)

// ComplexityScore analyzes API features and assigns a score
type ComplexityScore struct {
	Score       int
	NeedML      bool
	Strategy    string
	Features    map[string]int
	Description string
}

// AnalyzeComplexity scores the API based on features
func AnalyzeComplexity(features []string) *ComplexityScore {
	score := &ComplexityScore{
		Features: make(map[string]int),
		Score:    0,
	}

	if len(features) == 0 {
		score.Strategy = "RULES_PRIMARY"
		score.Description = "No business logic features detected"
		return score
	}

	for _, feature := range features {
		lower := strings.ToLower(feature)

		// Calculation-heavy features (HIGH priority for ML)
		if strings.Contains(lower, "discount") {
			score.Score += 3
			score.Features["discount"] = 3
			score.NeedML = true
		}
		if strings.Contains(lower, "tax") {
			score.Score += 3
			score.Features["tax"] = 3
			score.NeedML = true
		}
		if strings.Contains(lower, "pricing") || strings.Contains(lower, "price") {
			score.Score += 3
			score.Features["pricing"] = 3
			score.NeedML = true
		}

		// State machines (HIGH priority for ML)
		if strings.Contains(lower, "workflow") {
			score.Score += 3
			score.Features["workflow"] = 3
			score.NeedML = true
		}
		if strings.Contains(lower, "transition") {
			score.Score += 3
			score.Features["transition"] = 3
			score.NeedML = true
		}
		if strings.Contains(lower, "state") || strings.Contains(lower, "status") {
			score.Score += 3
			score.Features["state"] = 3
			score.NeedML = true
		}

		// Validation rules (MEDIUM - rules can do, ML can improve)
		if strings.Contains(lower, "validate") {
			score.Score += 2
			score.Features["validate"] = 2
		}
		if strings.Contains(lower, "rule") {
			score.Score += 2
			score.Features["rule"] = 2
		}
		if strings.Contains(lower, "required") || strings.Contains(lower, "constraint") {
			score.Score += 2
			score.Features["constraint"] = 2
		}

		// Auth (MEDIUM)
		if strings.Contains(lower, "auth") || strings.Contains(lower, "permission") {
			score.Score += 2
			score.Features["auth"] = 2
		}

		// Search/Filtering (LOW - rules can handle mostly)
		if strings.Contains(lower, "search") || strings.Contains(lower, "filter") {
			score.Score += 1
			score.Features["search"] = 1
		}
	}

	// Strategy decision based on score
	if score.Score > 7 {
		score.Strategy = "ML_PRIMARY"
		score.Description = "Heavy business logic - ML generates logic, rules validate"
	} else if score.Score > 4 {
		score.Strategy = "HYBRID_BALANCED"
		score.Description = "Mixed complexity - split responsibility cleanly"
	} else {
		score.Strategy = "RULES_PRIMARY"
		score.Description = "Simple CRUD - rules sufficient"
	}

	return score
}

// GetStrategy returns the recommended strategy
func (c *ComplexityScore) GetStrategy() string {
	return c.Strategy
}

// DebugInfo returns detailed scoring information
func (c *ComplexityScore) DebugInfo() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("📊 Complexity Analysis\n"))
	sb.WriteString(fmt.Sprintf("  Score: %d/10\n", c.Score))
	sb.WriteString(fmt.Sprintf("  Strategy: %s\n", c.Strategy))
	sb.WriteString(fmt.Sprintf("  Description: %s\n", c.Description))
	sb.WriteString(fmt.Sprintf("  Features Detected:\n"))
	for feat, weight := range c.Features {
		sb.WriteString(fmt.Sprintf("    - %s (weight: %d)\n", feat, weight))
	}
	return sb.String()
}
