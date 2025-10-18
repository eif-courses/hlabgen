package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Vote represents a vote entity.
type Vote struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CandidateID primitive.ObjectID `bson:"candidate_id" json:"candidate_id"`
	UserID      primitive.ObjectID `bson:"user_id" json:"user_id"`
}
