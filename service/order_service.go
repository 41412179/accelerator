package service

import (
	"accelerator/entity/errcode"
	"accelerator/entity/response"
	"accelerator/entity/table"
	"accelerator/mysql"
	"accelerator/util"
	"time"

	"github.com/gin-gonic/gin"
)

// OrderService 订单服务
type OrderService struct {
	GoodID    int64  `form:"good_id" json:"good_id" binding:"required"`
	PayType   string `form:"pay_type" json:"pay_type" binding:"required"`
	ChannelID int64  `form:"channel_id" json:"channel_id" binding:"required"`
	user      *table.User
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

// GetExpireTimeByUserId 获取过期时间
func (o *OrderService) GetExpireTimeByUserId(userID int64, remainingTime int64) int64 {
	return time.Now().Unix() + remainingTime
}

// GetRemainingTime 获取剩余时间
func (o *OrderService) GetRemainingTimeByUserId(userId int64) (int64, error) {
	orders, err := mysql.GetOrdersByUserID(userId)
	if err != nil {
		util.Log().Error("get orders by user id err: %+v", err)
		return 0, err
	}
	var remainingTime int64
	for _, order := range orders {
		if order.Status == mysql.OrderStatusPaid {
			if order.EndTime.Unix() > time.Now().Unix() {
				remainingTime += order.EndTime.Unix() - int64(time.Now().Unix())
			}
		}
	}

	return remainingTime, nil
}

// CreateOrder 创建订单
func (o *OrderService) CreateOrder(c *gin.Context) response.Response {
	// 获取用户
	user := util.GetUserByCtx(c)
	if user == nil {
		util.Log().Error("user not logined, user: %v", user)
		return errcode.NewErr(errcode.CodeCheckLogin, nil)
	}
	o.user = user
	// var order table.Order
	order := o.generateOrder()
	if order == nil {
		util.Log().Error("generate order err, user: %+v, orderInfo: %+v", user, o)
		return errcode.NewErr(errcode.CodeDBError, nil)
	}

	// 创建订单
	id, err := mysql.InsertOrder(order)
	if err != nil {
		util.Log().Error("insert order err: %+v", err)
		return errcode.NewErr(errcode.CodeDBError, err)
	}
	// 计算佣金
	go o.computerCommission(id, order)
	return o.setRsponse()

}

// computerCommission 计算佣金
func (o *OrderService) computerCommission(orderId int64, order *table.Order) {
	// 如果没有被邀请人，则不计算佣金
	if o.user.InviterId == 0 {
		return
	}

	// 计算佣金
	c := new(table.Commission)
	c.UserId = o.user.InviterId
	c.OrderId = orderId
	c.Type = table.AddCommissionType
	c.Change = float64(order.PayActualPrice) * float64(0.3) / 100

	err := mysql.InsertCommission(c)
	if err != nil {
		util.Log().Error("insert commission err: %+v", err)
		return
	}

}

func (o *OrderService) setRsponse() response.Response {
	return response.Response{
		Code: errcode.CodeSuccess,
		Msg:  errcode.Text(errcode.CodeSuccess),
	}
}

// generateOrder 生成订单
func (o *OrderService) generateOrder() *table.Order {
	good, err := mysql.GetGoodByID(o.GoodID)
	if err != nil {
		util.Log().Error("get good by id err: %+v", err)
		return nil
	}
	order := new(table.Order)
	order.GoodId = int(o.GoodID)
	order.PayType = o.PayType
	order.UserId = o.user.ID
	order.StartTime = time.Now()
	order.EndTime = time.Now().Add(time.Minute * time.Duration(good.Duration))
	order.Status = mysql.OrderStatusPaid
	order.ChannelId = o.ChannelID
	order.PayActualPrice = float64(good.Price) / 100
	return order
}
