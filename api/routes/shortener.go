package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ikatwal/url-shortener/api/helpers"
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
	// validate the url and domain
	if !helpers.IsValidURL(body.URL) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad url entered"})
		return
	}

}
