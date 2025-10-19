package models

type Shipment struct {
	ID          int    `json:"id"`
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
}
