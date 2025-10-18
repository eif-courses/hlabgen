package routes

import (
	"ShopAPI/internal/handlers"
	"github.com/gorilla/mux"
)

func Register() {
	r.HandleFunc("/products", handlers.CreateProduct).Methods("POST")
	r.HandleFunc("/products", handlers.GetProducts).Methods("GET")
	r.HandleFunc("/products/{id}", handlers.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", handlers.DeleteProduct).Methods("DELETE")

	r.HandleFunc("/orders", handlers.CreateOrder).Methods("POST")
	r.HandleFunc("/orders", handlers.GetOrders).Methods("GET")
	r.HandleFunc("/orders/{id}", handlers.UpdateOrder).Methods("PUT")
	r.HandleFunc("/orders/{id}", handlers.DeleteOrder).Methods("DELETE")

	r.HandleFunc("/customers", handlers.CreateCustomer).Methods("POST")
	r.HandleFunc("/customers", handlers.GetCustomers).Methods("GET")
	r.HandleFunc("/customers/{id}", handlers.UpdateCustomer).Methods("PUT")
	r.HandleFunc("/customers/{id}", handlers.DeleteCustomer).Methods("DELETE")

	r.HandleFunc("/carts", handlers.CreateCart).Methods("POST")
	r.HandleFunc("/carts", handlers.GetCarts).Methods("GET")
	r.HandleFunc("/carts/{id}", handlers.UpdateCart).Methods("PUT")
	r.HandleFunc("/carts/{id}", handlers.DeleteCart).Methods("DELETE")

	r.HandleFunc("/payments", handlers.CreatePayment).Methods("POST")
	r.HandleFunc("/payments", handlers.GetPayments).Methods("GET")
	r.HandleFunc("/payments/{id}", handlers.UpdatePayment).Methods("PUT")
	r.HandleFunc("/payments/{id}", handlers.DeletePayment).Methods("DELETE")
}
