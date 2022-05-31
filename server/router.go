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
	r.Use(middleware.Cors())
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

		// 计算利润
		v1.POST("profit", api.CalcProfit)

		// 查询版本号
		v1.GET("version", api.GetVersion)

		// 查询分享链接
		v1.GET("share", api.GetShare)

		// 支付宝回调通知
		v1.POST("alipay/notify", api.AlipayNotify)

		v1.GET("geos", api.GetGeos)

		// 需要token才能访问的接口
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// create order 下单接口
			auth.GET("order", api.CreateOrder)
			// 查询我的佣金统计
			auth.GET("self/commission", api.GetCommissionByUser)
			// 查询佣金明细
			auth.GET("withdraw", api.WithdrawByUser)
			// 查询剩余时长
			auth.GET("expire/time", api.GetExpireTime)

		}

		// 管理员接口
		admin := v1.Group("admin")
		// 管理员需要token才能访问的接口
		admin.GET("/login", api.AdminLogin)
		admin.Use(middleware.AdminRequired())
		{
			// 查询所有节点
			admin.GET("nodes", api.GetAdminNodes)

			// 管理员删除节点
			admin.GET("nodes/delete", api.DeleteNode)

			// 管理员新增节点
			admin.GET("nodes/add", api.AddNode)

			// 管理员编辑节点
			admin.GET("nodes/edit", api.EditNode)
		}
	}
	return r
}
