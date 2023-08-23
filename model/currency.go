package model

import (
	"time"
)

type Currency struct {
	ID           int64     `gorm:"primaryKey"`
	AdjustInd    string    `gorm:"size:1"`
	Begin        time.Time `gorm:"type:date"`
	Decimals     int       `gorm:"not null"`
	End          time.Time `gorm:"type:date"`
	EwuFlag      BoolBit   `gorm:"not null"`
	IsoCode      string    `gorm:"size:3"`
	RoundingUnit float64   `gorm:"not null"`
	LastUpdate   time.Time `gorm:"default:current_timestamp"`
	CurrencyID   int64     `gorm:"not null"`
}

func (Currency) TableName() string {
	return "currencys"
}
