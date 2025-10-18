package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Supplier represents a supplier in the inventory system.
type Supplier struct {
	ID      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"name" bson:"name"`
	Contact string             `json:"contact" bson:"contact"`
}
