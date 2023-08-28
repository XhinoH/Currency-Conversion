package model

import "time"

type CurrencyConversionRequest struct {
	Amount         float64   `json:"amount"`
	ConversionDate time.Time `json:"conversionDate"`
	SourceCurrency string    `json:"sourceCurrency"`
	TargetCurrency string    `json:"targetCurrency"`
}
