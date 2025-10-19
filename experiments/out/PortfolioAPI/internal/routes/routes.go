package routes

import (
	"PortfolioAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	// Project routes
	r.HandleFunc("/projects", handlers.CreateProject).Methods("POST")
	r.HandleFunc("/projects", handlers.GetProjects).Methods("GET")
	r.HandleFunc("/projects/{id}", handlers.GetProject).Methods("GET")
	r.HandleFunc("/projects/{id}", handlers.UpdateProject).Methods("PUT")
	r.HandleFunc("/projects/{id}", handlers.DeleteProject).Methods("DELETE")

	// Client routes
	r.HandleFunc("/clients", handlers.CreateClient).Methods("POST")
	r.HandleFunc("/clients", handlers.GetClients).Methods("GET")
	r.HandleFunc("/clients/{id}", handlers.GetClient).Methods("GET")
	r.HandleFunc("/clients/{id}", handlers.UpdateClient).Methods("PUT")
	r.HandleFunc("/clients/{id}", handlers.DeleteClient).Methods("DELETE")

	// Invoice routes
	r.HandleFunc("/invoices", handlers.CreateInvoice).Methods("POST")
	r.HandleFunc("/invoices", handlers.GetInvoices).Methods("GET")
	r.HandleFunc("/invoices/{id}", handlers.GetInvoice).Methods("GET")
	r.HandleFunc("/invoices/{id}", handlers.UpdateInvoice).Methods("PUT")
	r.HandleFunc("/invoices/{id}", handlers.DeleteInvoice).Methods("DELETE")
}
