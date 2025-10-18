package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Supplier represents a supplier in the inventory.
type Supplier struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name    string             `bson:"name" json:"name"`
	Contact string             `bson:"contact" json:"contact"`
}
