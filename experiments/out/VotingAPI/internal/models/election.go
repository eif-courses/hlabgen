package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Election represents an election entity.
type Election struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Date        string             `bson:"date" json:"date"`
}
