package models

import "encoding/json"

// MenuItem represents a menu item entity.
type MenuItem struct {
	ID           int     `json:"id"`
	RestaurantID int     `json:"restaurant_id"`
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
}

// ToJSON converts a MenuItem to JSON.
func (m *MenuItem) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}
