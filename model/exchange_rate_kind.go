package model

import (
	"time"
)

type ExchangeRateKind struct {
	ID                           int64
	Begin                        time.Time
	CurrencyTargetID             int64
	End                          time.Time
	EwuOverruledFlag             BoolBit
	ExchangeRateKindAlternativID int64
	GlobalFlag                   BoolBit
	Seq                          int64
	LastUpdate                   time.Time
}

func (c *ExchangeRateKind) TableName() string {
	return "exchange_rate_kinds"
}
