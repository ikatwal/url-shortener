package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ikatwal/url-shortener/api/database"
	"github.com/redis/go-redis/v9"
)

func ResolveURL(c *gin.Context) {
	url := c.Param("url")
	client := database.NewClient(0)
	defer client.Close()
	val, err := client.Get(database.Context, url).Result()
	if err == redis.Nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot connect to database"})
	}
	dbIncr := database.NewClient(1)
	defer dbIncr.Close()

	_ = dbIncr.Incr(database.Context, "counter")

	c.Redirect(301, val)
}
