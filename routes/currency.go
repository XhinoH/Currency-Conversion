package routes

import (
	"gocapri/handler"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()
	return r
}

func RegisterCurrencyRoutes(r *gin.Engine, ch *handler.CurrencyHandler) {

	r.GET("/currencies", ch.GetAllCurrencies)

	r.GET("/currencies/:id", ch.GetCurrencyByID)

	r.POST("/currencies", ch.CreateCurrency)

	r.GET("/currencies/iso/:isoCode", ch.GetCurrencyByIsoCode)

	r.PUT("/currencies/:id", ch.UpdateCurrency)

	r.DELETE("/currencies/:id", ch.DeleteCurrencyById)

	r.GET("/currencies/convert", ch.GetConversionRate)
}
