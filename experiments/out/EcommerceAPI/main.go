package main

import (
	"EcommerceAPI/internal/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.Register(r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
