package routes

import (
	"gocapri/handler"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()
	r.GET("/currencies", handler.GetAllCurrencies)
	r.GET("/currencies/:id", handler.GetCurrencyByID)
	r.POST("/currencies", handler.CreateCurrency)
	r.GET("/currencies/iso/:isoCode", handler.GetCurrencyByIsoCode)
	return r
}
