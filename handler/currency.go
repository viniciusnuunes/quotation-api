package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CurrencyHandler struct {
}

func (ch *CurrencyHandler) GetAllCurrencies(ctx *gin.Context) {
	fmt.Println("cheguei no Handler")
	ctx.JSON(http.StatusOK, gin.H{"message": "tudo deu certo"})
}
