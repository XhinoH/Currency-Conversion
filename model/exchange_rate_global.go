package model

import (
	"time"
)

type ExchangeRateGlobal struct {
	ID                 int64
	Begin              time.Time
	ConversionRate     float64
	CurrencyID         int64
	DivFlag            BoolBit
	End                time.Time
	ExchangeRateKindID int64
	Seq                int64
	UnitsSource        int
	UnitsTarget        int
	LastUpdate         time.Time
}

func (c *ExchangeRateGlobal) TableName() string {
	return "exchange_rate_globals"
}
