package model

type CurrencyConversionResponse struct {
	Amount       float64 `json:"amount"`
	ExchangeRate float64 `json:"exchangeRate"`
}
