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
