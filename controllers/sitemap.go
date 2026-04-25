package controllers

import (
	"fmt"
	"net/http"
	"nexfasten/config"
	"nexfasten/models"
	"strings"

	"github.com/gin-gonic/gin"
)

var siteURL string

func SetSiteURL(url string) {
	siteURL = strings.TrimRight(url, "/")
}

func getSiteURL() string {
	if siteURL != "" {
		return siteURL
	}
	cfg, err := config.LoadConfig("conf.yaml")
	if err == nil && cfg.Site.URL != "" {
		siteURL = strings.TrimRight(cfg.Site.URL, "/")
	}
	return siteURL
}

func SitemapHandler(c *gin.Context) {
	baseURL := getSiteURL()

	categories, _ := models.GetSitemapCategories()
	products, _ := models.GetSitemapProducts()
	newsList, _ := models.GetSitemapNews()

	var sb strings.Builder
	sb.WriteString(`<?xml version="1.0" encoding="UTF-8"?>`)
	sb.WriteString("\n")
	sb.WriteString(`<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">`)
	sb.WriteString("\n")

	// 静态页面
	staticURLs := []struct {
		loc        string
		changefreq string
		priority   string
	}{
		{"/", "weekly", "1.0"},
		{"/product", "daily", "0.9"},
		{"/news", "daily", "0.7"},
		{"/about", "monthly", "0.6"},
		{"/contact", "monthly", "0.6"},
		{"/faq", "monthly", "0.5"},
	}

	for _, u := range staticURLs {
		sb.WriteString("  <url>\n")
		sb.WriteString(fmt.Sprintf("    <loc>%s%s</loc>\n", baseURL, u.loc))
		sb.WriteString(fmt.Sprintf("    <changefreq>%s</changefreq>\n", u.changefreq))
		sb.WriteString(fmt.Sprintf("    <priority>%s</priority>\n", u.priority))
		sb.WriteString("  </url>\n")
	}

	// 分类页面
	for _, cat := range categories {
		if cat.Slug != "" {
			sb.WriteString("  <url>\n")
			sb.WriteString(fmt.Sprintf("    <loc>%s/product/%s</loc>\n", baseURL, cat.Slug))
			sb.WriteString("    <changefreq>weekly</changefreq>\n")
			sb.WriteString("    <priority>0.8</priority>\n")
			sb.WriteString("  </url>\n")
		}
	}

	// 产品页面
	for _, p := range products {
		if p.Slug != "" && p.CategorySlug != "" {
			sb.WriteString("  <url>\n")
			sb.WriteString(fmt.Sprintf("    <loc>%s/product/%s/%s</loc>\n", baseURL, p.CategorySlug, p.Slug))
			sb.WriteString("    <changefreq>weekly</changefreq>\n")
			sb.WriteString("    <priority>0.7</priority>\n")
			if !p.UpdatedAt.IsZero() {
				sb.WriteString(fmt.Sprintf("    <lastmod>%s</lastmod>\n", p.UpdatedAt.Format("2006-01-02")))
			}
			sb.WriteString("  </url>\n")
		}
	}

	// 新闻页面
	for _, n := range newsList {
		if n.Slug != "" {
			sb.WriteString("  <url>\n")
			sb.WriteString(fmt.Sprintf("    <loc>%s/news/%s</loc>\n", baseURL, n.Slug))
			sb.WriteString("    <changefreq>weekly</changefreq>\n")
			sb.WriteString("    <priority>0.7</priority>\n")
			if !n.UpdatedAt.IsZero() {
				sb.WriteString(fmt.Sprintf("    <lastmod>%s</lastmod>\n", n.UpdatedAt.Format("2006-01-02")))
			}
			sb.WriteString("  </url>\n")
		}
	}

	sb.WriteString("</urlset>\n")

	sitemapContent := sb.String()
	c.Data(http.StatusOK, "application/xml", []byte(sitemapContent))
}
