package routes

import (
	"fastener-pro/models"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// 首页
	r.GET("/", HomeHandler)

	// 产品页
	r.GET("/products", ProductsHandler)
	
	// 产品详情页
	r.GET("/product/:id", ProductDetailHandler)

	// 新闻页
	r.GET("/news", NewsListHandler)
	
	// 新闻详情页
	r.GET("/news/:id", NewsDetailHandler)

	// 关于我们
	r.GET("/about", AboutHandler)

	// 联系我们
	r.GET("/contact", ContactHandler)
	// 提交联系表单
	r.POST("/contact", ContactSubmitHandler)

	// FAQ页面
	r.GET("/faq", FAQHandler)

}

// 首页处理器
func HomeHandler(c *gin.Context) {
	// 获取产品列表（用于首页展示）
	products, _, err := models.GetProductsWithPagination(1, 9)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "index.html", gin.H{
			"title":   "Premium Fasteners & Bolts Supplier - High Quality Industrial Fasteners",
			"error":   "Failed to load products",
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

// 产品页处理器
func ProductsHandler(c *gin.Context) {
	// 获取分类参数
	category := c.Query("category")

	// 获取分页参数
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "12")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 || pageSize > 50 {
		pageSize = 12
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

	c.HTML(200, "products.html", gin.H{
		"title":    "Our Products - Industrial Fasteners, Bolts, Nuts, Screws | Yuanmao Fastener",
		"products": products,
		"category": category,
		"page":     page,
		"pageSize": pageSize,
		"total":    total,
		"pages":    pages,
		"active":   "products",
	})
}

// 产品详情页处理器
func ProductDetailHandler(c *gin.Context) {
	// 获取产品ID
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.HTML(http.StatusBadRequest, "product-detail.html", gin.H{
			"title": "Product | Yuanmao Fastener",
			"error": "Invalid product ID",
			"active": "products",
		})
		return
	}

	// 根据ID查询产品
	product, err := models.GetProductByID(uint(id))
	if err != nil {
		c.HTML(http.StatusNotFound, "product-detail.html", gin.H{
			"title": "Product | Yuanmao Fastener",
			"error": "Product not found",
			"active": "products",
		})
		return
	}

	// 创建HTML内容字段，避免HTML标签被转义
	htmlDescription := template.HTML(product.Description)

	c.HTML(200, "product-detail.html", gin.H{
		"title":   "Product | Yuanmao Fastener",
		"product": product,
		"htmlDescription": htmlDescription,
		"active":   "products",
	})
}

// 关于我们处理器
func AboutHandler(c *gin.Context) {
	c.HTML(200, "about.html", gin.H{
		"title": "About Us - Professional Fastener Manufacturer | Yuanmao Fastener",
		"active": "about",
	})
}

// 联系我们处理器
func ContactHandler(c *gin.Context) {
	c.HTML(200, "contact.html", gin.H{
		"title": "Contact Us - Yuanmao Fastener | Get Quote for Industrial Fasteners",
		"active": "contact",
	})
}

// 联系表单提交处理器
func ContactSubmitHandler(c *gin.Context) {
	// 解析表单数据
	name := c.PostForm("name")
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	company := c.PostForm("company")
	product := c.PostForm("product")
	message := c.PostForm("message")

	// 验证必填字段
	if name == "" || email == "" || message == "" {
		c.HTML(http.StatusBadRequest, "contact.html", gin.H{
			"title": "Contact Us - Yuanmao Fastener | Get Quote for Industrial Fasteners",
			"error": "Name, email, and message are required",
			"name":    name,
			"email":   email,
			"phone":   phone,
			"company": company,
			"product": product,
			"message": message,
			"active":   "contact",
		})
		return
	}

	// 创建联系表单记录
	contact := models.Contact{
		Name:    name,
		Email:   email,
		Phone:   phone,
		Company: company,
		Product: product,
		Message: message,
	}

	// 保存到数据库
	err := models.CreateContact(contact)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "contact.html", gin.H{
			"title": "Contact Us - Yuanmao Fastener | Get Quote for Industrial Fasteners",
			"error": "Failed to submit form",
			"name":    name,
			"email":   email,
			"phone":   phone,
			"company": company,
			"product": product,
			"message": message,
			"active":   "contact",
		})
		return
	}

	// 提交成功
	c.HTML(200, "contact.html", gin.H{
		"title":    "Contact Us - Yuanmao Fastener | Get Quote for Industrial Fasteners",
		"success":  "Form submitted successfully",
		"active":   "contact",
	})
}

// FAQ页面处理器
func FAQHandler(c *gin.Context) {
	c.HTML(200, "faq.html", gin.H{
		"title": "FAQ - Yuanmao Fastener | Frequently Asked Questions",
		"active": "faq",
	})
}

// NewsListHandler 新闻列表处理器
func NewsListHandler(c *gin.Context) {
	// 获取分页参数
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "9")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 || pageSize > 50 {
		pageSize = 9
	}

	// 分页查询新闻
	newsList, total, err := models.GetNewsList(page, pageSize)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "news.html", gin.H{
			"title":    "News - Industrial Fastener Industry Updates",
			"error":    "Failed to load news",
			"newsList": []models.News{},
			"page":     page,
			"pageSize": pageSize,
			"total":    0,
			"pages":    0,
			"active":   "news",
		})
		return
	}

	// 计算总页数
	pages := int(total) / pageSize
	if int(total)%pageSize > 0 {
		pages++
	}

	c.HTML(200, "news.html", gin.H{
		"title":    "News - Industrial Fastener Industry Updates",
		"newsList": newsList,
		"page":     page,
		"pageSize": pageSize,
		"total":    total,
		"pages":    pages,
		"active":   "news",
	})
}

// NewsDetailHandler 新闻详情处理器
func NewsDetailHandler(c *gin.Context) {
	// 获取新闻ID
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.HTML(http.StatusBadRequest, "news-detail.html", gin.H{
			"title": "News | Yuanmao Fastener",
			"error": "Invalid news ID",
			"active": "news",
		})
		return
	}

	// 根据ID查询新闻
	news, err := models.GetNewsByID(uint(id))
	if err != nil {
		c.HTML(http.StatusNotFound, "news-detail.html", gin.H{
			"title": "News | Yuanmao Fastener",
			"error": "News not found",
			"active": "news",
		})
		return
	}

	c.HTML(200, "news-detail.html", gin.H{
		"title": news.Title + " - Yuanmao Fastener News",
		"news":  news,
		"active": "news",
	})
}

