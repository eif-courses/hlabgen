package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Order represents an order in the inventory.
type Order struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ProductID primitive.ObjectID `bson:"product_id" json:"product_id"`
	Quantity  int                `bson:"quantity" json:"quantity"`
	Total     float64            `bson:"total" json:"total"`
}
