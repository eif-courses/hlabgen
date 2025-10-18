package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Category represents a news category.
type Category struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name string             `bson:"name" json:"name"`
}
