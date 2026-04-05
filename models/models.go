package models

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var CDNURL string

// SetCDNURL 设置CDN URL
func SetCDNURL(url string) {
	CDNURL = url
}

// prependCDN 为相对路径添加CDN前缀
func prependCDN(path string) string {
	if path == "" {
		return ""
	}
	// 如果已经是完整URL，直接返回
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		return path
	}
	// 添加CDN前缀
	return CDNURL + path
}

// Product 产品模型
type Product struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"size:255;not null"`
	Description template.HTML  `json:"description" gorm:"type:text"`
	Category    string         `json:"category" gorm:"size:100;not null"`
	Standard    string         `json:"standard" gorm:"size:100"`
	Material    string         `json:"material" gorm:"size:100"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	Images      []ProductImage `json:"images" gorm:"foreignKey:ProductID"`
}

// ProductImage 产品图片模型
type ProductImage struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	ProductID uint           `json:"product_id" gorm:"not null"`
	ImageURL  string         `json:"image_url" gorm:"size:255;not null"`
	Order     int            `json:"order" gorm:"default:0"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// Contact 联系表单模型
type Contact struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"size:255;not null"`
	Email     string    `json:"email" gorm:"size:255;not null"`
	Phone     string    `json:"phone" gorm:"size:100"`
	Company   string    `json:"company" gorm:"size:255"`
	Product   string    `json:"product" gorm:"size:255"`
	Message   string    `json:"message" gorm:"type:text;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

// News 新闻模型
type News struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"size:255;not null"`
	CoverImage  string    `json:"cover_image" gorm:"size:255"`
	PublishDate time.Time `json:"publish_date" gorm:"type:date"`
	Summary     string    `json:"summary" gorm:"size:500"`
	Content     string    `json:"content" gorm:"type:text"`
	Status      int       `json:"status" gorm:"default:1"` // 1: 已发布, 0: 草稿
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// AdPosition 广告位模型
type AdPosition struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Code        string         `json:"code" gorm:"size:50;not null"`
	Name        string         `json:"name" gorm:"size:100;not null"`
	Description string         `json:"description" gorm:"size:255"`
	Width       int            `json:"width"`
	Height      int            `json:"height"`
	Status      int            `json:"status" gorm:"default:1"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	Ads         []Ad           `json:"ads" gorm:"foreignKey:PositionID"`
}

// Ad 广告模型
type Ad struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	PositionID  uint           `json:"position_id" gorm:"not null;index"`
	Title       string         `json:"title" gorm:"size:255"`
	Subtitle    string         `json:"subtitle" gorm:"size:500"`
	ImageURL    string         `json:"image_url" gorm:"size:255;not null"`
	LinkURL     string         `json:"link_url" gorm:"size:255"`
	Order       int            `json:"order" gorm:"default:0"`
	Status      int            `json:"status" gorm:"default:1"`
	StartTime   *time.Time     `json:"start_time"`
	EndTime     *time.Time     `json:"end_time"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

// InitDB 初始化数据库连接
func InitDB(dsn string) error {
	// 配置GORM日志 - 只在出错时打印
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Warn, // 只打印警告和错误，避免启动时大量日志
		},
	)

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// 自动迁移表结构（静默执行）
	// Note: AdPosition and Ad tables need to be created manually due to index migration issue
	err = db.AutoMigrate(&Product{}, &Contact{}, &ProductImage{}, &News{})
	// err = db.AutoMigrate(&Product{}, &Contact{}, &ProductImage{}, &News{}, &AdPosition{}, &Ad{})
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	DB = db
	return nil
}

// GetProducts 获取所有产品
func GetProducts() ([]Product, error) {
	var products []Product
	result := DB.Find(&products)
	return products, result.Error
}

