package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Product represents a product in the inventory.
type Product struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name  string             `bson:"name" json:"name"`
	Price float64            `bson:"price" json:"price"`
	Stock int                `bson:"stock" json:"stock"`
}
