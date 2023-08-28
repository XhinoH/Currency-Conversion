package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"gocapri/model"
	"gocapri/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CurrencyHandler struct {
	cr repository.CurrencyRepository
	er repository.ExchangeRateRepository
}

func NewCurrencyHandler(cr repository.CurrencyRepository, er repository.ExchangeRateRepository) *CurrencyHandler {
	return &CurrencyHandler{cr: cr, er: er}
}

func (ch *CurrencyHandler) GetAllCurrencies(c *gin.Context) {

	currencies := ch.cr.GetAllCurrencies()
	c.IndentedJSON(http.StatusOK, currencies)

}

func (ch *CurrencyHandler) GetCurrencyByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.Status(http.StatusBadRequest)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID - must be integer"})
		return
	}

	currency, err := ch.cr.FindCurrencyById(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Currency not found"})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "An error occurred"})
		return
	}

	c.IndentedJSON(http.StatusOK, currency)
}

func (ch *CurrencyHandler) CreateCurrency(c *gin.Context) {

	var currency model.Currency
	err := c.BindJSON(&currency)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Data not valid"})
		return
	}
	err = ch.cr.CreateCurrency(&currency)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Could not create currency"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Currency created successfully"})
}

func (ch *CurrencyHandler) GetCurrencyByIsoCode(c *gin.Context) {
	isoCode := c.Param("isoCode")
	if isoCode == "" {
		c.Status(http.StatusBadRequest)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Empty ISO Code"})
		return
	}

	currency, err := ch.cr.FindCurrencyByIsoCode(isoCode)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Currency not found"})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "An error occurred"})
		return
	}

	c.IndentedJSON(http.StatusOK, currency)
}

func (ch *CurrencyHandler) UpdateCurrency(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.Status(http.StatusBadRequest)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID - must be integer"})
		return
	}

	var updatedCurrency model.Currency
	err = c.BindJSON(&updatedCurrency)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Data not valid"})
		return
	}

	existingCurrency, err := ch.cr.FindCurrencyById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Currency not found"})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "An error occurred"})
		return
	}

	// Update fields if they're provided in the JSON
	if updatedCurrency.AdjustInd != "" {
		existingCurrency.AdjustInd = updatedCurrency.AdjustInd
	}
	if !updatedCurrency.Begin.IsZero() {
		existingCurrency.Begin = updatedCurrency.Begin
	}
	if updatedCurrency.Decimals != 0 {
		existingCurrency.Decimals = updatedCurrency.Decimals
	}
	if !updatedCurrency.End.IsZero() {
		existingCurrency.End = updatedCurrency.End
	}
	if updatedCurrency.EwuFlag != false {
		existingCurrency.EwuFlag = updatedCurrency.EwuFlag
	}
	if updatedCurrency.IsoCode != "" {
		existingCurrency.IsoCode = updatedCurrency.IsoCode
	}
	if updatedCurrency.RoundingUnit != 0 {
		existingCurrency.RoundingUnit = updatedCurrency.RoundingUnit
	}
	// You can update other fields similarly

	err = ch.cr.UpdateCurrency(existingCurrency)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update currency"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Currency updated successfully"})
}

func (ch *CurrencyHandler) DeleteCurrencyById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.Status(http.StatusBadRequest)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID - must be integer"})
		return
	}

	err = ch.cr.DeleteCurrencyById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Currency not found"})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "An error occurred"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Currency deleted successfully"})

}

func (ch *CurrencyHandler) GetConversionRate(c *gin.Context) {
	var request model.CurrencyConversionRequest

	amount := request.Amount
	sourceCurrency := request.SourceCurrency
	targetCurrency := request.TargetCurrency
	conversionDate := request.ConversionDate

	if amount != 0 {
		c.Status(http.StatusBadRequest)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Empty amount"})
		return
	}
	if sourceCurrency == "" {
		c.Status(http.StatusBadRequest)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Empty sourceCurrency"})
		return
	}

	sourceCurrencyId := ch.cr.GetCurrencyIdFromIsoCode(sourceCurrency)
	targetCurrencyId := ch.cr.GetCurrencyIdFromIsoCode(targetCurrency)

	conversionSeq := ch.er.GetConversionSeqFromCurrencyId(int(sourceCurrencyId))

	conversionRate := ch.cr.GetConversionRate(targetCurrencyId, conversionSeq, conversionDate)
	response := model.CurrencyConversionResponse{
		Amount:       amount * conversionRate,
		ExchangeRate: conversionRate,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the JSON response
	fmt.Println(string(jsonResponse))
}
