package routes

import (
	"nexfasten/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// 首页
	r.GET("/", controllers.HomeHandler)

	// 搜索API
	r.GET("/api/search", controllers.SearchHandler)

	// 产品页
	r.GET("/product", controllers.ProductsHandler)
	// 产品分类页
	r.GET("/product/:category_slug", controllers.ProductsHandler)

	// 产品详情页
	r.GET("/product/:category_slug/:slug", controllers.ProductDetailHandler)

	// 新闻页
	r.GET("/news", controllers.NewsListHandler)

	// 新闻详情页
	r.GET("/news/:title_slug", controllers.NewsDetailHandler)

	// 关于我们
	r.GET("/about", controllers.AboutHandler)

	// 联系我们
	r.GET("/contact", controllers.ContactHandler)
	// 提交联系表单
	r.POST("/contact/submit", controllers.ContactSubmitHandler)

	// FAQ页面
	r.GET("/faq", controllers.FAQHandler)

	// 3D设计器页面
	r.GET("/3d", controllers.ThreeDHandler)

	// 站点地图
	r.GET("/sitemap.xml", controllers.SitemapHandler)
}
