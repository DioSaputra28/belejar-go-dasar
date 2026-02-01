package model

import "time"

type Product struct {
	ID        int       `json:"id" example:"1"`
	Name      string    `json:"name" example:"Indomie Godog"`
	Price     int       `json:"price" example:"3500"`
	Stock     int       `json:"stock" example:"10"`
	CreatedAt time.Time `json:"created_at"`
}
