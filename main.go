package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	database "github.com/viniciusnuunes/quotation-api/infrastructure/db"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}
	defer db.Close()

	r := gin.Default()

	r.GET("/hello", healthCheck)

	r.Run()
}

func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}
