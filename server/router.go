package server

import (
	"accelerator/api"
	"accelerator/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{

		// 服务健康检测
		v1.POST("ping", api.Ping)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 需要token才能访问的接口
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			// auth.GET("user/me", api.UserMe)
			// auth.DELETE("user/logout", api.UserLogout)
		}
	}
	return r
}
