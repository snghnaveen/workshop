package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/snghnaveen/workshop/src"
)

type ReqLongUrl struct {
	Url string `json:"url"`
}

func LongURL(c *gin.Context) {
	var reqUrl ReqLongUrl
	if err := c.BindJSON(&reqUrl); err != nil {
		c.JSON(400, gin.H{"error": "invalid json"})
		return
	}

	if err := src.ValidateURL(reqUrl.Url); err != nil {
		c.JSON(400, gin.H{"error": "invalid url"})
		return
	}

	if shortURL, ok := src.GetFirstURL(reqUrl.Url); ok {
		c.JSON(200, gin.H{"short_code": shortURL})
		return

	}

	shortCode, err := src.CreateAndGetShortURL(reqUrl.Url)

	if err != nil {
		c.JSON(500, gin.H{"error": "unable to genreate short code"})
		return
	}

	c.JSON(200, gin.H{"short_code": shortCode})

}

func ShortURL(c *gin.Context) {
	shortCode := c.Param("shortCode")

	longURL, err := src.ShortUrlToLongUrl(shortCode)
	if err != nil {
		c.JSON(400, gin.H{"error": "short_code not found"})
	}

	c.JSON(http.StatusOK, gin.H{"long_url": longURL})
	// c.Redirect(302, longURL)
}
