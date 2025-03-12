package main

import (
	"GO_APP/pkg/api"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Configura CORS si es necesario (ver código CORS anterior)
	router.GET("/stocks", api.FetchStocksHandler)

	if err := router.Run(":8081"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
