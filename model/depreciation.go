package model

type Depreciation struct {
	Id               int         `json:"item_id"`
	Name             string      `json:"name"`
	Price            interface{} `json:"initial_price"`
	DepreciatedValue interface{} `json:"depreciated_value"`
	DepreciationRate interface{} `json:"depreciation_rate"`
}
