package model

type Item struct {
	Id             int    `json:"id,omitempty"`
	Name           string `json:"name"`
	Category       string `json:"category"`
	PhotoUrl       string `json:"photo_url"`
	Price          int    `json:"price"`
	PurchaseDate   string `json:"purchase_date"`
	TotalUsageDays int    `json:"total_usage_days"`
}
