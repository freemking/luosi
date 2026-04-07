package controllers

import (
	"net/http"
	"nexfasten/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// SearchHandler handles search requests
func SearchHandler(c *gin.Context) {
	keyword := c.Query("q")
	limitStr := c.DefaultQuery("limit", "10")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}
	if limit > 50 {
		limit = 50
	}

	products, err := models.SearchProducts(keyword, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Search failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"products": products,
		"count":    len(products),
		"keyword":  keyword,
	})
}
