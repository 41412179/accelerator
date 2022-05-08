package service

import (
	"accelerator/mysql"
	"accelerator/util"
	"time"
)

type OrderService struct {
}

func NewOrderService() *OrderService {
	return &OrderService{}
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
			remainingTime += order.EndTime.Unix() - int64(time.Now().Unix())
		}
	}

	return remainingTime, nil
}
