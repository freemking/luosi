package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config 配置结构
type Config struct {
	Database DatabaseConfig `yaml:"database"`
	Mail     MailConfig     `yaml:"mail"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Driver    string `yaml:"driver"`
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	DBName    string `yaml:"dbname"`
	Charset   string `yaml:"charset"`
	ParseTime bool   `yaml:"parseTime"`
	Loc       string `yaml:"loc"`
}

// MailConfig 邮件配置
type MailConfig struct {
	SMTPHost     string `yaml:"smtp_host"`
	SMTPPort     int    `yaml:"smtp_port"`
	SMTPUsername string `yaml:"smtp_username"`
	SMTPPassword string `yaml:"smtp_password"`
	From         string `yaml:"from"`
	To           string `yaml:"to"`
}

// LoadConfig 加载配置文件
func LoadConfig(filePath string) (Config, error) {
	var config Config

	// 读取配置文件
	data, err := os.ReadFile(filePath)
	if err != nil {
		return config, fmt.Errorf("failed to read config file: %w", err)
	}

	// 解析YAML
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("failed to parse config file: %w", err)
	}

	return config, nil
}

// GetDSN 获取数据库连接字符串
func (c *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		c.User, c.Password, c.Host, c.Port, c.DBName, c.Charset, c.ParseTime, c.Loc)
}
