package routes

import (
	"fastener-pro/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// 首页
	r.GET("/", controllers.HomeHandler)

	// 产品页
	r.GET("/products", controllers.ProductsHandler)

	// 产品详情页
	r.GET("/product/:id", controllers.ProductDetailHandler)

	// 新闻页
	r.GET("/news", controllers.NewsListHandler)

	// 新闻详情页
	r.GET("/news/:id", controllers.NewsDetailHandler)

	// 关于我们
	r.GET("/about", controllers.AboutHandler)

	// 联系我们
	r.GET("/contact", controllers.ContactHandler)
	// 提交联系表单
	r.POST("/contact", controllers.ContactSubmitHandler)

	// FAQ页面
	r.GET("/faq", controllers.FAQHandler)
}
