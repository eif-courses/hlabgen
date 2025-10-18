package routes

import (
	"ShopAPI/internal/handlers"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	r.POST("/products", handlers.CreateProduct)
	r.GET("/products/:id", handlers.GetProduct)
	r.PUT("/products/:id", handlers.UpdateProduct)
	r.DELETE("/products/:id", handlers.DeleteProduct)
	r.GET("/products", handlers.ListProducts)

	r.POST("/orders", handlers.CreateOrder)
	r.GET("/orders/:id", handlers.GetOrder)
	r.PUT("/orders/:id", handlers.UpdateOrder)
	r.DELETE("/orders/:id", handlers.DeleteOrder)
	r.GET("/orders", handlers.ListOrders)

	r.POST("/customers", handlers.CreateCustomer)
	r.GET("/customers/:id", handlers.GetCustomer)
	r.PUT("/customers/:id", handlers.UpdateCustomer)
	r.DELETE("/customers/:id", handlers.DeleteCustomer)
	r.GET("/customers", handlers.ListCustomers)

	r.POST("/carts", handlers.CreateCart)
	r.GET("/carts/:id", handlers.GetCart)
	r.PUT("/carts/:id", handlers.UpdateCart)
	r.DELETE("/carts/:id", handlers.DeleteCart)
	r.GET("/carts", handlers.ListCarts)

	r.POST("/payments", handlers.CreatePayment)
	r.GET("/payments/:id", handlers.GetPayment)
	r.PUT("/payments/:id", handlers.UpdatePayment)
	r.DELETE("/payments/:id", handlers.DeletePayment)
	r.GET("/payments", handlers.ListPayments)
}
