package rules

import (
	"fmt"
	"os"
	"path/filepath"
)

type File struct {
	Name string
	Code []byte
}

func Scaffold(outDir string, appName string) ([]File, error) {
	dirs := []string{
		filepath.Join(outDir, "cmd"),
		filepath.Join(outDir, "internal", "handlers"),
		filepath.Join(outDir, "internal", "models"),
		filepath.Join(outDir, "internal", "routes"),
		filepath.Join(outDir, "tests"),
	}
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return nil, err
		}
	}

	mainGo := []byte(`package main

import (
	"log"
	"net/http"
	"github.com/eif-courses/` + appName + `/internal/routes"
)

func main() {
	mux := http.NewServeMux()
	routes.Register(mux)
	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
`)

	routesGo := []byte(`package routes

import "net/http"

func Register(mux *http.ServeMux) {
	// TODO: handlers wiring is appended by ML layer
}
`)
	goMod := []byte(fmt.Sprintf("module github.com/eif-courses/%s\n\ngo 1.25\n", appName))

	files := []File{
		{Name: filepath.Join(outDir, "cmd", "main.go"), Code: mainGo},
		{Name: filepath.Join(outDir, "internal", "routes", "routes.go"), Code: routesGo},
		{Name: filepath.Join(outDir, "go.mod"), Code: goMod},
	}
	for _, f := range files {
		if err := os.WriteFile(f.Name, f.Code, 0o644); err != nil {
			return nil, err
		}
	}
	return files, nil

}
