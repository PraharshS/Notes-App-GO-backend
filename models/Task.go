package models

type Task struct {
	ID          int    `json:"id" bson:"id,omitempty"`
	Message     string `json:"message"`
	Description string `json:"description"`
	User        User   `json:"user"`
}
