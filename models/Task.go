package models

type Task struct {
	ID      int    `json:"id" bson:"id,omitempty"`
	Message string `json:"message"`
	User    User   `json:"user"`
}
