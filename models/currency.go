package models

type Currency struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	Symbol  string `json:"symbol"`
	Country string `json:"country"`
	Decimal int    `json:"decimal"`
}
