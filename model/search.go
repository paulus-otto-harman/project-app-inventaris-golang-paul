package model

type Search struct {
	Name  string `json:"name,omitempty"`
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
}
