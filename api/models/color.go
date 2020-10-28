package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Create color struct
type Color struct {
	ID   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name,omitempty" bson:"name,omitempty"`
}
