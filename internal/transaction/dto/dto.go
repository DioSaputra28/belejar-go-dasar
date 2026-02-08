package dto

type CheckoutItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type CheckoutRequest struct {
	Items []CheckoutItem `json:"items"`
}

type ReportRequest struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}
