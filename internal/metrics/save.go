package metrics

import (
	"encoding/json"
	"os"
)

func Save(path string, m Result) error {
	b, _ := json.MarshalIndent(m, "", " ")
	return os.WriteFile(path, b, 0o644)
}
