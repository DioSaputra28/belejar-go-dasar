package model

// Category represents a product category in the system
type Category struct {
	ID          int    `json:"id" example:"1"`                            // Category ID
	Name        string `json:"name" example:"Makanan"`                    // Category name
	Description string `json:"description" example:"Makanan sehari-hari"` // Category description
}

var Categories = []Category{
	{ID: 1, Name: "Makanan", Description: "Makanan sehari-hari"},
	{ID: 2, Name: "Minuman", Description: "Minuman sehari-hari"},
	{ID: 3, Name: "Alat Tulis", Description: "Alat tulis sehari-hari"},
}
