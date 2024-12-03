package models

type UserSchema struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Type     string `json:"type"`
	Age      int    `json:"age"`
	IsActive bool   `json:"is_active"`
}
