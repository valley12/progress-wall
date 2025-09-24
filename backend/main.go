package main

import (
	"fmt"
	"log"
	"progress-wall-backend/config"
	"progress-wall-backend/database"
)

func main() {
	// 加载配置
	cfg := config.Load()

	// 打印配置信息
	fmt.Printf("服务器配置:\n")
	fmt.Printf("- 端口: %s\n", cfg.Server.Port)
	fmt.Printf("- 模式: %s\n", cfg.Server.Mode)
	fmt.Printf("- 数据库类型: %s\n", cfg.DB.Type)
	fmt.Printf("- 数据库名称: %s\n", cfg.DB.Name)
	fmt.Printf("- JWT密钥长度: %d\n", len(cfg.JWT.Secret))
	fmt.Printf("- CORS允许来源: %s\n", cfg.CORS.AllowOrigins)

	// 初始化数据库
	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatal("数据库初始化失败:", err)
	}

	// 执行数据库迁移
	if err := database.Migrate(db); err != nil {
		log.Fatal("数据库迁移失败:", err)
	}

	// 初始化基础数据
	if err := database.Seed(db); err != nil {
		log.Fatal("初始化基础数据失败:", err)
	}

	log.Println("数据库初始化完成")
}
