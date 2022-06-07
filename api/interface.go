package api

import (
	"accelerator/entity/response"
	"accelerator/service"
	"accelerator/util"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var service service.UserService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(200, response.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}

// GetNodes 获取节点列表
func GetNodes(c *gin.Context) {
	var service service.NodeService

	if err := c.ShouldBind(&service); err == nil {
		res := service.GetNodes(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// GetAdminNodes admin获取节点列表
func GetAdminNodes(c *gin.Context) {
	var service service.AdminNodeService

	if err := c.ShouldBind(&service); err == nil {
		res := service.GetNodes(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DeleteNode 删除节点
func DeleteNode(c *gin.Context) {
	var service service.AdminDeleteNodeService

	if err := c.ShouldBind(&service); err == nil {
		res := service.DeleteNode(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// AddNode 添加节点
func AddNode(c *gin.Context) {
	var service service.AdminAddNodeService

	if err := c.ShouldBind(&service); err == nil {
		res := service.AddNode(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// EditNode 编辑节点
func EditNode(c *gin.Context) {
	var service service.AdminEditNodeService

	if err := c.ShouldBind(&service); err == nil {
		res := service.EditNode(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// EditVersion 编辑版本
func EditVersion(c *gin.Context) {
	var service service.EditVersionService

	if err := c.ShouldBind(&service); err == nil {
		res := service.EditVersion()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// GoodList 获取套餐列表
func GoodList(c *gin.Context) {
	var service service.GoodService

	if err := c.ShouldBind(&service); err == nil {
		res := service.GetGoods(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// CreateOrder 创建订单
func CreateOrder(c *gin.Context) {
	var service service.OrderService

	if err := c.ShouldBind(&service); err == nil {
		util.Log().Info("bind order service: %+v", service)
		res := service.CreateOrder(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// GetCommissionByUser 获取用户佣金
func GetCommissionByUser(c *gin.Context) {
	var service service.CommissionService

	if err := c.ShouldBind(&service); err == nil {
		res := service.GetCommissionByUser(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func WithdrawByUser(c *gin.Context) {
	var service service.WithdrawService

	if err := c.ShouldBind(&service); err == nil {
		res := service.WithdrawByUser(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// GetOrdersByChannelID 获取渠道订单
func GetOrdersByChannelID(c *gin.Context) {
	var service service.ChannelOrderService

	if err := c.ShouldBind(&service); err == nil {
		res := service.GetOrdersByChannelID(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// CalcProfit 计算收益
func CalcProfit(c *gin.Context) {
	var service service.ProfitService

	if err := c.ShouldBind(&service); err == nil {
		res := service.CalcProfit(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func GetExpireTime(c *gin.Context) {
	var service service.ExpireService

	if err := c.ShouldBind(&service); err == nil {
		res := service.GetExpireTime(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func GetVersion(c *gin.Context) {
	var service service.VersionService

	if err := c.ShouldBind(&service); err == nil {
		res := service.GetVersion(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func AdminLogin(c *gin.Context) {
	var service service.AdminService

	if err := c.ShouldBind(&service); err == nil {
		res := service.AdminLogin(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func GetShare(c *gin.Context) {
	var service service.ShareService

	if err := c.ShouldBind(&service); err == nil {
		res := service.GetShare(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// AlipayNotify 支付宝通知
func AlipayNotify(c *gin.Context) {
	var service service.AlipayNotifyService

	if err := c.ShouldBind(&service); err == nil {
		res := service.AlipayNotify(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func GetGeos(c *gin.Context) {
	var service service.GeoService

	if err := c.ShouldBind(&service); err == nil {
		res := service.GetGeos(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func ExistUser(c *gin.Context) {
	var service service.ExistUserService

	if err := c.ShouldBind(&service); err == nil {
		res := service.ExistUser(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func CountUser(c *gin.Context) {
	var service service.CountUserService

	if err := c.ShouldBind(&service); err == nil {
		res := service.CountUser(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
