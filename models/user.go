package models

type User struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	// IsCurrentUser bool   `json:"isCurrentUser"`
}