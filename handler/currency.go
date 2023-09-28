package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type CurrencyHandler struct {
}

func (ch *CurrencyHandler) GetAllCurrencies(ctx *gin.Context) {
	fmt.Println("cheguei no Handler")
}
