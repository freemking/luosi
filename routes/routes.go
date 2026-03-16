package routes

import (
	"fastener-pro/models"
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

	// 关于我们
	r.GET("/about", AboutHandler)

	// 联系我们
	r.GET("/contact", ContactHandler)
	// 提交联系表单
	r.POST("/contact", ContactSubmitHandler)

	// FAQ页面
	r.GET("/faq", FAQHandler)

	// Feedback页面
	r.GET("/feedback", FeedbackHandler)
	// 提交Feedback表单
	r.POST("/feedback", FeedbackSubmitHandler)
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
			"title":    "Our Products - Industrial Fasteners, Bolts, Nuts, Screws | Fastener Pro",
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
		"title":    "Our Products - Industrial Fasteners, Bolts, Nuts, Screws | Fastener Pro",
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
			"title": "Product | Fastener Pro",
			"error": "Invalid product ID",
			"active": "products",
		})
		return
	}

	// 根据ID查询产品
	product, err := models.GetProductByID(uint(id))
	if err != nil {
		c.HTML(http.StatusNotFound, "product-detail.html", gin.H{
			"title": "Product | Fastener Pro",
			"error": "Product not found",
			"active": "products",
		})
		return
	}

	c.HTML(200, "product-detail.html", gin.H{
		"title":   "Product | Fastener Pro",
		"product": product,
		"active":   "products",
	})
}

// 关于我们处理器
func AboutHandler(c *gin.Context) {
	c.HTML(200, "about.html", gin.H{
		"title": "About Us - Professional Fastener Manufacturer | Fastener Pro",
		"active": "about",
	})
}

// 联系我们处理器
func ContactHandler(c *gin.Context) {
	c.HTML(200, "contact.html", gin.H{
		"title": "Contact Us - Fastener Pro | Get Quote for Industrial Fasteners",
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
			"title": "Contact Us - Fastener Pro | Get Quote for Industrial Fasteners",
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
			"title": "Contact Us - Fastener Pro | Get Quote for Industrial Fasteners",
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
		"title":    "Contact Us - Fastener Pro | Get Quote for Industrial Fasteners",
		"success":  "Form submitted successfully",
		"active":   "contact",
	})
}

// FAQ页面处理器
func FAQHandler(c *gin.Context) {
	c.HTML(200, "faq.html", gin.H{
		"title": "FAQ - Fastener Pro | Frequently Asked Questions",
		"active": "faq",
	})
}

// Feedback页面处理器
func FeedbackHandler(c *gin.Context) {
	c.HTML(200, "feedback.html", gin.H{
		"title": "Feedback - Fastener Pro | Send Us an Inquiry",
		"active": "feedback",
	})
}

// Feedback表单提交处理器
func FeedbackSubmitHandler(c *gin.Context) {
	// 解析表单数据
	name := c.PostForm("name")
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	company := c.PostForm("company")
	product := c.PostForm("product")
	message := c.PostForm("message")

	// 验证必填字段
	if name == "" || email == "" || message == "" {
		c.HTML(http.StatusBadRequest, "feedback.html", gin.H{
			"title": "Feedback - Fastener Pro | Send Us an Inquiry",
			"error": "Name, email, and message are required",
			"name":    name,
			"email":   email,
			"phone":   phone,
			"company": company,
			"product": product,
			"message": message,
			"active":   "feedback",
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
		c.HTML(http.StatusInternalServerError, "feedback.html", gin.H{
			"title": "Feedback - Fastener Pro | Send Us an Inquiry",
			"error": "Failed to submit form",
			"name":    name,
			"email":   email,
			"phone":   phone,
			"company": company,
			"product": product,
			"message": message,
			"active":   "feedback",
		})
		return
	}

	// 提交成功
	c.HTML(200, "feedback.html", gin.H{
		"title":    "Feedback - Fastener Pro | Send Us an Inquiry",
		"success":  "Form submitted successfully",
		"active":   "feedback",
	})
}
