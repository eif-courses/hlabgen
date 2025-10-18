package models

import "encoding/json"

// Car represents a car in the rental system.
type Car struct {
	ID          int     `json:"id"`
	Make        string  `json:"make"`
	Model       string  `json:"model"`
	Year        int     `json:"year"`
	PricePerDay float64 `json:"price_per_day"`
}

// ToJSON converts a Car to JSON.
func (c *Car) ToJSON() ([]byte, error) {
	return json.Marshal(c)
}

// FromJSON converts JSON to a Car.
func FromJSON(data []byte) (*Car, error) {
	var c Car
	err := json.Unmarshal(data, &c)
	return &c, err
}
