package routes

import (
	"fastener-pro/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// 首页
	r.GET("/", controllers.HomeHandler)

	// 搜索API
	r.GET("/api/search", controllers.SearchHandler)

	// 产品页
	r.GET("/products.html", controllers.ProductsHandler)
	// 产品分类页
	r.GET("/products-:category", controllers.ProductsHandler)

	// 产品详情页
	r.GET("/product-:id", controllers.ProductDetailHandler)

	// 新闻页
	r.GET("/news.html", controllers.NewsListHandler)

	// 新闻详情页
	r.GET("/news-:id", controllers.NewsDetailHandler)

	// 关于我们
	r.GET("/about.html", controllers.AboutHandler)

	// 联系我们
	r.GET("/contact.html", controllers.ContactHandler)
	// 提交联系表单
	r.POST("/contact.html", controllers.ContactSubmitHandler)

	// FAQ页面
	r.GET("/faq.html", controllers.FAQHandler)

	// 3D设计器页面
	r.GET("/3d.html", controllers.ThreeDHandler)
}