// GetProductsWithPagination 分页获取产品（用于首页，只获取名称和图片）
func GetProductsWithPagination(page, pageSize int) ([]Product, int64, error) {
	var products []Product
	var total int64

	// 计算总数
	DB.Model(&Product{}).Count(&total)

	// 分页查询 - 只选择需要的字段
	offset := (page - 1) * pageSize
	result := DB.Select("id, name").Offset(offset).Limit(pageSize).Find(&products)
	if result.Error != nil {
		return products, total, result.Error
	}

	// 批量查询图片
	if len(products) > 0 {
		productIDs := make([]uint, len(products))
		for i, p := range products {
			productIDs[i] = p.ID
		}

		var images []ProductImage
		DB.Select("id, product_id, image_url, `order`").
			Where("product_id IN ?", productIDs).
			Order("`order` ASC").
			Find(&images)

		// 为图片URL添加CDN前缀并按product_id分组，过滤掉空图片
		imageMap := make(map[uint][]ProductImage)
		for i := range images {
			// 跳过空的图片URL（已删除的图片）
			if images[i].ImageURL == "" {
				continue
			}
			images[i].ImageURL = prependCDN(images[i].ImageURL)
			imageMap[images[i].ProductID] = append(imageMap[images[i].ProductID], images[i])
		}

		// 关联图片到产品
		for i := range products {
			products[i].Images = imageMap[products[i].ID]
		}
	}

	return products, total, nil
}

// GetProductsByCategoryWithPagination 分页根据分类获取产品（用于列表页，只获取名称和图片）
func GetProductsByCategoryWithPagination(category string, page, pageSize int) ([]Product, int64, error) {
	var products []Product
	var total int64

	// 计算总数
	if category != "" {
		DB.Model(&Product{}).Where("category = ?", category).Count(&total)
	} else {
		DB.Model(&Product{}).Count(&total)
	}

	// 分页查询 - 只选择需要的字段
	offset := (page - 1) * pageSize
	query := DB.Select("id, name").Offset(offset).Limit(pageSize)
	if category != "" {
		query = query.Where("category = ?", category)
	}
	result := query.Find(&products)
	if result.Error != nil {
		return products, total, result.Error
	}

	// 批量查询图片
	if len(products) > 0 {
		productIDs := make([]uint, len(products))
		for i, p := range products {
			productIDs[i] = p.ID
		}

		var images []ProductImage
		DB.Select("id, product_id, image_url, `order`").
			Where("product_id IN ?", productIDs).
			Order("`order` ASC").
			Find(&images)

		// 为图片URL添加CDN前缀并按product_id分组，过滤掉空图片
		imageMap := make(map[uint][]ProductImage)
		for i := range images {
			// 跳过空的图片URL（已删除的图片）
			if images[i].ImageURL == "" {
				continue
			}
			images[i].ImageURL = prependCDN(images[i].ImageURL)
			imageMap[images[i].ProductID] = append(imageMap[images[i].ProductID], images[i])
		}

		// 关联图片到产品
		for i := range products {
			products[i].Images = imageMap[products[i].ID]
		}
	}

	return products, total, nil
}

// GetProductByID 根据ID获取产品
func GetProductByID(id uint) (Product, error) {
	var product Product
	result := DB.Preload("Images", func(db *gorm.DB) *gorm.DB {
		return db.Order("`order` ASC")
	}).First(&product, id)

	// 为图片URL添加CDN前缀并去重，过滤掉空图片
	seen := make(map[string]bool)
	var uniqueImages []ProductImage
	for i := range product.Images {
		// 跳过空的图片URL（已删除的图片）
		if product.Images[i].ImageURL == "" {
			continue
		}
		product.Images[i].ImageURL = prependCDN(product.Images[i].ImageURL)
		// 去重：只保留第一次出现的图片URL
		if !seen[product.Images[i].ImageURL] {
			seen[product.Images[i].ImageURL] = true
			uniqueImages = append(uniqueImages, product.Images[i])
		}
	}
	product.Images = uniqueImages

	return product, result.Error
}

// GetProductsByCategory 根据分类获取产品
func GetProductsByCategory(category string) ([]Product, error) {
	var products []Product
	result := DB.Where("category = ?", category).Find(&products)
	return products, result.Error
}

// CreateContact 创建联系表单
func CreateContact(contact Contact) error {
	result := DB.Create(&contact)
	return result.Error
}

// GetContacts 获取所有联系表单
func GetContacts() ([]Contact, error) {
	var contacts []Contact
	result := DB.Find(&contacts)
	return contacts, result.Error
}

