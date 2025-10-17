package input

import (
	"encoding/json"
	"os"
)

func Load(path string) (Schema, error) {
	var schema Schema
	file, err := os.ReadFile(path)
	if err != nil {
		return schema, err
	}
	err = json.Unmarshal(file, &schema)
	return schema, err
}
