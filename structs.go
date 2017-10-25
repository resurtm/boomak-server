package main

type AuthEntry struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Id       int    `json:"id" bson:"_id,omitempty"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