// GetNewsList 获取新闻列表（分页，只获取列表页需要的字段）
func GetNewsList(page, pageSize int) ([]News, int64, error) {
	var newsList []News
	var total int64

	// 计算总数
	DB.Model(&News{}).Where("status = ?", 1).Count(&total)

	// 分页查询 - 只选择列表页需要的字段
	offset := (page - 1) * pageSize
	result := DB.Select("id, title, cover_image, publish_date, summary, status").
		Where("status = ?", 1).
		Order("publish_date DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&newsList)

	// 为封面图片URL添加CDN前缀
	for i := range newsList {
		newsList[i].CoverImage = prependCDN(newsList[i].CoverImage)
	}

	return newsList, total, result.Error
}

// GetAllNews 获取所有新闻（管理用，只获取列表需要的字段）
func GetAllNews(page, pageSize int) ([]News, int64, error) {
	var newsList []News
	var total int64

	// 计算总数
	DB.Model(&News{}).Count(&total)

	// 分页查询 - 只选择列表需要的字段
	offset := (page - 1) * pageSize
	result := DB.Select("id, title, cover_image, publish_date, summary, status, created_at").
		Order("created_at DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&newsList)

	// 为封面图片URL添加CDN前缀
	for i := range newsList {
		newsList[i].CoverImage = prependCDN(newsList[i].CoverImage)
	}

	return newsList, total, result.Error
}

// GetNewsByID 根据ID获取新闻（详情页需要完整内容）
func GetNewsByID(id uint) (News, error) {
	var news News
	result := DB.First(&news, id)

	// 为封面图片URL添加CDN前缀
	news.CoverImage = prependCDN(news.CoverImage)

	return news, result.Error
}

// CreateNews 创建新闻
func CreateNews(news News) error {
	result := DB.Create(&news)
	return result.Error
}

// UpdateNews 更新新闻
func UpdateNews(news News) error {
	result := DB.Save(&news)
	return result.Error
}

// DeleteNews 删除新闻
func DeleteNews(id uint) error {
	result := DB.Delete(&News{}, id)
	return result.Error
}

// GetNewsCount 获取新闻总数
func GetNewsCount() (int64, error) {
	var count int64
	result := DB.Model(&News{}).Where("status = ?", 1).Count(&count)
	return count, result.Error
}

// SearchProducts searches products by name or description using LIKE query
func SearchProducts(keyword string, limit int) ([]Product, error) {
	var products []Product

	if keyword == "" {
		return products, nil
	}

	// Use LIKE query to search in name and description fields
	searchPattern := "%" + keyword + "%"
	query := DB.Where("name LIKE ? OR mini_description LIKE ?", searchPattern, searchPattern)

	if limit > 0 {
		query = query.Limit(limit)
	}

	result := query.Find(&products)
	if result.Error != nil {
		return products, result.Error
	}

	// Batch query images
	if len(products) > 0 {
		productIDs := make([]uint, len(products))
		for i, p := range products {
			productIDs[i] = p.ID
		}

		var images []ProductImage
		DB.Select("id, product_id, image_url, `order`").
			Where("product_id IN ?", productIDs).
			Order("`order` ASC").
			Find(&images)

		// Add CDN prefix and group by product_id, filter out empty images
		imageMap := make(map[uint][]ProductImage)
		for i := range images {
			// Skip empty image URLs (deleted images)
			if images[i].ImageURL == "" {
				continue
			}
			images[i].ImageURL = prependCDN(images[i].ImageURL)
			imageMap[images[i].ProductID] = append(imageMap[images[i].ProductID], images[i])
		}

		// Associate images with products
		for i := range products {
			products[i].Images = imageMap[products[i].ID]
		}
	}

	return products, nil
}

// GetAdsByPositionCode 根据广告位代码获取广告列表
func GetAdsByPositionCode(code string) ([]Ad, error) {
	var position AdPosition
	result := DB.Where("code = ? AND status = 1", code).First(&position)
	if result.Error != nil {
		return nil, result.Error
	}

	var ads []Ad
	now := time.Now()
	result = DB.Where("position_id = ? AND status = 1", position.ID).
		Where("(start_time IS NULL OR start_time <= ?)", now).
		Where("(end_time IS NULL OR end_time >= ?)", now).
		Order("`order` ASC").
		Find(&ads)

	for i := range ads {
		ads[i].ImageURL = prependCDN(ads[i].ImageURL)
	}

	return ads, result.Error
}
