package handler

// import (
// 	"gocapri/model"
// 	"gocapri/repository"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func GetConversionRate(c *gin.Context) {
// 	var request model.CurrencyConversionRequest
// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
// 		return
// 	}

// 	// TODO : Add error handling

// 	convertedAmount := request.Amount * conversionRate

// 	response := model.CurrencyConversionResponse{
// 		Amount:       convertedAmount,
// 		ExchangeRate: conversionRate,
// 	}

// 	c.JSON(http.StatusOK, response)
// }
