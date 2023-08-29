package repository

import (
	"gocapri/model"

	"gorm.io/gorm"
)

type ExchangeRateRepository interface {
	GetConversionSeqFromCurrencyId(currencyId int) int64
}

type ExchangeRateRepositoryImpl struct {
	db *gorm.DB
}

func NewExchangeRateRepository(db *gorm.DB) *ExchangeRateRepositoryImpl {
	return &ExchangeRateRepositoryImpl{db: db}
}

func (r *ExchangeRateRepositoryImpl) GetConversionSeqFromCurrencyId(currencyId int) int64 {
	exchangeRate := model.ExchangeRateKind{}

	r.db.Table("exchange_rate_kinds").
		Select("seq").
		Where("currencyTargetId = ?", currencyId).
		Limit(1).
		First(&exchangeRate)
	return int64(exchangeRate.Seq)
}
