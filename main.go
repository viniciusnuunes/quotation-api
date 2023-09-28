package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciusnuunes/quotation-api/handler"
	database "github.com/viniciusnuunes/quotation-api/infrastructure/db"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}
	defer db.Close()

	h := handler.MustInit()

	r := gin.Default()

	r.GET("/hello", healthCheck)
	r.GET("/currency", h.Currency.GetAllCurrencies)

	r.Run()
}

func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}
