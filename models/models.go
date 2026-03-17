package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Product 产品模型
type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"size:255;not null"`
	Description string    `json:"description" gorm:"type:text"`
	Category    string    `json:"category" gorm:"size:100;not null"`
	Standard    string    `json:"standard" gorm:"size:100"`
	Material    string    `json:"material" gorm:"size:100"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Images      []ProductImage `json:"images" gorm:"foreignKey:ProductID"`
}

// ProductImage 产品图片模型
type ProductImage struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	ProductID uint      `json:"product_id" gorm:"not null"`
	ImageURL  string    `json:"image_url" gorm:"size:255;not null"`
	Order     int       `json:"order" gorm:"default:0"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
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

// InitDB 初始化数据库连接
func InitDB(dsn string) error {
	// 配置GORM日志
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
		},
	)

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// 自动迁移表结构
	err = db.AutoMigrate(&Product{}, &Contact{}, &ProductImage{})
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

// GetProductsWithPagination 分页获取产品
func GetProductsWithPagination(page, pageSize int) ([]Product, int64, error) {
	var products []Product
	var total int64

	// 计算总数
	DB.Model(&Product{}).Count(&total)

	// 分页查询
	offset := (page - 1) * pageSize
	result := DB.Preload("Images").Offset(offset).Limit(pageSize).Find(&products)

	return products, total, result.Error
}

// GetProductsByCategoryWithPagination 分页根据分类获取产品
func GetProductsByCategoryWithPagination(category string, page, pageSize int) ([]Product, int64, error) {
	var products []Product
	var total int64

	// 计算总数
	if category != "" {
		DB.Model(&Product{}).Where("category = ?", category).Count(&total)
	} else {
		DB.Model(&Product{}).Count(&total)
	}

	// 分页查询
	offset := (page - 1) * pageSize
	query := DB.Preload("Images").Offset(offset).Limit(pageSize)
	if category != "" {
		query = query.Where("category = ?", category)
	}
	result := query.Find(&products)

	return products, total, result.Error
}

// GetProductByID 根据ID获取产品
func GetProductByID(id uint) (Product, error) {
	var product Product
	result := DB.Preload("Images").First(&product, id)
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
