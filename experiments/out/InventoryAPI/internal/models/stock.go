package models

type Stock struct {
	ID          int `json:"id"`
	ProductID   int `json:"product_id"`
	WarehouseID int `json:"warehouse_id"`
	Quantity    int `json:"quantity"`
}
