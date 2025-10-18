package models

// Restaurant represents a restaurant entity.
type Restaurant struct {
	ID       int        `json:"id"`
	Name     string     `json:"name"`
	Location string     `json:"location"`
	Menu     []MenuItem `json:"menu"`
}

// MenuItem represents a menu item entity.
type MenuItem struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// Order represents an order entity.
type Order struct {
	ID           int        `json:"id"`
	RestaurantID int        `json:"restaurant_id"`
	MenuItems    []MenuItem `json:"menu_items"`
	TotalPrice   float64    `json:"total_price"`
	Status       string     `json:"status"`
}
