package models

type Body struct {
	Message string `json:"message" validate:"required"`
}
