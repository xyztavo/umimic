package models

type Message struct {
	Role    string `json:"role" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type Body struct {
	Message string    `json:"message" validate:"required"`
	History []Message `json:"history"`
}
