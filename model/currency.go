package model

import (
	"time"
)

type Currency struct {
	ID           int       `gorm:"primaryKey"`
	AdjustInd    string    `gorm:"size:1 ; column:adjustInd"`
	Begin        time.Time `gorm:"type:date; column:begin"`
	Decimals     int       `gorm:"not null; column:decimals"`
	End          time.Time `gorm:"type:date; column:end"`
	EwuFlag      BoolBit   `gorm:"not null; column:ewuFlag"`
	IsoCode      string    `gorm:"size:3; column:isoCode"`
	RoundingUnit float64   `gorm:"not null; column:roundingUnit"`
	LastUpdate   time.Time `gorm:"default:current_timestamp; column:last_update"`
	CurrencyID   int       `gorm:"not null; column:currencyId"`
}

func (Currency) TableName() string {
	return "currencys"
}
