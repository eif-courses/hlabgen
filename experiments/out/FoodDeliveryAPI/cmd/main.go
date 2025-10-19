// Auto-generated scaffold for FoodDeliveryAPI (2025-10-19 12:27:43)
package main

import (
	"FoodDeliveryAPI/internal/routes"
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
