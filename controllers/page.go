package controllers

import (
	"fastener-pro/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AboutHandler 关于我们处理器
func AboutHandler(c *gin.Context) {
	c.HTML(200, "about.html", gin.H{
		"title":  "About Us - Professional Fastener Manufacturer | Yuanmao Fastener",
		"active": "about",
	})
}

// ContactHandler 联系我们处理器
func ContactHandler(c *gin.Context) {
	c.HTML(200, "contact.html", gin.H{
		"title":  "Contact Us - Yuanmao Fastener | Get Quote for Industrial Fasteners",
		"active": "contact",
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
		c.HTML(http.StatusBadRequest, "contact.html", gin.H{
			"title":   "Contact Us - Yuanmao Fastener | Get Quote for Industrial Fasteners",
			"error":   "Name, email, and message are required",
			"name":    name,
			"email":   email,
			"phone":   phone,
			"company": company,
			"product": product,
			"message": message,
			"active":  "contact",
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
			"title":   "Contact Us - Yuanmao Fastener | Get Quote for Industrial Fasteners",
			"error":   "Failed to submit form",
			"name":    name,
			"email":   email,
			"phone":   phone,
			"company": company,
			"product": product,
			"message": message,
			"active":  "contact",
		})
		return
	}

	// 提交成功
	c.HTML(200, "contact.html", gin.H{
		"title":   "Contact Us - Yuanmao Fastener | Get Quote for Industrial Fasteners",
		"success": "Form submitted successfully",
		"active":  "contact",
	})
}

// FAQHandler FAQ页面处理器
func FAQHandler(c *gin.Context) {
	c.HTML(200, "faq.html", gin.H{
		"title":  "FAQ - Yuanmao Fastener | Frequently Asked Questions",
		"active": "faq",
	})
}
