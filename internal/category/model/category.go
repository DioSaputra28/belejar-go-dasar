package model

import "time"

type Category struct {
	ID          int       `json:"id" example:"1"`
	Name        string    `json:"name" example:"Makanan"`
	Description string    `json:"description" example:"Makanan sehari-hari"`
	CreatedAt   time.Time `json:"created_at"`
}
