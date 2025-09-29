package main

import (
	"log"
	"progress-wall-backend/config"
	"progress-wall-backend/database"
	"progress-wall-backend/models"
	"progress-wall-backend/repository"
	"progress-wall-backend/router"
	"progress-wall-backend/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	cfg := config.Load()

	// 初始化数据库
	if err := database.InitDB(cfg); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 获取数据库实例
	db := database.GetDB()

	// 自动迁移数据库表
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// 设置Gin模式
	gin.SetMode(cfg.Server.Mode)

	// 初始化依赖注入层
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	// 初始化路由
	deps := router.HandlerDependencies{
		UserService: userService,
	}
	r := router.NewRouter(deps)

	// 启动服务器
	log.Printf("Server starting on port %s", cfg.Server.Port)
	if err := r.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
