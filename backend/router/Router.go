package router

import (
	"progress-wall-backend/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type HandlerDependencies struct {
	UserService services.UserService
}

func NewRouter(deps HandlerDependencies) *gin.Engine {
	router := gin.Default()

	// 配置 CORS 中间件
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	userhandler := NewUserHandler(deps.UserService)
	// 用户相关路由
	userGroup := router.Group("/api/users")
	{
		userGroup.POST("/register", userhandler.Register)
	}

	return router
}
