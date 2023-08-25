package handler

import (
	"errors"
	"gocapri/db"
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

	currency, err := repository.Repo().FindById(id)

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
	err = repository.Repo().Create(&currency)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Could not create currency"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Currency created successfully"})
}

func UpdateCurrency(c *gin.Context) {
	var currency model.Currency
	id := c.Param("id")
	db.DB.First(&currency, id)
	c.BindJSON(&currency)
	db.DB.Save(&currency)
	c.IndentedJSON(http.StatusOK, currency)
}

func DeleteCurrency(c *gin.Context) {
	var currency model.Currency
	id := c.Param("id")
	db.DB.Delete(&currency, id)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Currency deleted successfully"})
}

func GetCurrencyByIsoCode(c *gin.Context) {
	isoCode := c.Param("isoCode")
	if isoCode == "" {
		c.Status(http.StatusBadRequest)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Empty ISO Code"})
		return
	}

	currency, err := repository.Repo().FindByISOCode(isoCode)

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
