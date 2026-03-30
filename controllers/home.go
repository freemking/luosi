package controllers

import (
	"fastener-pro/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HomeHandler 首页处理器
func HomeHandler(c *gin.Context) {
	// 获取产品列表（用于首页展示）
	products, _, err := models.GetProductsWithPagination(1, 9)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "index.html", gin.H{
			"title":    "Premium Fasteners & Bolts Supplier - High Quality Industrial Fasteners",
			"error":    "Failed to load products",
			"products": []models.Product{},
			"active":   "home",
		})
		return
	}

	c.HTML(200, "index.html", gin.H{
		"title":    "Premium Fasteners & Bolts Supplier - High Quality Industrial Fasteners",
		"products": products,
		"active":   "home",
	})
}
