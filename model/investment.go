package model

type Investment struct {
	TotalInvestment  interface{} `json:"total_investment"`
	DepreciatedValue interface{} `json:"depreciated_value"`
}
