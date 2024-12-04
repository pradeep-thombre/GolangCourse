package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserSchema struct {
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Type     string `json:"type" bson:"type"`
	Age      int    `json:"age" bson:"age"`
	IsActive bool   `json:"is_active" bson:"is_active"`
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
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	TopicID   primitive.ObjectID `json:"topic_id" bson:"topic_id"`
	Title     string             `json:"title" bson:"title"`
	Type      string             `json:"type" bson:"type"`
	Content   interface{}        `json:"content" bson:"content"`
	CreatedAt string             `json:"created_at" bson:"created_at"`
	UpdatedAt string             `json:"updated_at" bson:"updated_at"`
	DeletedAt string             `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
	Hidden    bool               `json:"is_hidden,omitempty" bson:"isHidden,omitempty"`
}
