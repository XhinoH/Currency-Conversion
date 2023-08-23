package repository

import (
	"gocapri/model"

	"gorm.io/gorm"
)

type CurrencyRepository interface {
	Create(currency *model.Currency) error
	FindById(id int64) (*model.Currency, error)
	Update(currency *model.Currency) error
	Delete(id int64) error
}

type currencyRepository struct {
	db *gorm.DB
}

func NewCurrencyRepository(db *gorm.DB) CurrencyRepository {
	return &currencyRepository{db: db}
}

func (r *currencyRepository) Create(currency *model.Currency) error {
	return r.db.Create(currency).Error
}

func (r *currencyRepository) FindById(id int64) (*model.Currency, error) {
	var currency model.Currency
	err := r.db.First(&currency, id).Error
	if err != nil {
		return nil, err
	}
	return &currency, nil
}

func (r *currencyRepository) Update(currency *model.Currency) error {
	return r.db.Save(currency).Error
}

func (r *currencyRepository) Delete(id int64) error {
	return r.db.Delete(&model.Currency{}, id).Error
}
