package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"your-app/models"
	"your-app/repository"
	"your-app/service"
)

func main() {
	dsn := "root:roor@tcp(your-db-host:your-db-port)/your-db-name?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	currencyRepo := repository.NewCurrencyRepository(db)
	currencyService := service.NewCurrencyService(currencyRepo)

	// Create a new currency
	newCurrency := models.Currency{
		AdjustInd:    "A",
		Begin:        time.Now(),
		Decimals:     2,
		End:          time.Now().AddDate(1, 0, 0),
		EwuFlag:      true,
		IsoCode:      "USD",
		RoundingUnit: 0.01,
		CurrencyID:   1,
	}
	err = currencyService.CreateCurrency(&newCurrency)
	if err != nil {
		log.Printf("Error creating currency: %v", err)
	}

	// Get currency by ID
	currencyID := 1
	currency, err := currencyService.GetCurrencyByID(currencyID)
	if err != nil {
		log.Printf("Error getting currency: %v", err)
	} else {
		fmt.Printf("Currency: %+v\n", currency)
	}

	// Update currency
	currency.IsoCode = "EUR"
	err = currencyService.UpdateCurrency(currency)
	if err != nil {
		log.Printf("Error updating currency: %v", err)
	}

	// Delete currency
	err = currencyService.DeleteCurrency(currencyID)
	if err != nil {
		log.Printf("Error deleting currency: %v", err)
	}
}
