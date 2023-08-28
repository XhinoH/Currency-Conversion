package main

import (
	"gocapri/db"
	"gocapri/handler"
	"gocapri/repository"
	"gocapri/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	db.Init()

	r := gin.Default()
	cr := repository.NewCurrencyRepository(db.DB)
	er := repository.NewExchangeRateRepository(db.DB)
	ch := handler.NewCurrencyHandler(cr, er)
	routes.RegisterCurrencyRoutes(r, ch)
	r.Run("localhost:8080")

}
