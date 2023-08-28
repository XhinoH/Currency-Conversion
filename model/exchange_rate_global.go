package model

import (
	"time"
)

type ExchangeRateGlobal struct {
	ID                 int      `gorm:"primaryKey"`
	Begin              *time.Time `gorm:"type:date; column:begin"`
	ConversionRate     float64    `gorm:"not null; type:decimal(11,6); column:conversionRate"`
	CurrencyID         int      `gorm:"not null; column:currencyId"`
	DivFlag            BoolBit    `gorm:"not null; column:divFlag"`
	End                *time.Time `gorm:"type:date; column:end"`
	ExchangeRateKindID int      `gorm:"default:NULL; column:exchangeRateKindId"`
	Seq                int      `gorm:"not null; column:seq"`
	UnitsSource        int        `gorm:"not null; column:unitsSource"`
	UnitsTarget        int        `gorm:"not null; column:unitsTarget"`
	LastUpdate         time.Time  `gorm:"default:current_timestamp; column:last_update"`
}

func (ExchangeRateGlobal) TableName() string {
	return "exchange_rate_globals"
}
