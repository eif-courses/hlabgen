package assemble

import (
	"os"
	"path/filepath"
)

type File struct {
	Filename string
	Content  string
}

func WriteMany(base string, files []File) error {

	for _, file := range files {
		p := filepath.Join(base, file.Filename)
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			return err
		}
		if err := os.WriteFile(p, []byte(file.Content), 0o644); err != nil {
			return err
		}
	}
	return nil
}
