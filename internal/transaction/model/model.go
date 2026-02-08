package model

import "time"

type Transaction struct {
	ID          int       `json:"id" example:"1"`
	TotalAmount int       `json:"total_amount" example:"15000"`
	CreatedAt   time.Time `json:"created_at"`
}

type TransactionDetail struct {
	ID            int `json:"id" example:"1"`
	TransactionID int `json:"transaction_id" example:"1"`
	ProductID     int `json:"product_id" example:"1"`
	Quantity      int `json:"quantity" example:"2"`
	Subtotal      int `json:"subtotal" example:"7000"`
}

type ProdukTerlaris struct {
	Nama       string `json:"nama"`
	QtyTerjual int    `json:"qty_terjual"`
}

type ReportResponse struct {
	TotalRevenue   int            `json:"total_revenue"`
	TotalTransaksi int            `json:"total_transaksi"`
	ProdukTerlaris ProdukTerlaris `json:"produk_terlaris"`
}
