package main

import (
	"GO_APP/pkg/api"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Configuración de CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:8080"
		},
		MaxAge: 12 * time.Hour,
	}))
	// Configura CORS si es necesario (ver código CORS anterior)
	router.GET("/stocks", api.FetchStocksHandler)

	if err := router.Run(":8081"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
