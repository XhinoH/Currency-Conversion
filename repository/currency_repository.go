package repository

import (
	"gocapri/db"
	"gocapri/model"
	"time"

	"gorm.io/gorm"
)

type CurrencyRepository interface {
	CreateCurrency(currency *model.Currency) error
	FindCurrencyById(id int64) (*model.Currency, error)
	UpdateCurrency(currency *model.Currency) error
	DeleteCurrencyById(id int64) error
	GetAllCurrencies() []model.Currency
	FindCurrencyByIsoCode(isoCode string) (*model.Currency, error)
	GetCurrencyIdFromIsoCode(isoCode string) int64
	GetConversionRate(targetCurrencyId int64, exchangeRateKindId int64, end time.Time) float64
}

type CurrencyRepositoryImpl struct {
	db *gorm.DB
}

func NewCurrencyRepository(db *gorm.DB) *CurrencyRepositoryImpl {
	return &CurrencyRepositoryImpl{db: db}
}

func (r *CurrencyRepositoryImpl) CreateCurrency(currency *model.Currency) error {
	return r.db.Create(currency).Error
}

func (r *CurrencyRepositoryImpl) FindCurrencyById(id int64) (*model.Currency, error) {

	var currency model.Currency
	result := db.DB.First(&currency, id)

	if result.Error != nil {

		return nil, result.Error

	}

	return &currency, nil
}

func (repo *CurrencyRepositoryImpl) UpdateCurrency(currency *model.Currency) error {
	result := db.DB.Save(currency)
	if result.Error != nil {
		return result.Error
	}
	return nil // no error occurred
}

func (r *CurrencyRepositoryImpl) DeleteCurrencyById(id int64) error {

	result := db.DB.Delete(&model.Currency{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil // no error occurred
}

func (r *CurrencyRepositoryImpl) GetAllCurrencies() []model.Currency {
	var currencies []model.Currency
	r.db.Find(&currencies)
	return currencies
}

func (r *CurrencyRepositoryImpl) FindCurrencyByIsoCode(isoCode string) (*model.Currency, error) {
	var currency model.Currency
	result := db.DB.Where("isoCode = ?", isoCode).First(&currency)

	if result.Error != nil {

		return nil, result.Error

	}

	return &currency, nil
}

func (r *CurrencyRepositoryImpl) GetCurrencyIdFromIsoCode(isoCode string) int64 {
	currency := model.Currency{}

	r.db.Table("currencys").
		Select("currencyId").
		Where("isoCode = ?", isoCode).
		First(&currency)
	return int64(currency.CurrencyID)
}

func (r *CurrencyRepositoryImpl) GetConversionRate(targetCurrencyId int64, exchangeRateKindId int64, end time.Time) float64 {

	exchangeRateGlobal := model.ExchangeRateGlobal{}
	r.db.Table("exchange_rate_globals").
		Select("conversionRate").
		Where("currencyId = ? AND exchangeRateKindId = ? AND end = ?", targetCurrencyId, exchangeRateKindId, end).
		First(&exchangeRateGlobal)

	return exchangeRateGlobal.ConversionRate
}

// func (r *CurrencyRepositoryImpl) convertSourceCurrencyToEurUsd (sourceCurrencyId int64) int64 {
// }

