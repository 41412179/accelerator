package mysql

import (
	"accelerator/entity/db"
	"accelerator/entity/table"
)

func GetCommissionsByUser(userID int64) ([]*table.Commission, error) {
	var commissions []*table.Commission
	err := db.DB.Where("user_id = ?", userID).Find(&commissions).Error
	return commissions, err
}
