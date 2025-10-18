package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Article represents a news article.
type Article struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title      string             `bson:"title" json:"title"`
	Content    string             `bson:"content" json:"content"`
	AuthorID   primitive.ObjectID `bson:"author_id" json:"author_id"`
	CategoryID primitive.ObjectID `bson:"category_id" json:"category_id"`
}
