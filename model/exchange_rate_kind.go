package model

import (
	"time"
)

type ExchangeRateKind struct {
	ID                           int      `gorm:"primaryKey"`
	Begin                        *time.Time `gorm:"type:date;colum:begin"`
	CurrencyTargetID             int      `gorm:"default:NULL;colum:currencyTargetId"`
	End                          *time.Time `gorm:"type:date;colum:end"`
	EwuOverruledFlag             BoolBit    `gorm:"not null;colum:ewuOverruledFlag"`
	ExchangeRateKindAlternativID int      `gorm:"default:NULL;colum:exchangeRateKindAlternativId"`
	GlobalFlag                   BoolBit    `gorm:"not null;colum:globalFlag"`
	Seq                          int      `gorm:"not null;colum:seq"`
	LastUpdate                   time.Time  `gorm:"default:current_timestamp;colum:last_update"`
}

func (ExchangeRateKind) TableName() string {
	return "exchange_rate_kinds"
}
