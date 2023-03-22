package routes

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ikatwal/url-shortener/api/database"
	"github.com/ikatwal/url-shortener/api/helpers"
	"github.com/redis/go-redis/v9"
)

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type response struct {
	URL            string        `json:"url"`
	CustomShort    string        `json:"short"`
	Expiry         time.Duration `json:"expiry"`
	XRateRemaining int           `json:"rate_limit"`
	XRateLimitRest time.Duration `json:"rate_limit_reset"`
}

func ShortenURL(c *gin.Context) {
	body := new(request)
	if err := c.ShouldBindJSON(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// rate limiting
	db := database.NewClient(1)
	defer db.Close()
	val, err := db.Get(database.Context, c.ClientIP()).Result()
	if err == redis.Nil {
		_ = db.Set(database.Context, c.ClientIP(), os.Getenv("API_QUOTA"), 10*time.Minute).Err()
	} else {
		valInt, _ := strconv.Atoi(val)
		if valInt <= 0 {
			limit, _ := db.TTL(database.Context, c.ClientIP()).Result()
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": fmt.Sprintf("rate limit exceeded %v", limit.Minutes())})
			return
		}
	}

	// validate the url and domain
	if !helpers.IsValidURL(body.URL) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad url entered"})
		return
	}

	//enfirce https, SSL
	body.URL = helpers.EnforceHTTP(body.URL)

	db.Decr(database.Context, c.ClientIP())

}
