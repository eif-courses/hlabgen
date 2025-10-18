package main

import (
	"github.com/eif-courses/LibraryAPI/internal/routes"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	routes.Register(mux)
	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
