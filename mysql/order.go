package mysql

import (
	"accelerator/entity/db"
	"accelerator/entity/table"
)

// GetOrderByID get order by id
func GetOrdersByUserID(userID uint) ([]table.Order, error) {
	var orders []table.Order
	if err := db.DB.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
