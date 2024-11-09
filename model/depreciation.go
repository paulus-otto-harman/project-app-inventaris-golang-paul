package model

type Depreciation struct {
	Id    int    `json:"item_id"`
	Name  string `json:"name"`
	Price string `json:"initial_price"`
	Value string `json:"depreciated_value"`
	Rate  string `json:"depreciation_rate"`
}
