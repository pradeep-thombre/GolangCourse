package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
	Name     string             `json:"name"`
	Email    string             `json:"email"`
	Type     string             `json:"type"`
	Age      int                `json:"age"`
	IsActive bool               `json:"is_active"`
}

// Topic represents the structure of a topic in the application.
type Topic struct {
	ID          string `json:"id" bson:"_id,omitempty"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Position    int    `json:"position" bson:"position"`
	Hidden      bool   `json:"hidden" bson:"hidden"`
	CreatedAt   string `json:"createdAt" bson:"createdAt"`
	UpdatedAt   string `json:"updatedAt" bson:"updatedAt"`
}
