package mysql

import (
	"accelerator/entity/db"
	// "accelerator/entity/mysql"
	"accelerator/entity/table"
	"time"
	// "gorm.io/driver/mysql"
)

const (
	OrderStatusUnpaid   = 0  // 未支付
	OrderStatusPaying   = 2  // 支付中
	OrderStatusPaid     = 4  // 已支付
	OrderStatusRefund   = 6  // 已退款,未使用这个状态目前
	OrderStatusCanceled = 8  // 已取消,未付款
	OrderStatusFinished = 10 // 已完成

)

// GetOrderByID get order by id
func GetOrdersByUserID(userID int64) ([]table.Order, error) {
	var orders []table.Order
	if err := db.DB.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

// GetOrdersByChannelID get orders by channel id
func GetOrdersByChannelID(channelID int64, startTime, endTime time.Time) ([]table.Order, error) {
	var orders []table.Order

	if err := db.DB.Where("channel_id = ? AND created_at >= ? AND created_at <= ? AND status = ?", channelID, startTime, endTime, OrderStatusPaid).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

// GetPaidOrdersByTime get paid orders by time
func GetPaidOrdersByTime(startTime, endTime time.Time) ([]table.Order, error) {
	var orders []table.Order

	if err := db.DB.Where("created_at >= ? AND created_at <= ? AND status = ?", startTime, endTime, OrderStatusPaid).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func InsertOrder(order *table.Order) (int64, error) {
	if err := db.DB.Create(order).Error; err != nil {
		return 0, err
	}
	return order.Id, nil
}

// GetOrderByOutTradeNo get order by out trade no
func GetOrderByOutTradeNo(outTradeNo string) (*table.Order, error) {
	var order table.Order
	if err := db.DB.Where("out_trade_no = ?", outTradeNo).First(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func UpdateOrder(order *table.Order) error {
	if err := db.DB.Save(order).Error; err != nil {
		return err
	}
	return nil
}
