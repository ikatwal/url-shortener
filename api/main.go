package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ikatwal/url-shortener/api/routes"
	"github.com/joho/godotenv"
)

func initializeRoutes(router *gin.Engine) {
	router.GET("/:url", routes.ResolveURL)
	router.POST("/api/v1", routes.ShortenURL)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env file: %v", err)
	}
	router := gin.Default()
	initializeRoutes(router)
	log.Fatal(router.Run(os.Getenv("APP_PORT")))
}
