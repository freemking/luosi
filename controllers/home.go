package controllers

import (
	"nexfasten/models"

	"github.com/gin-gonic/gin"
)

// HomeHandler 首页处理器
func HomeHandler(c *gin.Context) {
	// 获取首页轮播广告
	heroAds, _ := models.GetAdsByPositionCode("home")
	// 获取国际销售版块广告
	internationalSalesAds, _ := models.GetAdsByPositionCode("home-international-sales")
	// 获取关于版块广告
	aboutAds, _ := models.GetAdsByPositionCode("home-about")

	c.HTML(200, "index.html", gin.H{
		"title":                "Premium Fasteners & Bolts Supplier - High Quality Industrial Fasteners",
		"heroAds":              heroAds,
		"internationalSalesAd": internationalSalesAds,
		"aboutAd":              aboutAds,
		"categories":           models.GetCategories(),
		"active":               "home",
	})
}
