package controllers

import (
	"html/template"
	"net/http"
	"nexfasten/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ProductsHandler 产品页处理器
func ProductsHandler(c *gin.Context) {
	// 获取分类 slug 参数 - 优先从URL路径获取，其次从查询参数获取
	slug := c.Param("category_slug")

	// 通过 slug 获取分类信息
	category := ""
	if slug != "" {
		if cat := models.GetCategoryBySlug(slug); cat != nil {
			category = cat.Name
		}
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
			"title":      "Our Products - Industrial Fasteners, Bolts, Nuts, Screws | Yuanmao Fastener",
			"error":      "Failed to load products",
			"products":   []models.Product{},
			"category":   category,
			"page":       page,
			"pageSize":   pageSize,
			"total":      0,
			"pages":      0,
			"categories": models.GetCategories(),
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
		"slug":          slug,
		"page":          page,
		"pageSize":      pageSize,
		"total":         total,
		"pages":         pages,
		"pageHeaderAd":  firstAd,
		"categories":    models.GetCategories(),
		"active":        "products",
	})
}

// ProductDetailHandler 产品详情页处理器
func ProductDetailHandler(c *gin.Context) {
	categorySlug := c.Param("category_slug")
	slug := c.Param("slug")

	product, err := models.GetProductBySlug(slug)
	if err != nil {
		c.HTML(http.StatusNotFound, "product-detail.html", gin.H{
			"title":  "Product | Yuanmao Fastener",
			"error":  "Product not found",
			"active": "products",
		})
		return
	}

	htmlDescription := template.HTML(product.Description)

	relatedProducts, _ := models.GetRelatedProducts(categorySlug, product.ID, 4)

	c.HTML(200, "product-detail.html", gin.H{
		"title":            "Product | Yuanmao Fastener",
		"product":          product,
		"categorySlug":     categorySlug,
		"htmlDescription":  htmlDescription,
		"relatedProducts":  relatedProducts,
		"categories":       models.GetCategories(),
		"active":           "products",
	})
}
