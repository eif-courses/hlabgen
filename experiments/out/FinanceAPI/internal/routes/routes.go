package routes

import (
	"FinanceAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register() {
	r.HandleFunc("/accounts", handlers.CreateAccount).Methods("POST")
	r.HandleFunc("/accounts", handlers.GetAccounts).Methods("GET")
	r.HandleFunc("/accounts/{id}", handlers.UpdateAccount).Methods("PUT")
	r.HandleFunc("/accounts/{id}", handlers.DeleteAccount).Methods("DELETE")
	r.HandleFunc("/transactions", handlers.CreateTransaction).Methods("POST")
	r.HandleFunc("/transactions", handlers.GetTransactions).Methods("GET")
	r.HandleFunc("/transactions/{id}", handlers.UpdateTransaction).Methods("PUT")
	r.HandleFunc("/transactions/{id}", handlers.DeleteTransaction).Methods("DELETE")
	r.HandleFunc("/invoices", handlers.CreateInvoice).Methods("POST")
	r.HandleFunc("/invoices", handlers.GetInvoices).Methods("GET")
	r.HandleFunc("/invoices/{id}", handlers.UpdateInvoice).Methods("PUT")
	r.HandleFunc("/invoices/{id}", handlers.DeleteInvoice).Methods("DELETE")
}
