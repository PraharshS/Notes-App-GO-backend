package models

type User struct {
	ID       int    `bson:"_id,omitempty"`
	Username string `json:"username"`
	Password string `json:"password"`
}
