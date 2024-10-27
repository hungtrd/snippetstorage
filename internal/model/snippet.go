package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const SnippetCollection = "snippets"

type Snippet struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Content   string             `bson:"content,omitempty" json:"content"`
	Title     string             `bson:"title,omitempty" json:"title"`
	Author    string             `bson:"author,omitempty" json:"author"`
	SharedTo  []string           `bson:"sharedTo,omitempty" json:"shared_to"`
	IsPublic  bool               `bson:"isPublic,omitempty" json:"is_public"`
	CreatedAt time.Time          `bson:"createdAt,omitempty" json:"created_at"`
	UpdatedAt time.Time          `bson:"updatedAt,omitempty" json:"updated_at"`
	DeletedAt *time.Time         `bson:"deletedAt,omitempty" json:"deleted_at"`
}
