package main

import (
	"fastener-pro/config"
	"fastener-pro/cron"
	"fastener-pro/models"
	"fastener-pro/routes"
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置文件
	cfg, err := config.LoadConfig("conf.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 初始化数据库连接
	dsn := cfg.Database.GetDSN()
	err = models.InitDB(dsn)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 设置定时任务
	cron.SetupCron(cfg)

	// 创建Gin引擎
	r := gin.Default()

	// 静态文件服务
	r.Static("/static", "./static")

	// 添加模板函数
	r.SetFuncMap(template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
		"sub": func(a, b int) int {
			return a - b
		},
		"mul": func(a, b int) int {
			return a * b
		},
		"eq": func(a, b interface{}) bool {
			return a == b
		},
		"seq": func(start, end int) []int {
			seq := make([]int, end-start+1)
			for i := range seq {
				seq[i] = start + i
			}
			return seq
		},
	})

	// 模板加载
	r.LoadHTMLGlob("./templates/*")

	// 注册路由
	routes.SetupRoutes(r)

	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
