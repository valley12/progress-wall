package main

import (
	"fmt"
	"log"
	"progress-wall-backend/config"
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

	log.Println("配置加载完成")
}
