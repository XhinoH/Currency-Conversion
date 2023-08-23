package service

import (
	"gocapri/model"
	"gocapri/repository"
)

type currencyService struct {
	repo repository.CurrencyRepository
}

func NewCurrencyService(repo repository.CurrencyRepository) CurrencyService {
	return &currencyService{repo: repo}
}

func (s *currencyService) CreateCurrency(currency *model.Currency) error {
	return s.repo.Create(currency)
}

func (s *currencyService) GetCurrencyByID(id int64) (*model.Currency, error) {
	return s.repo.FindById(id)
}

func (s *currencyService) UpdateCurrency(currency *model.Currency) error {
	return s.repo.Update(currency)
}

func (s *currencyService) DeleteCurrency(id int64) error {
	return s.repo.Delete(id)
}
