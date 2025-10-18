package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Election represents an election entity.
type Election struct {
	ID         primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Name       string               `bson:"name" json:"name"`
	Date       string               `bson:"date" json:"date"`
	Candidates []primitive.ObjectID `bson:"candidates" json:"candidates"`
}
