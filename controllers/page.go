package controllers

import (
	"fastener-pro/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AboutHandler 关于我们处理器
func AboutHandler(c *gin.Context) {
	// 获取page header广告
	pageHeaderAd, _ := models.GetAdsByPositionCode("about")
	var firstAd *models.Ad
	if len(pageHeaderAd) > 0 {
		firstAd = &pageHeaderAd[0]
	}

	c.HTML(200, "about.html", gin.H{
		"title":         "About Us - Professional Fastener Manufacturer | Yuanmao Fastener",
		"pageHeaderAd":  firstAd,
		"categories":    models.GetCategories(),
		"active":        "about",
	})
}

// ContactHandler 联系我们处理器
func ContactHandler(c *gin.Context) {
	// 获取page header广告
	pageHeaderAd, _ := models.GetAdsByPositionCode("contact")
	var firstAd *models.Ad
	if len(pageHeaderAd) > 0 {
		firstAd = &pageHeaderAd[0]
	}

	c.HTML(200, "contact.html", gin.H{
		"title":         "Contact Us - Yuanmao Fastener | Get Quote for Industrial Fasteners",
		"pageHeaderAd":  firstAd,
		"categories":    models.GetCategories(),
		"active":        "contact",
	})
}

// ContactSubmitHandler 联系表单提交处理器
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
		c.HTML(http.StatusBadRequest, "contact-error.html", gin.H{
			"title":        "Submission Failed - Yuanmao Fastener",
			"error":        "Name, email, and message are required",
			"categories":   models.GetCategories(),
			"active":       "contact",
		})
		return
	}

	// 创建留言反馈记录
	feedback := models.Feedback{
		Name:    name,
		Email:   email,
		Phone:   phone,
		Company: company,
		Product: product,
		Message: message,
	}

	// 保存到数据库
	err := models.CreateFeedback(feedback)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "contact-error.html", gin.H{
			"title":        "Submission Failed - Yuanmao Fastener",
			"error":        "Failed to submit form: " + err.Error(),
			"categories":   models.GetCategories(),
			"active":       "contact",
		})
		return
	}

	// 提交成功
	c.HTML(200, "contact-success.html", gin.H{
		"title":        "Submission Successful - Yuanmao Fastener",
		"categories":   models.GetCategories(),
		"active":       "contact",
	})
}

// FAQHandler FAQ页面处理器
func FAQHandler(c *gin.Context) {
	// 获取page header广告
	pageHeaderAd, _ := models.GetAdsByPositionCode("faq")
	var firstAd *models.Ad
	if len(pageHeaderAd) > 0 {
		firstAd = &pageHeaderAd[0]
	}

	c.HTML(200, "faq.html", gin.H{
		"title":        "FAQ - Yuanmao Fastener | Frequently Asked Questions",
		"pageHeaderAd": firstAd,
		"categories":   models.GetCategories(),
		"active":       "faq",
	})
}

// ThreeDHandler 3D设计器页面处理器
func ThreeDHandler(c *gin.Context) {
	c.HTML(200, "3d.html", gin.H{
		"title":  "3D Fastener Designer - Visualize & Customize | Yuanmao",
		"active": "products",
	})
}
