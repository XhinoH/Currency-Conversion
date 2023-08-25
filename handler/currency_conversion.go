package handler

// import (
// 	"errors"
// 	"gocapri/repository"
// 	"net/http"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// 	"gorm.io/gorm"
// )

// func ConvertCurrency(c *gin.Context) {
// 	from := c.Param("from")
// 	to := c.Param("to")
// 	amount, err := strconv.ParseFloat(c.Param("amount"), 64)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount"})
// 		return
// 	}

// 	currency, err := repository.Repo().FindByISOCode(from)
// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			c.JSON(http.StatusNotFound, gin.H{"error": "Currency not found"})
// 			return
// 		}
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred"})
// 		return
// 	}

// }
