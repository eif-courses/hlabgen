package models

import "encoding/json"

// Shipment represents a shipment in the warehouse.
type Shipment struct {
	ID       int    `json:"id"`
	ItemID   int    `json:"item_id"`
	Quantity int    `json:"quantity"`
	Status   string `json:"status"`
}

// ToJSON converts the Shipment to JSON format.
func (s *Shipment) ToJSON() ([]byte, error) {
	return json.Marshal(s)
}
