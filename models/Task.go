package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Message     string             `json:"message"`
	Description string             `json:"description"`
	User        User               `json:"user"`
}
