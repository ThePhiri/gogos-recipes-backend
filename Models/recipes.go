package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Recipe struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title,omitempty"`
	UserId      string             `json:"userId" bson:"userId,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
	Culture     string             `json:"culture" bson:"culture,omitempty"`
	Country     string             `json:"country" bson:"country,omitempty"`
	Author      string             `json:"author" bson:"author,omitempty"`
	Ingredients []string           `json:"ingredients" bson:"ingredients,omitempty"`
	Steps       []interface{}      `json:"steps" bson:"steps,omitempty"`
	Image       string             `json:"image" bson:"image,omitempty"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
}
