package controllers

import (
	"fastener-pro/models"

	"github.com/gin-gonic/gin"
)

// HomeHandler 首页处理器
func HomeHandler(c *gin.Context) {
	// 获取首页轮播广告
	heroAds, _ := models.GetAdsByPositionCode("home")

	c.HTML(200, "index.html", gin.H{
		"title":   "Premium Fasteners & Bolts Supplier - High Quality Industrial Fasteners",
		"heroAds": heroAds,
		"active":  "home",
	})
}
