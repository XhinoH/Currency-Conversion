package model

import (
	"time"
)

type CurrencyReform struct {
	ID             int64
	Begin          time.Time
	ConvRate       float64
	CurrencyFromID int64
	CurrencyToID   int64
	DivFlag        BoolBit
	End            time.Time
	Seq            int64
	UnitsSource    int
	UnitsTarget    int
	LastUpdate     time.Time
}

func (c *CurrencyReform) TableName() string {
	return "currency_reforms"
}
