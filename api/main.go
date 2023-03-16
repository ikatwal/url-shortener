package main

import "github.com/gin-gonic/gin"

func initializeRoutes(router *gin.Engine) {
	router.GET("/:url", routes.ResolveURL)
	router.POST("/api/v1", routes.ShortenURL)
}

func main() {
	err := godotenv.Load()
	router := gin.Default()
	initializeRoutes(router)
}
