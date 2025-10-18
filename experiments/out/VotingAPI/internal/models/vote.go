package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Vote represents a vote entity.
type Vote struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ElectionID  primitive.ObjectID `bson:"election_id" json:"election_id"`
	CandidateID primitive.ObjectID `bson:"candidate_id" json:"candidate_id"`
	VoterID     primitive.ObjectID `bson:"voter_id" json:"voter_id"`
}
