package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Order represents an order in the inventory system.
type Order struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	ProductID primitive.ObjectID `json:"product_id" bson:"product_id"`
	Quantity  int                `json:"quantity" bson:"quantity"`
	Status    string             `json:"status" bson:"status"`
}
