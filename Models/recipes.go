package models

import "time"

type Recipe struct {
	ID          string    `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string    `json:"name" bson:"name,omitempty"`
	Description string    `json:"description" bson:"description,omitempty"`
	Ingredients []string  `json:"ingredients" bson:"ingredients,omitempty"`
	Steps       []string  `json:"steps" bson:"steps,omitempty"`
	Image       string    `json:"image" bson:"image,omitempty"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at,omitempty"`
}
