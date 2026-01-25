package model

// Produk represents a product in the system
type Produk struct {
	ID    int    `json:"id" example:"1"`               // Product ID
	Nama  string `json:"nama" example:"Indomie Godog"` // Product name
	Harga int    `json:"harga" example:"3500"`         // Product price
	Stok  int    `json:"stok" example:"10"`            // Product stock
}

var Produksi = []Produk{
	{ID: 1, Nama: "Indomie Godog", Harga: 3500, Stok: 10},
	{ID: 2, Nama: "Vit 1000ml", Harga: 3000, Stok: 40},
	{ID: 3, Nama: "kecap", Harga: 12000, Stok: 20},
}
