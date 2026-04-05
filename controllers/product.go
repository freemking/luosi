package controllers

import (
	"fastener-pro/models"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// ProductsHandler 产品页处理器
func ProductsHandler(c *gin.Context) {
	// 获取分类参数 - 优先从URL路径获取，其次从查询参数获取
	category := c.Param("category")
	if category == "" {
		category = c.Query("category")
	}
	// Remove .html extension if present
	if strings.HasSuffix(category, ".html") {
		category = strings.TrimSuffix(category, ".html")
	}

	// 获取分页参数
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "15")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 || pageSize > 50 {
		pageSize = 15
	}

	// 分页查询产品
	products, total, err := models.GetProductsByCategoryWithPagination(category, page, pageSize)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "products.html", gin.H{
			"title":    "Our Products - Industrial Fasteners, Bolts, Nuts, Screws | Yuanmao Fastener",
			"error":    "Failed to load products",
			"products": []models.Product{},
			"category": category,
			"page":     page,
			"pageSize": pageSize,
			"total":    0,
			"pages":    0,
		})
		return
	}

	// 计算总页数
	pages := int(total) / pageSize
	if int(total)%pageSize > 0 {
		pages++
	}

	// 获取page header广告
	pageHeaderAd, _ := models.GetAdsByPositionCode("products")
	var firstAd *models.Ad
	if len(pageHeaderAd) > 0 {
		firstAd = &pageHeaderAd[0]
	}

	c.HTML(200, "products.html", gin.H{
		"title":         "Our Products - Industrial Fasteners, Bolts, Nuts, Screws | Yuanmao Fastener",
		"products":      products,
		"category":      category,
		"page":          page,
		"pageSize":      pageSize,
		"total":         total,
		"pages":         pages,
		"pageHeaderAd":  firstAd,
		"active":        "products",
	})
}

// ProductDetailHandler 产品详情页处理器
func ProductDetailHandler(c *gin.Context) {
	// 获取产品ID
	idStr := c.Param("id")

	// Handle .html extension if present
	if strings.HasSuffix(idStr, ".html") {
		idStr = strings.TrimSuffix(idStr, ".html")
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.HTML(http.StatusBadRequest, "product-detail.html", gin.H{
			"title":  "Product | Yuanmao Fastener",
			"error":  "Invalid product ID",
			"active": "products",
		})
		return
	}

	// 根据ID查询产品
	product, err := models.GetProductByID(uint(id))
	if err != nil {
		c.HTML(http.StatusNotFound, "product-detail.html", gin.H{
			"title":  "Product | Yuanmao Fastener",
			"error":  "Product not found",
			"active": "products",
		})
		return
	}

	// 创建HTML内容字段，避免HTML标签被转义
	htmlDescription := template.HTML(product.Description)

	c.HTML(200, "product-detail.html", gin.H{
		"title":           "Product | Yuanmao Fastener",
		"product":         product,
		"htmlDescription": htmlDescription,
		"active":          "products",
	})
}
