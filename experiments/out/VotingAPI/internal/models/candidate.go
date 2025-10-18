package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Candidate represents a candidate entity.
type Candidate struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name       string             `bson:"name" json:"name"`
	ElectionID primitive.ObjectID `bson:"election_id" json:"election_id"`
}
