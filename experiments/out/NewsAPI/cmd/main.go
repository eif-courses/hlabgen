// Auto-generated scaffold for NewsAPI (2025-10-19 12:28:03)
package main

import (
	"NewsAPI/internal/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.Register(r)
	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
