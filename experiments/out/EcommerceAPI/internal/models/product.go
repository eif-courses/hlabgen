package models

// Product represents a product in the e-commerce system.
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

// User represents a user in the e-commerce system.
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Cart represents a user's shopping cart.
type Cart struct {
	ID     int       `json:"id"`
	UserID int       `json:"user_id"`
	Items  []Product `json:"items"`
}

// Order represents a user's order.
type Order struct {
	ID       int       `json:"id"`
	UserID   int       `json:"user_id"`
	Products []Product `json:"products"`
	Total    float64   `json:"total"`
}
