package service

import "gocapri/model"

type CurrencyService interface {
	CreateCurrency(currency *model.Currency) error
	GetCurrencyByID(id int64) (*model.Currency, error)
	UpdateCurrency(currency *model.Currency) error
	DeleteCurrency(id int64) error
}
