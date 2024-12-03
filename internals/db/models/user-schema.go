package models

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
