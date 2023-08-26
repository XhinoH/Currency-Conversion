package repository

import (
	"gocapri/db"
	"gocapri/model"

	"gorm.io/gorm"
)

type CurrencyRepository interface {
	CreateCurrency(currency *model.Currency) error
	FindCurrencyById(id int64) (*model.Currency, error)
	UpdateCurrency(currency *model.Currency) error
	DeleteCurrencyById(id int64) error
	GetAllCurrencies() []model.Currency
	FindCurrencyByIsoCode(isoCode string) (*model.Currency, error)
}

type currencyRepository struct {
	db *gorm.DB
}

func NewCurrencyRepository(db *gorm.DB) *currencyRepository {
	return &currencyRepository{db: db}
}

var repo *currencyRepository

func Repo() *currencyRepository {
	if repo == nil {
		repo = NewCurrencyRepository(db.DB)

	}
	return repo
}

func (r *currencyRepository) CreateCurrency(currency *model.Currency) error {
	return r.db.Create(currency).Error
}

func (r *currencyRepository) FindCurrencyById(id int64) (*model.Currency, error) {

	var currency model.Currency
	result := db.DB.First(&currency, id)

	if result.Error != nil {

		return nil, result.Error

	}

	return &currency, nil
}

func (repo *currencyRepository) UpdateCurrency(currency *model.Currency) error {
	result := db.DB.Save(currency)
	if result.Error != nil {
		return result.Error
	}
	return nil // no error occurred
}

func (r *currencyRepository) DeleteCurrencyById(id int64) error {

	result := db.DB.Delete(&model.Currency{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil // no error occurred
}

func (r *currencyRepository) GetAllCurrencies() []model.Currency {
	var currencies []model.Currency
	r.db.Find(&currencies)
	return currencies
}

func (r *currencyRepository) FindCurrencyByIsoCode(isoCode string) (*model.Currency, error) {
	var currency model.Currency
	result := db.DB.Where("isoCode = ?", isoCode).First(&currency)

	if result.Error != nil {

		return nil, result.Error

	}

	return &currency, nil
}
