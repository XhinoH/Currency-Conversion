package handler

import (
	"errors"
	"gocapri/model"
	"gocapri/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllCurrencies(c *gin.Context) {
	currencies := repository.Repo().GetAllCurrencies()
	c.IndentedJSON(http.StatusOK, currencies)
}

func GetCurrencyByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.Status(http.StatusBadRequest)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID - must be integer"})
		return
	}

	currency, err := repository.Repo().FindCurrencyById(id)

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

func CreateCurrency(c *gin.Context) {

	var currency model.Currency
	err := c.BindJSON(&currency)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Data not valid"})
		return
	}
	err = repository.Repo().CreateCurrency(&currency)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Could not create currency"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Currency created successfully"})
}

func GetCurrencyByIsoCode(c *gin.Context) {
	isoCode := c.Param("isoCode")
	if isoCode == "" {
		c.Status(http.StatusBadRequest)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Empty ISO Code"})
		return
	}

	currency, err := repository.Repo().FindCurrencyByIsoCode(isoCode)

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

func UpdateCurrency(c *gin.Context) {
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

	existingCurrency, err := repository.Repo().FindCurrencyById(id)
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

	err = repository.Repo().UpdateCurrency(existingCurrency)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update currency"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Currency updated successfully"})
}

func DeleteCurrencyById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.Status(http.StatusBadRequest)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID - must be integer"})
		return
	}

	err = repository.Repo().DeleteCurrencyById(id)
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
