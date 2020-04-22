package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Product defines the properties of a product
type Product struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	Price     float32            `json:"price" bson:"price"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"-" bson:"updated_at"`
}
