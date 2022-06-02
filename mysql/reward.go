package mysql

import (
	"accelerator/entity/db"
	"accelerator/entity/table"
)

func GetRewardByQq(qq string) (*table.Reward, error) {
	var reward table.Reward
	if err := db.DB.Where("qq = ?", qq).First(&reward).Error; err != nil {
		return nil, err
	}
	return &reward, nil
}

func InsertReward(reward *table.Reward) (int64, error) {
	result := db.DB.Create(reward)
	return result.RowsAffected, result.Error
}
