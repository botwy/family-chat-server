package models

type Message struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	User    User   `json:"user"`
}
