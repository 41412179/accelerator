package mysql

import (
	"accelerator/entity/db"
	"accelerator/entity/table"
	"time"
)

func GetCommissionsByUser(userID int64) ([]*table.Commission, error) {
	var commissions []*table.Commission
	err := db.DB.Where("user_id = ?", userID).Find(&commissions).Error
	return commissions, err
}

func GetCommissionsByChannelId(channelID int64) ([]*table.Commission, error) {
	var commissions []*table.Commission
	err := db.DB.Where("channel_id = ?", channelID).Find(&commissions).Error
	return commissions, err
}

func GetCommissionsByTime(startTime, endTime time.Time) ([]*table.Commission, error) {
	var commissions []*table.Commission
	err := db.DB.Where("created_at >= ? AND created_at <= ?", startTime, endTime).Find(&commissions).Error
	return commissions, err
}

// InsertCommission 插入佣金
func InsertCommission(commission *table.Commission) error {
	return db.DB.Create(commission).Error
}
