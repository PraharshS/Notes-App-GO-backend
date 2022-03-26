package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Message       string             `json:"message"`
	Description   string             `json:"description"`
	IsDone        bool               `json:"is_done" bson:"is_done"`
	Targeted_Date string             `json:"targeted_date" bson:"targeted_date"`
	User          User               `json:"user"`
}
