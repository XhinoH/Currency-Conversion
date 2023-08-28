package model

import (
	"time"
)

type CurrencyReform struct {
	ID             int      `gorm:"primaryKey"`
	Begin          *time.Time `gorm:"type:date; column:begin"`
	ConvRate       float64    `gorm:"not null; type:decimal(11,6);column:convRate"`
	CurrencyFromID int      `gorm:"default:NULL;colum:currencyFromId"`
	CurrencyToID   int      `gorm:"default:NULL;colum:currencyToId"`
	DivFlag        BoolBit    `gorm:"not null;column:divFlag"`
	End            *time.Time `gorm:"type:date;column:end"`
	Seq            int      `gorm:"not null;column:seq"`
	UnitsSource    int        `gorm:"not null;column:unitsSource"`
	UnitsTarget    int        `gorm:"not null;column:unitsTarget"`
	LastUpdate     time.Time  `gorm:"default:current_timestamp;colum:last_update"`
}

func (CurrencyReform) TableName() string {
	return "currency_reforms"
}
