package controllers

import (
	"net/http"
	"nexfasten/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
			"title":      "News - Industrial Fastener Industry Updates",
			"error":      "Failed to load news",
			"newsList":   []models.News{},
			"page":       page,
			"pageSize":   pageSize,
			"total":      0,
			"pages":      0,
			"categories": models.GetCategories(),
			"active":     "news",
		})
		return
	}

	// 计算总页数
	pages := int(total) / pageSize
	if int(total)%pageSize > 0 {
		pages++
	}

	// 获取page header广告
	pageHeaderAd, _ := models.GetAdsByPositionCode("news")
	var firstAd *models.Ad
	if len(pageHeaderAd) > 0 {
		firstAd = &pageHeaderAd[0]
	}

	c.HTML(200, "news.html", gin.H{
		"title":         "News - Industrial Fastener Industry Updates",
		"newsList":      newsList,
		"page":          page,
		"pageSize":      pageSize,
		"total":         total,
		"pages":         pages,
		"pageHeaderAd":  firstAd,
		"categories":    models.GetCategories(),
		"active":        "news",
	})
}

// NewsDetailHandler 新闻详情处理器
func NewsDetailHandler(c *gin.Context) {
	slug := c.Param("title_slug")

	news, err := models.GetNewsBySlug(slug)
	if err != nil {
		c.HTML(http.StatusNotFound, "news-detail.html", gin.H{
			"title":      "News | Yuanmao Fastener",
			"error":      "News not found",
			"categories": models.GetCategories(),
			"active":     "news",
		})
		return
	}

	c.HTML(200, "news-detail.html", gin.H{
		"title":      news.Title + " - Yuanmao Fastener News",
		"news":       news,
		"categories": models.GetCategories(),
		"active":     "news",
	})
}
