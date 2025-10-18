package routes

import (
	"FinanceAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register() {
	router.HandleFunc("/accounts", handlers.CreateAccount).Methods("POST")
	router.HandleFunc("/accounts", handlers.GetAccounts).Methods("GET")
	router.HandleFunc("/accounts/{id}", handlers.UpdateAccount).Methods("PUT")
	router.HandleFunc("/accounts/{id}", handlers.DeleteAccount).Methods("DELETE")
	router.HandleFunc("/transactions", handlers.CreateTransaction).Methods("POST")
	router.HandleFunc("/transactions", handlers.GetTransactions).Methods("GET")
	router.HandleFunc("/transactions/{id}", handlers.UpdateTransaction).Methods("PUT")
	router.HandleFunc("/transactions/{id}", handlers.DeleteTransaction).Methods("DELETE")
	router.HandleFunc("/invoices", handlers.CreateInvoice).Methods("POST")
	router.HandleFunc("/invoices", handlers.GetInvoices).Methods("GET")
	router.HandleFunc("/invoices/{id}", handlers.UpdateInvoice).Methods("PUT")
	router.HandleFunc("/invoices/{id}", handlers.DeleteInvoice).Methods("DELETE")
}
