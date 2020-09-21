package main

import (
	"log"
	"os"

	"github.com/danielthank/exchat-server/handler"
	"github.com/danielthank/exchat-server/injector"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}
	gin.SetMode(os.Getenv("GIN_MODE"))
}

func main() {
	wsHandler := injector.InjectWSHandler()
	authHandler := injector.InjectAuthHandler()

	router := gin.Default()
	handler.AddWSRoutes(router, wsHandler)
	handler.AddAuthRoutes(router, authHandler)

	router.Run(":8080")
}
