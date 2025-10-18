package routes

import (
	"InventoryAPI/internal/handlers"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	r.POST("/products", handlers.CreateProduct)
	r.GET("/products", handlers.GetProducts)
	r.PUT("/products/:id", handlers.UpdateProduct)
	r.DELETE("/products/:id", handlers.DeleteProduct)

	r.POST("/warehouses", handlers.CreateWarehouse)
	r.GET("/warehouses", handlers.GetWarehouses)
	r.PUT("/warehouses/:id", handlers.UpdateWarehouse)
	r.DELETE("/warehouses/:id", handlers.DeleteWarehouse)

	r.POST("/stocks", handlers.CreateStock)
	r.GET("/stocks", handlers.GetStocks)
	r.PUT("/stocks/:id", handlers.UpdateStock)
	r.DELETE("/stocks/:id", handlers.DeleteStock)

	r.POST("/suppliers", handlers.CreateSupplier)
	r.GET("/suppliers", handlers.GetSuppliers)
	r.PUT("/suppliers/:id", handlers.UpdateSupplier)
	r.DELETE("/suppliers/:id", handlers.DeleteSupplier)

	r.POST("/orders", handlers.CreateOrder)
	r.GET("/orders", handlers.GetOrders)
	r.PUT("/orders/:id", handlers.UpdateOrder)
	r.DELETE("/orders/:id", handlers.DeleteOrder)
}
