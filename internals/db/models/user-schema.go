package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserSchema struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Type     string `json:"type"`
	Age      int    `json:"age"`
	IsActive bool   `json:"is_active"`
}

// TopicSchema represents the structure of a topic in the database.
type TopicSchema struct {
	ID          string `json:"id" bson:"_id,omitempty"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Position    int    `json:"position" bson:"position"`
	Hidden      bool   `json:"isHidden" bson:"isHidden"`
	CreatedAt   string `json:"createdAt" bson:"createdAt"`
	UpdatedAt   string `json:"updatedAt" bson:"updatedAt"`
}

type PageSchema struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	TopicID   primitive.ObjectID `bson:"topic_id"`
	Title     string             `bson:"title"`
	Content   string             `bson:"content"`
	CreatedAt string             `bson:"created_at"`
	UpdatedAt string             `bson:"updated_at"`
	DeletedAt string             `bson:"deleted_at,omitempty"`
	Hidden    bool               `bson:"hidden,omitempty"`
}
