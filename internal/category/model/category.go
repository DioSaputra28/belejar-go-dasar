package model

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
}

var Categories = []Category{
	{ID: 1, Name: "Makanan", Description: "Makanan sehari-hari"},
	{ID: 2, Name: "Minuman", Description: "Minuman sehari-hari"},
	{ID: 3, Name: "Alat Tulis", Description: "Alat tulis sehari-hari"},
}