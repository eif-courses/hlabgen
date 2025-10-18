package models

import "encoding/json"

// Shipment represents a shipment in the warehouse.
type Shipment struct {
	ID          int    `json:"id"`
	ItemID      int    `json:"item_id"`
	Quantity    int    `json:"quantity"`
	Destination string `json:"destination"`
}

// ToJSON converts a Shipment to JSON.
func (s *Shipment) ToJSON() ([]byte, error) {
	return json.Marshal(s)
}
