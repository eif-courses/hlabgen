package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Author represents a news author.
type Author struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name  string             `bson:"name" json:"name"`
	Email string             `bson:"email" json:"email"`
}
