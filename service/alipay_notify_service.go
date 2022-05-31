package service

import (
	"accelerator/entity/errcode"
	"accelerator/entity/response"
	"accelerator/mysql"
	"accelerator/util"

	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
)

type AlipayNotifyService struct {
}

func NewAlipayNotifyService() *AlipayNotifyService {
	return &AlipayNotifyService{}
}

// AlipayNotify 支付宝回调
func (a *AlipayNotifyService) AlipayNotify(c *gin.Context) response.Response {

	o := new(OrderService)
	bm, err := o.ParseNotifyAndVerifySign(c.Request)
	if err != nil {
		util.Log().Error("parse notify and verify sign failed, err: %v", err)
		return response.Response{
			Code: errcode.ErrAlipayNotifySignVerifyFailed,
			Msg:  errcode.Text(errcode.ErrAlipayNotifySignVerifyFailed),
		}
	}
	a.UpdateOrderStatus(bm)

	return response.Response{
		Code: errcode.CodeSuccess,
		Msg:  errcode.Text(errcode.CodeSuccess),
		Data: nil,
	}
}

// UpdateOrderStatus 更新订单状态
func (a *AlipayNotifyService) UpdateOrderStatus(bm gopay.BodyMap) {
	order, err := mysql.GetOrderByOutTradeNo(bm["out_trade_no"].(string))
	if err != nil {
		util.Log().Error("get order by out_trade_no failed, err: %v", err)
		return
	}
	if bm["trade_status"].(string) == "TRADE_SUCCESS" {
		order.Status = mysql.OrderStatusPaid
		// order.EndTime = time.Now()
	}
	if bm["trade_status"].(string) == "TRADE_CLOSED" {
		order.Status = mysql.OrderStatusCanceled
		// order.EndTime = time.Now()
	}
	if bm["trade_status"].(string) == "TRADE_FINISHED" {
		order.Status = mysql.OrderStatusFinished
		// order.EndTime = time.Now()
	}

	order.Status = mysql.OrderStatusPaid
	err = mysql.UpdateOrder(order)
	if err != nil {
		util.Log().Error("update order status failed, err: %v, order: %+v", err, order)
		return
	}
}
