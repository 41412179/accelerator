package server

import (
	"accelerator/api"
	"accelerator/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	// session信息
	// r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	// 跨域问题
	// r.Use(middleware.Cors())
	// 获取当前用户
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{

		// 服务健康检测
		v1.POST("ping", api.Ping)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 套餐列表
		v1.POST("good/list", api.GoodList)

		// 节点列表
		v1.GET("nodes", api.GetNodes)

		// 查询某个渠道的order
		v1.GET("channel/orders", api.GetOrdersByChannelID)

		// 需要token才能访问的接口
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// create order 下单接口
			auth.POST("order", api.CreateOrder)
			// 查询我的佣金统计
			auth.GET("self/commission", api.GetCommissionByUser)
			// 查询佣金明细
			auth.GET("withdraw", api.WithdrawByUser)

		}
	}
	return r
}
