package model

type Item struct {
	Id             int         `json:"id,omitempty"`
	Name           string      `json:"name"`
	CategoryId     int         `json:"-"`
	Category       string      `json:"category"`
	PhotoUrl       string      `json:"photo_url"`
	Price          interface{} `json:"price"`
	PurchaseDate   string      `json:"purchase_date"`
	TotalUsageDays int         `json:"total_usage_days"`
}
