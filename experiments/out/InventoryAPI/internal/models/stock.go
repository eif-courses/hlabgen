package models

type Stock struct {
	ProductID   int `json:"product_id"`
	WarehouseID int `json:"warehouse_id"`
	Quantity    int `json:"quantity"`
}
