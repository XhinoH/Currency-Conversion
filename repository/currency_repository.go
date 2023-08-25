package repository

import (
	"gocapri/db"
	"gocapri/model"

	"gorm.io/gorm"
)

type CurrencyRepository interface {
	Create(currency *model.Currency) error
	FindById(id int64) (*model.Currency, error)
	Update(currency *model.Currency) error
	Delete(id int64) error
	GetAllCurrencies() []model.Currency
	FindByISOCode(isoCode string) (*model.Currency, error)
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

func (r *currencyRepository) Create(currency *model.Currency) error {
	return r.db.Create(currency).Error
}

func (r *currencyRepository) FindById(id int64) (*model.Currency, error) {

	var currency model.Currency
	result := db.DB.First(&currency, id)

	if result.Error != nil {

		return nil, result.Error

	}

	return &currency, nil
}

func (r *currencyRepository) Update(currency *model.Currency) error {
	return r.db.Save(currency).Error
}

func (r *currencyRepository) Delete(id int64) error {
	return r.db.Delete(&model.Currency{}, id).Error
}

func (r *currencyRepository) GetAllCurrencies() []model.Currency {
	var currencies []model.Currency
	r.db.Find(&currencies)
	return currencies
}

func (r *currencyRepository) FindByISOCode(isoCode string) (*model.Currency, error) {
	var currency model.Currency
	result := db.DB.Where("isoCode = ?", isoCode).First(&currency)

	if result.Error != nil {

		return nil, result.Error

	}

	return &currency, nil
}
